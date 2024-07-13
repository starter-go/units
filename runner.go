package units

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/starter-go/application"
	"github.com/starter-go/base/util"
	"github.com/starter-go/starter"
	"github.com/starter-go/units/src/main/golang/unitcore"
	"github.com/starter-go/vlog"
)

// Runner 单元测试执行器
type Runner interface {

	// Dependencies(deps ...application.Module) Runner
	// ModuleT(mb *application.ModuleBuilder) Runner

	Module(m application.Module) Runner

	Testing(t *testing.T) Runner

	EnablePanic(enabled bool) Runner

	SetProperties(table map[string]string) Runner

	Run(args []string) error
}

// NewRunner 新建一个 Runner
func NewRunner() Runner {
	return &runner{enPanic: true}
}

////////////////////////////////////////////////////////////////////////////////

type runner struct {
	deps []application.Module
	// mb      *application.ModuleBuilder
	mod     application.Module
	t       *testing.T
	props   map[string]string
	enPanic bool
}

// func (inst *runner) Dependencies(mods ...application.Module) Runner {
// 	inst.deps = mods
// 	return inst
// }

// func (inst *runner) ModuleT(mb *application.ModuleBuilder) Runner {
// 	inst.mb = mb
// 	return inst
// }

func (inst *runner) Module(m application.Module) Runner {
	inst.mod = m
	return inst
}

func (inst *runner) SetProperties(t map[string]string) Runner {
	inst.props = t
	return inst
}

func (inst *runner) Testing(t *testing.T) Runner {
	inst.t = t
	return inst
}

func (inst *runner) EnablePanic(enabled bool) Runner {
	inst.enPanic = enabled
	return inst
}

func (inst *runner) Run(args []string) error {
	r := &innerRunner{parent: inst}
	err := r.run(args)
	if err != nil {
		if inst.enPanic {
			panic(err)
		} else {
			vlog.Error(err.Error())
		}
	}
	return err
}

////////////////////////////////////////////////////////////////////////////////

type innerRunner struct {
	parent  *runner
	current *Registration
	err     error
	skipped bool

	countTotal   int
	countSkipped int
	countRun     int
	countError   int
}

func (inst *innerRunner) run(args []string) error {
	i := starter.Init(args)
	i.MainModule(inst.getTargetModule())
	i.WithPanic(inst.parent.enPanic)
	inst.setupRunnerHolder(i)
	inst.loadAdditionProps(i)
	err := i.Run()
	if err != nil {
		return err
	}
	return inst.countErrors()
}

func (inst *innerRunner) getTargetModule() application.Module {
	return inst.parent.mod
}

func (inst *innerRunner) countErrors() error {

	const f = "units::count total:%d run:%d skip:%d error:%d"
	vlog.Info(f, inst.countTotal, inst.countRun, inst.countSkipped, inst.countError)

	count := inst.countError
	if count == 0 {
		vlog.Info("units::success")
		return nil
	}
	return fmt.Errorf("%d error(s)", count)
}

// func (inst *innerRunner) makeCoreModule() application.Module {
// 	mb := ModuleT()
// 	mb.Components(main4units.ExportConfig)
// 	return mb.Create()
// }

// func (inst *innerRunner) makeMainModule() application.Module {

// 	core := inst.makeCoreModule()

// 	mb := inst.parent.mb
// 	if mb == nil {
// 		mb = &application.ModuleBuilder{}
// 	}
// 	mb.Name(core.Name() + "#main")
// 	mb.Version(core.Version())
// 	mb.Revision(core.Revision())
// 	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)

// 	deps := inst.parent.deps
// 	if deps != nil {
// 		mb.Depend(deps...)
// 	}
// 	mb.Depend(core)

// 	return mb.Create()
// }

func (inst *innerRunner) setupRunnerHolder(i starter.Initializer) {
	holder := &unitcore.RunnerHolder{Run: inst.runWithContext}
	name := holder.AttributeName()
	i.GetAttributes().SetAttribute(name, holder)
}

func (inst *innerRunner) loadAdditionProps(i starter.Initializer) {
	src := inst.parent.props
	dst := i.GetProperties()
	for key, val := range src {
		dst.SetProperty(key, val)
	}
}

func (inst *innerRunner) runWithContext(ctx *unitcore.Context) error {
	all, err := inst.listUnits(ctx.ApplicationContext)
	if err != nil {
		return err
	}
	if ctx.RunAll {
		for _, item := range all {
			err := inst.runUnit(ctx, item)
			inst.handleError(err)
		}
	} else {
		namelist := ctx.RunList
		for _, name := range namelist {
			item, err := inst.findUnitByName(name, all)
			if err != nil {
				return err
			}
			err = inst.runUnit(ctx, item)
			inst.handleError(err)
		}
	}
	return nil
}

func (inst *innerRunner) findUnitByName(name string, all []*Registration) (*Registration, error) {
	want := strings.TrimSpace(name)
	for _, reg := range all {
		have := strings.TrimSpace(reg.Name)
		if have == want {
			return reg, nil
		}
	}
	return nil, fmt.Errorf("no test unit with name: %s", want)
}

func (inst *innerRunner) listUnits(ac application.Context) ([]*Registration, error) {

	// list all
	comSet := ac.GetComponents()
	ids := comSet.ListIDs()
	all := make([]*Registration, 0)

	for _, id := range ids {
		h, err := comSet.Get(id)
		if err != nil {
			return nil, err
		}
		ci, err := h.GetInstance()
		if err != nil {
			return nil, err
		}
		obj := ci.Get()
		reg, ok := obj.(Units)
		if ok {
			all = reg.Units(all)
		}
	}

	// filter
	dst := make([]*Registration, 0)
	for _, item := range all {
		if item == nil {
			continue
		}
		if item.Test != nil {
			dst = append(dst, item)
		}
	}

	// sort
	s := util.Sorter{
		OnSwap: func(i1, i2 int) { dst[i1], dst[i2] = dst[i2], dst[i1] },
		OnLen:  func() int { return len(dst) },
		OnLess: func(i1, i2 int) bool {
			p1 := dst[i1].Priority
			p2 := dst[i2].Priority
			if p1 == p2 {
				name1 := dst[i1].Name
				name2 := dst[i2].Name
				return strings.Compare(name1, name2) < 0
			}
			return p1 > p2
		},
	}
	s.Sort()

	inst.countTotal = len(all)
	return dst, nil
}

func (inst *innerRunner) runUnit(ctx *unitcore.Context, u *Registration) error {

	defer func() {
		inst.onUnitEnd(ctx, u)
	}()
	defer func() {
		x := recover()
		inst.handleErrorX(x)
	}()
	inst.onUnitBegin(ctx, u)

	if !u.Enabled {
		inst.skipped = true
		inst.countSkipped++
		return nil
	}

	inst.countRun++
	err := u.Test()
	inst.handleError(err)
	return err
}

func (inst *innerRunner) onUnitBegin(ctx *unitcore.Context, u *Registration) {
	inst.err = nil
	inst.skipped = false
	inst.current = u
}

func (inst *innerRunner) onUnitEnd(ctx *unitcore.Context, u *Registration) {
	if u == nil {
		return
	}
	name := u.Name
	p := u.Priority
	// fmt.Printf("[unit name:'%s'  Priority:%d   ]", name, p)

	b := strings.Builder{}
	b.WriteString("Unit priority:")
	b.WriteString(strconv.Itoa(p))
	b.WriteString(" name:'")
	b.WriteString(name)
	b.WriteString("'")
	b.WriteString(" ................ ")

	err := inst.err
	if err == nil {
		if inst.skipped {
			b.WriteString("[skipped]")
		} else {
			b.WriteString("[OK]")
		}
		vlog.Info(b.String())
	} else {
		b.WriteString("Error:")
		b.WriteString(err.Error())
		vlog.Error(b.String())
		inst.countError++
	}

	inst.current = nil
}

func (inst *innerRunner) handleError(err error) {
	if err == nil {
		return
	}
	inst.err = err
}

func (inst *innerRunner) handleErrorX(x any) {
	if x == nil {
		return
	}

	err, ok := x.(error)
	if ok {
		inst.handleError(err)
	}

	str, ok := x.(string)
	if ok {
		inst.handleError(fmt.Errorf(str))
	}
}
