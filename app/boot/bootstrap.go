package boot

import (
	"fmt"
	"os"

	"github.com/starter-go/application"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

type Boot struct {

	//starter:component

	AC          application.Context    //starter:inject("context")
	RunnerList  []units.RunnerRegistry //starter:inject(".")
	RunnerAlias string                 //starter:inject("${units.runner}")

	runner units.Runner // the cached runner
}

// Life implements application.Lifecycle.
func (inst *Boot) Life() *application.Life {
	l := new(application.Life)
	l.OnStart = inst.onStart
	l.OnLoop = inst.run
	return l
}

func (inst *Boot) onStart() error {
	inst.innerListAllUnits()
	return nil
}

func (inst *Boot) run() error {

	a := os.Args
	c := units.NewContext()

	c.Arguments = a

	runner, err := inst.innerGetRunner()
	if err != nil {
		return err
	}
	return runner.Run(c)
}

func (inst *Boot) innerGetRunner() (units.Runner, error) {
	r := inst.runner
	if r == nil {
		loader := new(innerRunnerLoader)
		loader.boot = inst
		loaded, err := loader.load()
		if err != nil {
			return nil, err
		}
		r = loaded
		inst.runner = loaded
	}
	return r, nil
}

func (inst *Boot) innerListAllUnits() error {

	// 列出所有的测试单元

	const bar = "---------------------------------------------------------------"

	vlog.Info("List All Units")
	vlog.Info(bar)

	ac := inst.AC
	allcom := ac.GetComponents()
	ids := allcom.ListIDs()
	src := make([]units.Unit, 0)
	tmp := make([]*units.Registration, 0)

	for _, id := range ids {
		holder, err := allcom.Get(id)
		if err != nil {
			continue
		}
		ref, err := holder.GetInstance()
		if err != nil {
			continue
		}
		com := ref.Get()
		u1, ok := com.(units.Unit)
		if ok {
			src = append(src, u1)
		}
	}

	for _, u := range src {
		tmp = u.ListRegistrations(tmp)
	}

	for _, u := range tmp {
		name := u.Name
		id := u.ID
		cl := u.Class
		en := u.Enabled
		priority := u.Priority
		vlog.Info("units.[Unit id:'%s' name:'%s' class:'%s' enabled:%v priority:%v]", id, name, cl, en, priority)
	}

	vlog.Info(bar)
	return nil
}

func (inst *Boot) _impl() application.Lifecycle {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type innerRunnerLoader struct {
	boot *Boot
}

func (inst *innerRunnerLoader) load() (units.Runner, error) {
	info, err := inst.innerLoadOne()
	if err != nil {
		return nil, err
	}
	return info.Runner, nil
}

func (inst *innerRunnerLoader) innerLoadAll() (map[string]*units.RunnerRegistration, error) {

	dst := make(map[string]*units.RunnerRegistration)
	src := inst.boot.RunnerList

	for _, reg1 := range src {

		info, err := inst.innerGetRunnerInfo(reg1)
		if err != nil {
			return nil, err
		}

		err = inst.innerPutRunnerInfo(info, dst)
		if err != nil {
			vlog.Warn("%s", err.Error())
		}
	}

	return dst, nil
}

func (inst *innerRunnerLoader) innerLoadOne() (*units.RunnerRegistration, error) {

	all, err := inst.innerLoadAll()
	if err != nil {
		return nil, err
	}

	alias := inst.boot.RunnerAlias
	res := all[alias]
	if res == nil {
		return nil, fmt.Errorf("no units.Runner with name(alias) of '%s'", alias)
	}

	return res, nil
}

func (inst *innerRunnerLoader) innerPutRunnerInfo(info *units.RunnerRegistration, table map[string]*units.RunnerRegistration) error {

	if info == nil || table == nil {
		return fmt.Errorf("param(s) (info,table) is nil")
	}

	if !info.Enabled {
		return nil
	}

	if info.Runner == nil {
		return fmt.Errorf("runner is nil")
	}

	// put with name
	name1 := info.Alias
	name2 := info.Name
	older1 := table[name1]
	older2 := table[name2]
	if older1 != nil {
		return fmt.Errorf("units:runner:name (alias) '%s' is duplicate", name1)
	}
	if older2 != nil {
		return fmt.Errorf("units:runner:name (alias) '%s' is duplicate", name2)
	}

	table[name1] = info
	table[name2] = info
	return nil
}

func (inst *innerRunnerLoader) innerGetRunnerInfo(r1 units.RunnerRegistry) (*units.RunnerRegistration, error) {
	r2 := new(units.RunnerRegistration)
	err := r1.GetRegistration(r2)
	return r2, err
}

////////////////////////////////////////////////////////////////////////////////
// EOF
