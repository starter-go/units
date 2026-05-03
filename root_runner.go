package units

import (
	"github.com/starter-go/starter"
)

type RootRunner struct {
}

func (inst *RootRunner) Run(c *Context) error {

	i := starter.Init(c.Arguments)
	i.MainModule(c.Module)
	i.WithPanic(c.UsePanic)
	err := i.Run()

	if c.UsePanic && (err != nil) {
		panic(err)
	}

	return err
}

func (inst *RootRunner) _impl() Runner {
	return inst
}
