package units

import (
	"github.com/starter-go/starter"
)

type bootRunner struct {

	// deps []application.Module
	// // mb      *application.ModuleBuilder
	// mod     application.Module
	// t       *testing.T
	// props   map[string]string
	// enPanic bool

}

// func (inst *bootRunner) Module(m application.Module) Runner {
// 	inst.mod = m
// 	return inst
// }

// func (inst *bootRunner) SetProperties(t map[string]string) Runner {
// 	inst.props = t
// 	return inst
// }

// func (inst *bootRunner) Testing(t *testing.T) Runner {
// 	inst.t = t
// 	return inst
// }

// func (inst *bootRunner) EnablePanic(enabled bool) Runner {
// 	inst.enPanic = enabled
// 	return inst
// }

func (inst *bootRunner) Run(c *Context) error {

	i := starter.Init(c.Arguments)
	i.MainModule(c.Module)
	i.WithPanic(c.UsePanic)
	return i.Run()

}

func (inst *bootRunner) _impl() Runner {
	return inst
}
