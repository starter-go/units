package units

import (
	"context"
	"os"
	"testing"

	"github.com/starter-go/application"
)

type Context struct {
	CC context.Context

	Module application.Module

	T *testing.T

	UsePanic bool

	Properties map[string]string

	Selector string // 用于选择需要测试的单元, 包括(ID选择器'#'; Class选择器'.'; 全部选择器'*')

	Arguments []string
}

////////////////////////////////////////////////////////////////////////////////

// Runner 单元测试执行器
type Runner interface {

	// Dependencies(deps ...application.Module) Runner
	// ModuleT(mb *application.ModuleBuilder) Runner
	// Module() Runner
	// Testing() Runner
	// EnablePanic() Runner

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
