package main4units
import (
    p0d2a11d16 "github.com/starter-go/afs"
    p0ef6f2938 "github.com/starter-go/application"
    p0dc072ed4 "github.com/starter-go/units"
    pfe207c121 "github.com/starter-go/units/app/boot"
    p18376513c "github.com/starter-go/units/app/dirmanager"
    pef802692c "github.com/starter-go/units/app/runner1"
     "github.com/starter-go/application"
)

// type pfe207c121.Boot in package:github.com/starter-go/units/app/boot
//
// id:com-fe207c12166f6afd-boot-Boot
// class:
// alias:
// scope:singleton
//
type pfe207c1216_boot_Boot struct {
}

func (inst* pfe207c1216_boot_Boot) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-fe207c12166f6afd-boot-Boot"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pfe207c1216_boot_Boot) new() any {
    return &pfe207c121.Boot{}
}

func (inst* pfe207c1216_boot_Boot) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pfe207c121.Boot)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)
    com.RunnerList = inst.getRunnerList(ie)
    com.RunnerAlias = inst.getRunnerAlias(ie)


    return nil
}


func (inst*pfe207c1216_boot_Boot) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*pfe207c1216_boot_Boot) getRunnerList(ie application.InjectionExt)[]p0dc072ed4.RunnerRegistry{
    dst := make([]p0dc072ed4.RunnerRegistry, 0)
    src := ie.ListComponents(".class-0dc072ed44b3563882bff4e657a52e62-RunnerRegistry")
    for _, item1 := range src {
        item2 := item1.(p0dc072ed4.RunnerRegistry)
        dst = append(dst, item2)
    }
    return dst
}


func (inst*pfe207c1216_boot_Boot) getRunnerAlias(ie application.InjectionExt)string{
    return ie.GetString("${units.runner}")
}



// type p18376513c.DirManagerImpl in package:github.com/starter-go/units/app/dirmanager
//
// id:com-18376513c555e28e-dirmanager-DirManagerImpl
// class:
// alias:alias-0dc072ed44b3563882bff4e657a52e62-DirManager
// scope:singleton
//
type p18376513c5_dirmanager_DirManagerImpl struct {
}

func (inst* p18376513c5_dirmanager_DirManagerImpl) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-18376513c555e28e-dirmanager-DirManagerImpl"
	r.Classes = ""
	r.Aliases = "alias-0dc072ed44b3563882bff4e657a52e62-DirManager"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p18376513c5_dirmanager_DirManagerImpl) new() any {
    return &p18376513c.DirManagerImpl{}
}

func (inst* p18376513c5_dirmanager_DirManagerImpl) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p18376513c.DirManagerImpl)
	nop(ie, com)

	
    com.AC = inst.getAC(ie)
    com.AFS = inst.getAFS(ie)
    com.TheUnitsWorkingDir = inst.getTheUnitsWorkingDir(ie)


    return nil
}


func (inst*p18376513c5_dirmanager_DirManagerImpl) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*p18376513c5_dirmanager_DirManagerImpl) getAFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}


func (inst*p18376513c5_dirmanager_DirManagerImpl) getTheUnitsWorkingDir(ie application.InjectionExt)string{
    return ie.GetString("${units.working.dir}")
}



// type pef802692c.RunnerV1 in package:github.com/starter-go/units/app/runner1
//
// id:com-ef802692c2592b59-runner1-RunnerV1
// class:class-0dc072ed44b3563882bff4e657a52e62-RunnerRegistry
// alias:
// scope:singleton
//
type pef802692c2_runner1_RunnerV1 struct {
}

func (inst* pef802692c2_runner1_RunnerV1) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-ef802692c2592b59-runner1-RunnerV1"
	r.Classes = "class-0dc072ed44b3563882bff4e657a52e62-RunnerRegistry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pef802692c2_runner1_RunnerV1) new() any {
    return &pef802692c.RunnerV1{}
}

func (inst* pef802692c2_runner1_RunnerV1) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pef802692c.RunnerV1)
	nop(ie, com)

	
    com.UnitNameList = inst.getUnitNameList(ie)
    com.AC = inst.getAC(ie)
    com.MyName = inst.getMyName(ie)
    com.MyAlias = inst.getMyAlias(ie)
    com.Enabled = inst.getEnabled(ie)


    return nil
}


func (inst*pef802692c2_runner1_RunnerV1) getUnitNameList(ie application.InjectionExt)string{
    return ie.GetString("${units.list}")
}


func (inst*pef802692c2_runner1_RunnerV1) getAC(ie application.InjectionExt)p0ef6f2938.Context{
    return ie.GetContext()
}


func (inst*pef802692c2_runner1_RunnerV1) getMyName(ie application.InjectionExt)string{
    return ie.GetString("${runner.runner1.name}")
}


func (inst*pef802692c2_runner1_RunnerV1) getMyAlias(ie application.InjectionExt)string{
    return ie.GetString("${runner.runner1.alias}")
}


func (inst*pef802692c2_runner1_RunnerV1) getEnabled(ie application.InjectionExt)bool{
    return ie.GetBool("${runner.runner1.enabled}")
}


