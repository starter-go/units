package runner1

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/starter-go/application"
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

	MyName  string //starter:inject("${runner.UnitTestRunnerV1.name}")
	MyAlias string //starter:inject("${runner.UnitTestRunnerV1.alias}")
	Enabled bool   //starter:inject("${runner.UnitTestRunnerV1.enabled}")

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

	steps = append(steps, inst.innerDoLoadTasks)
	steps = append(steps, inst.innerDoLogInitState)
	steps = append(steps, inst.innerDoRunTasks)
	steps = append(steps, inst.innerDoLogFinishState)
	steps = append(steps, inst.innerDoCheckLastError)

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

	defer func() {
		x := recover()
		inst.innerHandleErrorX(nil, x)
	}()

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

func (inst *RunnerV1) innerDoCheckLastError(rc *runningContext) error {
	return rc.lastError
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

func (inst *RunnerV1) innerRun1(rc *runningContext, t *innerTask) error {

	if !inst.innerAcceptTask(t) {
		t.State = units.TaskStateIgnored
		return nil
	}

	defer func() {
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

	err := inst.innerRun2(rc, t)
	if err != nil {
		rc.abort = true
		t.Error = err
		t.State = units.TaskStateError
	} else {
		t.State = units.TaskStateOK
	}

	return err
}

func (inst *RunnerV1) innerRun2(rc *runningContext, t *innerTask) error {

	now := time.Now()
	bar := inst.innerGetBarString()

	vlog.Info(bar)
	inst.innerLogTaskInfo(t)

	defer func() {
		now = time.Now()
		t.Done = true
		t.StoppedAt = now

		vlog.Info("Done.")
		vlog.Info(bar)

	}()

	t.StartedAt = now
	fn := t.Info.Do

	err := fn()

	if err == nil {
		t.OK = true
	} else {
		t.Error = err
	}

	err = inst.innerFilterError(t, err)

	return err
}

func (inst *RunnerV1) innerLogTaskInfo(t *innerTask) {

	const nl = "\n"
	builder := new(strings.Builder)

	builder.WriteString("units.Unit" + nl)
	builder.WriteString("  name: " + t.Info.Name + nl)
	builder.WriteString("  alias: " + t.Info.Alias + nl)
	builder.WriteString("  description: " + t.Info.Description + nl)

	str := builder.String()
	vlog.Info("%s", str)
}

func (inst *RunnerV1) innerGetBarString() string {
	return "////////////////////////////////////////////////////////////////////////////////"
}

func (inst *RunnerV1) innerAcceptTask(t *innerTask) bool {

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

func (inst *RunnerV1) innerFilterError(t *innerTask, err error) error {

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

func (inst *RunnerV1) innerDoLoadTasks(rc *runningContext) error {
	loader := new(innerTaskLoader)
	rawNameList := inst.UnitNameList
	rawUnitList := inst.innerLoadUnitList()
	tasks, err := loader.load(rawUnitList, rawNameList)

	if err == nil {
		rc.tasks = tasks
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
