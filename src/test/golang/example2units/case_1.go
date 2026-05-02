package example2units

import (
	"fmt"

	"github.com/starter-go/units"
)

// Test1 ...
type Test1 struct {

	//starter:component

	NameOfUnit1 string //starter:inject("${unit.u-1.name}")
	NameOfUnit2 string //starter:inject("${unit.u-2.name}")
	NameOfUnit3 string //starter:inject("${unit.u-3.name}")
	NameOfUnit4 string //starter:inject("${unit.u-4.name}")

}

func (inst *Test1) _impl() units.Unit {
	return inst
}

// Units ...
func (inst *Test1) ListRegistrations(list []*units.Registration) []*units.Registration {

	list = append(list, &units.Registration{
		Enabled:  false,
		Priority: 0,
		Name:     inst.NameOfUnit1,
		OnError:  units.OnErrorAbort,
		Do:       inst.test1,
	})
	list = append(list, &units.Registration{
		Enabled:  true,
		Priority: 0,
		Name:     inst.NameOfUnit2,
		OnError:  units.OnErrorWarn,
		Do:       inst.test2,
	})
	list = append(list, &units.Registration{
		Enabled:  true,
		Priority: 3,
		Name:     inst.NameOfUnit3,
		Do:       inst.test3,
	})
	list = append(list, &units.Registration{
		Enabled:  true,
		Priority: 0,
		Name:     inst.NameOfUnit4,
		Do:       inst.test4,
	})

	return list
}

func (inst *Test1) test1() error {
	return nil
}

func (inst *Test1) test2() error {
	return fmt.Errorf("test2: error")
}

func (inst *Test1) test3() error {
	panic("test3: panic")
	// return nil
}

func (inst *Test1) test4() error {
	return nil
}
