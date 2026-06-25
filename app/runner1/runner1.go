package runner1

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/starter-go/application"
	"github.com/starter-go/base/lang"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

type innerRunnerV1config struct {
	t          *testing.T
	module     application.Module
	usePanic   bool
	properties map[string]string
}

type RunnerV1 struct {

	//starter:component

	_as func(units.RunnerRegistry) //starter:as(".")

	// UnitCompList []units.Unit //x-starter:inject(".")

	UnitNameList string              //starter:inject("${units.list}")
	AC           application.Context //starter:inject("context")

	MyName  string //starter:inject("${runner.runner1.name}")
	MyAlias string //starter:inject("${runner.runner1.alias}")
	Enabled bool   //starter:inject("${runner.runner1.enabled}")

}

// // EnablePanic implements units.Runner.
// func (inst *RunnerV1) EnablePanic(enabled bool) units.Runner {
// 	// inst.config.usePanic = enabled
// 	return inst
// }

// // Module implements units.Runner.
// func (inst *RunnerV1) Module(m application.Module) units.Runner {
// 	// inst.config.module = m
// 	return inst
// }

// Run implements units.Runner.
func (inst *RunnerV1) Run(c *units.Context) error {

	steps := make([]func(rc *runningContext) error, 0)

	rc := new(runningContext)
	rc.tc = c
	rc.cc = c.CC

	steps = append(steps, inst.innerDoPrepareContexts)
	steps = append(steps, inst.innerDoLoadTasks)
	steps = append(steps, inst.innerDoLogInitState)
	steps = append(steps, inst.innerDoRunTasks)
	steps = append(steps, inst.innerDoLogFinishState)
	steps = append(steps, inst.innerDoCheckErrorList)

	for _, fn := range steps {
		err := fn(rc)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *RunnerV1) innerHandleErrorX(err error, x any) {

	if err == nil && x == nil {
		return
	}

	if x != nil {
		err = fmt.Errorf("panic(%v)", x)
	}

	vlog.Error("%s", err.Error())
}

func (inst *RunnerV1) innerDoXX(rc *runningContext) error {
	return nil
}

func (inst *RunnerV1) innerDoRunTasks(rc *runningContext) error {

	list := rc.tasks
	tc := rc.tc

	defer func() {
		x := recover()
		inst.innerHandleErrorX(nil, x)
		tc.StoppedAt = lang.Now()
	}()

	tc.StartedAt = lang.Now()

	for idx, task := range list {
		task.Index = idx
		err := inst.innerRun1(rc, task)
		if err != nil {
			inst.innerHandleErrorX(err, nil)
			rc.lastError = err
		}
		if rc.abort {
			break
		}
	}

	return nil
}

func (inst *RunnerV1) innerDoLogInitState(rc *runningContext) error {

	vlog.Info("tasks-initial-state:")

	return inst.innerDoLogAllTasksCurrentState(rc)
}

func (inst *RunnerV1) innerDoCheckErrorList(rc *runningContext) error {
	var err error
	all := rc.tasks
	for _, it := range all {
		e := it.Error
		if e == nil {
			continue
		}
		err = e
	}
	return err
}

func (inst *RunnerV1) innerDoLogFinishState(rc *runningContext) error {

	vlog.Info("tasks-finished-state:")

	return inst.innerDoLogAllTasksCurrentState(rc)
}

func (inst *RunnerV1) innerDoLogAllTasksCurrentState(rc *runningContext) error {
	logger := new(taskStateLogger)
	logger.Log(rc.tasks)
	return nil
}

func (inst *RunnerV1) innerReloadUnitContext(rc *runningContext, t *units.UnitHolder) (*units.UnitContext, error) {

	cc := rc.cc
	tc := rc.tc
	uc := new(units.UnitContext)

	holder, err := units.GetContextHolder(cc)
	if err != nil {
		return nil, err
	}

	p1 := tc.Properties
	p2 := make(map[string]string)
	for k, v := range p1 {
		p2[k] = v
	}

	uc.Parent = tc
	uc.Unit = t.Ref
	uc.Properties = p2

	holder.UC = uc
	rc.uc = uc

	return uc, nil
}

func (inst *RunnerV1) innerRun1(rc *runningContext, t *units.UnitHolder) error {

	if !inst.innerAcceptTask(t) {
		t.State = units.TaskStateIgnored
		return nil
	}

	uc, err := inst.innerReloadUnitContext(rc, t)
	if err != nil {
		return err
	}

	defer func() {

		uc.StoppedAt = lang.Now()

		x := recover()
		if x == nil {
			return
		}

		rc.abort = true
		t.Error = fmt.Errorf("panic(%v)", x)
		t.Done = true
		t.State = units.TaskStatePanic

		// panic(x)
	}()

	uc.StartedAt = lang.Now()

	err = inst.innerRun2(rc, t)
	if err != nil {
		rc.abort = true
		t.Error = err
		t.State = units.TaskStateError
	} else {
		t.State = units.TaskStateOK
	}

	return err
}

func (inst *RunnerV1) innerRun2(rc *runningContext, t *units.UnitHolder) error {

	now := lang.Now()
	bar := inst.innerGetBarString()
	cc := rc.cc

	vlog.Info(bar)
	inst.innerLogTaskInfo(t)

	defer func() {
		now = lang.Now()
		t.Done = true
		t.StoppedAt = now

		vlog.Info("Done.")
		vlog.Info(bar)

	}()

	t.StartedAt = now
	fn := t.Info.Do

	err := fn(cc)

	if err == nil {
		t.OK = true
	} else {
		t.Error = err
	}

	err = inst.innerFilterError(t, err)

	return err
}

func (inst *RunnerV1) innerLogTaskInfo(t *units.UnitHolder) {

	const nl = "\n"
	builder := new(strings.Builder)

	builder.WriteString("units.Unit" + nl)
	builder.WriteString("  name: " + t.Info.Name + nl)
	builder.WriteString("  id: " + t.Info.ID + nl)
	builder.WriteString("  description: " + t.Info.Description + nl)

	str := builder.String()
	vlog.Info("%s", str)
}

func (inst *RunnerV1) innerGetBarString() string {
	return "////////////////////////////////////////////////////////////////////////////////"
}

func (inst *RunnerV1) innerAcceptTask(t *units.UnitHolder) bool {

	if t == nil {
		return false
	}

	if !t.Selected {
		return false
	}

	if !t.Info.Enabled {
		return false
	}

	return true
}

func (inst *RunnerV1) innerFilterError(t *units.UnitHolder, err error) error {

	if err == nil {
		return nil
	}

	method := t.Info.OnError
	switch method {

	case units.OnErrorWarn:
		vlog.Warn("%s", err.Error())
		return nil

	case units.OnErrorPanic:
		panic(err)

	case units.OnErrorMute:
	case units.OnErrorIgnore:
	case units.OnErrorNone:
		return nil

	case units.OnErrorAbort:
	case units.OnErrorError:
	default:
	}
	return err
}

func (inst *RunnerV1) innerLoadUnitList() []units.Unit {

	all := inst.AC.GetComponents()
	ids := all.ListIDs()
	dst := make([]units.Unit, 0)

	for _, id := range ids {
		h, err := all.Get(id)
		if h == nil || err != nil {
			continue
		}
		ins, err := h.GetInstance()
		if err != nil {
			continue
		}
		obj := ins.Get()
		unit, ok := obj.(units.Unit)
		if ok {
			dst = append(dst, unit)
		}
	}

	return dst
}

func (inst *RunnerV1) innerDoPrepareContexts(rc *runningContext) error {

	cc := rc.cc
	tc := rc.tc
	uc := new(units.UnitContext)
	ac := inst.AC

	if cc == nil {
		cc = context.Background()
	}

	holder, err := units.SetupContextHolder(cc)
	if err != nil {
		return err
	}

	cc = holder.CC

	rc.cc = cc
	rc.tc = tc
	rc.uc = uc

	tc.Module = ac.GetMainModule()

	uc.Parent = tc
	uc.Properties = tc.Properties

	holder.CC = cc
	holder.TC = tc
	holder.UC = uc

	return nil
}

func (inst *RunnerV1) innerDoLoadTasks(rc *runningContext) error {

	loader := new(innerTaskLoader)
	rawNameList := inst.UnitNameList
	rawUnitList := inst.innerLoadUnitList()
	tasks, err := loader.load(rawUnitList, rawNameList)

	if err == nil {

		tc := rc.tc

		rc.tasks = tasks
		tc.Units = tasks
	}

	return err
}

// // SetProperties implements units.Runner.
// func (inst *RunnerV1) SetProperties(table map[string]string) units.Runner {
// 	inst.config.properties = table
// 	return inst
// }

// // Testing implements units.Runner.
// func (inst *RunnerV1) Testing(t *testing.T) units.Runner {
// 	// inst.config.t = t
// 	return inst
// }

// GetRegistration implements units.RunnerRegistry.
func (inst *RunnerV1) GetRegistration(ret *units.RunnerRegistration) error {

	ret.Runner = inst
	ret.Name = inst.MyName
	ret.Alias = inst.MyAlias
	ret.Enabled = inst.Enabled

	return nil
}

func (inst *RunnerV1) _impl() units.RunnerRegistry {
	return inst
}

////////////////////////////////////////////////////////////////////////////////
// EOF
