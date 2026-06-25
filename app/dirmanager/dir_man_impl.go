package dirmanager

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/starter-go/afs"
	"github.com/starter-go/application"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/units"
)

type DirManagerImpl struct {

	//starter:component

	_as func(units.DirManager) //starter:as("#")

	AC  application.Context //starter:inject("context")
	AFS afs.FS              //starter:inject("#")

	TheUnitsWorkingDir string //starter:inject("${units.working.dir}")

}

// GetDir implements units.DirManager.
func (inst *DirManagerImpl) GetDir(want *units.DirHolder) (*units.DirHolder, error) {

	if want == nil {
		return nil, fmt.Errorf("want is nil")
	}

	steps := make([]func(*locating) error, 0)

	h2 := new(locating)
	h2.want = want
	h2.have = new(units.DirHolder)

	steps = append(steps, inst.innerDoInit)
	steps = append(steps, inst.innerDoPrepareContext)
	steps = append(steps, inst.innerDoPrepareModuleName)
	steps = append(steps, inst.innerDoPrepareUnitName)
	steps = append(steps, inst.innerDoPrepareRuntimeName)
	steps = append(steps, inst.innerDoPrepareScope)
	steps = append(steps, inst.innerDoPrepareKey)
	steps = append(steps, inst.innerDoBuildPath)
	steps = append(steps, inst.innerDoMakeResult)

	for _, st := range steps {
		err := st(h2)
		if err != nil {
			return nil, err
		}
	}

	return h2.have, nil
}

func (inst *DirManagerImpl) innerDoPrepareModuleName(h2 *locating) error {

	tc := h2.tc
	mod := tc.Module
	name := mod.Name()

	if name == "" {
		name = "unnamed"
	}

	name = strings.ReplaceAll(name, "#", "_")

	h2.moduleRef = mod
	h2.moduleName = name

	return nil
}

func (inst *DirManagerImpl) innerDoPrepareUnitName(h2 *locating) error {

	uc := h2.uc
	reg := uc.Unit
	name := reg.Name

	if name == "" {
		name = "unnamed"
	}

	h2.unitRef = reg
	h2.unitName = name

	return nil
}

func (inst *DirManagerImpl) innerDoPrepareRuntimeName(h2 *locating) error {

	tc := h2.tc
	t := tc.StartedAt
	n := t.Int()

	str := strconv.FormatInt(n, 10)

	h2.runAtTime = t
	h2.runAtStr = str

	return nil
}

func (inst *DirManagerImpl) innerDoInit(h2 *locating) error {

	dir := inst.TheUnitsWorkingDir // "~/.starter/units"

	pb := &h2.builder
	pb.WriteString(dir)
	return nil
}

func (inst *DirManagerImpl) innerDoPrepareContext(h2 *locating) error {

	cc := h2.cc
	want := h2.want

	if cc == nil {
		cc = want.Context
	}
	if cc == nil {
		cc = context.Background()
	}

	holder, err := units.GetContextHolder(cc)
	if err != nil {
		return err
	}

	tc := holder.TC
	uc := holder.UC

	h2.cc = cc
	h2.tc = tc
	h2.uc = uc

	return nil
}

// func (inst *DirManagerImpl) innerGetRuntimeStartedAt(cc context.Context) (lang.Time, error) {
// 	tc, err := units.GetContext(cc)
// 	if err != nil {
// 		return 0, err
// 	}
// 	at := tc.StartedAt
// 	return at, nil
// }

func (inst *DirManagerImpl) innerDoPrepareScope(h2 *locating) error {

	want := h2.want
	scope := h2.scope

	if scope == "" {
		scope = want.Scope
	}
	if scope == "" {
		scope = units.DirScopeRuntime
	}

	h2.scope = scope

	return nil
}

func (inst *DirManagerImpl) innerDoPrepareKey(h2 *locating) error {

	want := h2.want
	key := h2.key

	if key == "" {
		key = want.Key
	}

	if key == "" {
		key = units.DirKeyTemp
	}

	h2.key = key

	return nil
}

func (inst *DirManagerImpl) innerDoBuildPath(h2 *locating) error {

	scope := h2.scope
	b := &h2.builder

	// module
	mod := h2.moduleName
	b.WriteString("/modules/")
	b.WriteString(mod)

	if scope == units.DirScopeModule {
		return nil
	}

	// unit

	unit := h2.unitName
	b.WriteString("/units/")
	b.WriteString(unit)

	if scope == units.DirScopeUnit {
		return nil
	}

	// runtime

	at := h2.runAtStr
	b.WriteString("/at/")
	b.WriteString(at)

	return nil
}

func (inst *DirManagerImpl) innerDoMakeResult(h2 *locating) error {

	b := &h2.builder
	key := h2.key

	b.WriteString("/")
	b.WriteString(string(key))

	// build

	pathStr := b.String()
	fs := inst.AFS
	path := fs.NewPath(pathStr)

	// have

	have := new(units.DirHolder)
	want := h2.want

	have.Path = path
	have.Context = h2.cc
	have.Key = want.Key
	have.Scope = want.Scope
	have.Unit = h2.unitRef

	h2.path = path
	h2.have = have

	return nil
}

func (inst *DirManagerImpl) _impl() units.DirManager {
	return inst
}

////////////////////////////////////////////////////////////////////////////////

type locating struct {
	cc context.Context
	tc *units.TestContext
	uc *units.UnitContext

	want *units.DirHolder
	have *units.DirHolder

	key   units.DirKey
	scope units.DirScope

	// module, unit, runtime

	moduleRef  application.Module
	moduleName string

	unitRef  *units.Registration
	unitName string

	runAtTime lang.Time
	runAtStr  string

	// builder

	builder strings.Builder // the path-builder
	path    afs.Path
}

////////////////////////////////////////////////////////////////////////////////
// EOF
