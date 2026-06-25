package units

import (
	"context"
	"os"
)

////////////////////////////////////////////////////////////////////////////////

// Runner 单元(set)测试执行器
type Runner interface {
	Run(c *Context) error
}

////////////////////////////////////////////////////////////////////////////////

// NewRunner 新建一个 Runner
func NewRunner() Runner {
	return &RootRunner{}
}

// NewRunner 新建一个 units.Context
func NewContext() *Context {

	ctx := new(Context)

	ctx.Arguments = os.Args
	ctx.CC = context.Background()
	ctx.UsePanic = true
	ctx.Properties = make(map[string]string)

	return ctx
}

func Run(c *Context) error {
	r := NewRunner()
	return r.Run(c)
}

////////////////////////////////////////////////////////////////////////////////

type RunnerRegistration struct {
	Name    string
	Alias   string
	Enabled bool
	Runner  Runner
}

////////////////////////////////////////////////////////////////////////////////

type RunnerRegistry interface {
	GetRegistration(ret *RunnerRegistration) error
}

////////////////////////////////////////////////////////////////////////////////
