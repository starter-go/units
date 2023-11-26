package test4units
import (
    pff286ca77 "github.com/starter-go/units/src/test/golang/unit"
     "github.com/starter-go/application"
)

// type pff286ca77.Test1 in package:github.com/starter-go/units/src/test/golang/unit
//
// id:com-ff286ca7719b7ca3-unit-Test1
// class:
// alias:
// scope:singleton
//
type pff286ca771_unit_Test1 struct {
}

func (inst* pff286ca771_unit_Test1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ff286ca7719b7ca3-unit-Test1"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pff286ca771_unit_Test1) new() any {
    return &pff286ca77.Test1{}
}

func (inst* pff286ca771_unit_Test1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pff286ca77.Test1)
	nop(ie, com)

	


    return nil
}


