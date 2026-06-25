package test4units
import (
    p0dc072ed4 "github.com/starter-go/units"
    p370fb552a "github.com/starter-go/units/src/test/golang/example2units"
     "github.com/starter-go/application"
)

// type p370fb552a.Test1 in package:github.com/starter-go/units/src/test/golang/example2units
//
// id:com-370fb552a1007af0-example2units-Test1
// class:
// alias:
// scope:singleton
//
type p370fb552a1_example2units_Test1 struct {
}

func (inst* p370fb552a1_example2units_Test1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-370fb552a1007af0-example2units-Test1"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p370fb552a1_example2units_Test1) new() any {
    return &p370fb552a.Test1{}
}

func (inst* p370fb552a1_example2units_Test1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p370fb552a.Test1)
	nop(ie, com)

	
    com.NameOfUnit1 = inst.getNameOfUnit1(ie)
    com.NameOfUnit2 = inst.getNameOfUnit2(ie)
    com.NameOfUnit3 = inst.getNameOfUnit3(ie)
    com.NameOfUnit4 = inst.getNameOfUnit4(ie)


    return nil
}


func (inst*p370fb552a1_example2units_Test1) getNameOfUnit1(ie application.InjectionExt)string{
    return ie.GetString("${unit.u-1.name}")
}


func (inst*p370fb552a1_example2units_Test1) getNameOfUnit2(ie application.InjectionExt)string{
    return ie.GetString("${unit.u-2.name}")
}


func (inst*p370fb552a1_example2units_Test1) getNameOfUnit3(ie application.InjectionExt)string{
    return ie.GetString("${unit.u-3.name}")
}


func (inst*p370fb552a1_example2units_Test1) getNameOfUnit4(ie application.InjectionExt)string{
    return ie.GetString("${unit.u-4.name}")
}



// type p370fb552a.CaseTryDirMan in package:github.com/starter-go/units/src/test/golang/example2units
//
// id:com-370fb552a1007af0-example2units-CaseTryDirMan
// class:class-0dc072ed44b3563882bff4e657a52e62-Unit
// alias:
// scope:singleton
//
type p370fb552a1_example2units_CaseTryDirMan struct {
}

func (inst* p370fb552a1_example2units_CaseTryDirMan) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-370fb552a1007af0-example2units-CaseTryDirMan"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-Unit"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p370fb552a1_example2units_CaseTryDirMan) new() any {
    return &p370fb552a.CaseTryDirMan{}
}

func (inst* p370fb552a1_example2units_CaseTryDirMan) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p370fb552a.CaseTryDirMan)
	nop(ie, com)

	
    com.DirMan = inst.getDirMan(ie)


    return nil
}


func (inst*p370fb552a1_example2units_CaseTryDirMan) getDirMan(ie application.InjectionExt)p0dc072ed4.DirManager{
    return ie.GetComponent("#alias-0dc072ed44b3563882bff4e657a52e62-DirManager").(p0dc072ed4.DirManager)
}


