package unit

import (
	"fmt"

	"github.com/starter-go/units"
)

// Test1 ...
type Test1 struct {

	//starter:component

}

func (inst *Test1) _impl() units.Units {
	return inst
}

// ListUnits ...
func (inst *Test1) ListUnits(list []*units.Registration) []*units.Registration {

	list = append(list, &units.Registration{
		Enabled:  false,
		Priority: 0,
		Name:     "test-91",
		Test:     inst.test1,
	})
	list = append(list, &units.Registration{
		Enabled:  true,
		Priority: 0,
		Name:     "test-2",
		Test:     inst.test2,
	})
	list = append(list, &units.Registration{
		Enabled:  true,
		Priority: 3,
		Name:     "test-3",
		Test:     inst.test3,
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
