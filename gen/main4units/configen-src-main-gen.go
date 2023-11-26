package main4units
import (
    p0ef6f2938 "github.com/starter-go/application"
    p0b30545e5 "github.com/starter-go/units/src/main/golang/unitcore"
     "github.com/starter-go/application"
)

// type p0b30545e5.UnitCore in package:github.com/starter-go/units/src/main/golang/unitcore
//
// id:com-0b30545e5e1b101e-unitcore-UnitCore
// class:
// alias:alias-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle
// scope:singleton
//
type p0b30545e5e_unitcore_UnitCore struct {
}

func (inst* p0b30545e5e_unitcore_UnitCore) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-0b30545e5e1b101e-unitcore-UnitCore"
	r.Classes = ""
	r.Aliases = "alias-0ef6f2938681e99da4b0c19ce3d3fb4f-Lifecycle"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p0b30545e5e_unitcore_UnitCore) new() any {
    return &p0b30545e5.UnitCore{}
}

func (inst* p0b30545e5e_unitcore_UnitCore) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p0b30545e5.UnitCore)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)
    com.RunAll = inst.getRunAll(ie)
    com.RunList = inst.getRunList(ie)


    return nil
}


func (inst*p0b30545e5e_unitcore_UnitCore) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p0b30545e5e_unitcore_UnitCore) getRunAll(ie application.InjectionExt)bool{
    return ie.GetBool("${test.units.run.all}")
}


func (inst*p0b30545e5e_unitcore_UnitCore) getRunList(ie application.InjectionExt)string{
    return ie.GetString("${test.units.run.list}")
}


