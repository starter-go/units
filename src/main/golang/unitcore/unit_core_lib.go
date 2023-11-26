package unitcore

import (
	"fmt"
	"strings"

	"github.com/starter-go/application"
)

// UnitCore ...
type UnitCore struct {

	//starter:component
	_as func(application.Lifecycle) //starter:as("#")

	AC      application.Context //starter:inject("context")
	RunAll  bool                //starter:inject("${test.units.run-all}")
	RunList string              //starter:inject("${test.units.run-list}")

}

func (inst *UnitCore) _impl() application.Lifecycle {
	return inst
}

// Life ...
func (inst *UnitCore) Life() *application.Life {
	return &application.Life{OnLoop: inst.run}
}

func (inst *UnitCore) run() error {
	ctx := inst.AC
	holder := &RunnerHolder{}
	name := holder.AttributeName()
	obj, err := ctx.GetAttributes().GetAttributeRequired(name)
	if err != nil {
		return err
	}
	holder = obj.(*RunnerHolder)
	fn := holder.Run
	if fn == nil {
		return fmt.Errorf("func [unitcore.RunnerHolder.Run] is nil")
	}
	return inst.invoke(holder)
}

func (inst *UnitCore) invoke(h *RunnerHolder) error {

	list := strings.Split(inst.RunList, ",")

	ctx := &Context{}
	ctx.ApplicationContext = inst.AC
	ctx.RunAll = inst.RunAll
	ctx.RunList = list

	fn := h.Run
	return fn(ctx)
}
