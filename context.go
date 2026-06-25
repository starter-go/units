package units

import (
	"context"
	"fmt"
	"testing"

	"github.com/starter-go/application"
	"github.com/starter-go/base/lang"
)

////////////////////////////////////////////////////////////////////////////////

type TestContext struct {
	CC context.Context

	Module application.Module

	T *testing.T

	UsePanic bool

	Properties map[string]string

	Selector string // 用于选择需要测试的单元, 包括(ID选择器'#'; Class选择器'.'; 全部选择器'*')

	StartedAt lang.Time

	StoppedAt lang.Time

	Units []*UnitHolder

	Arguments []string
}

type UnitContext struct {
	Parent *Context

	Properties map[string]string

	StartedAt lang.Time

	StoppedAt lang.Time

	Unit *Registration // 当前执行的测试单元

}

type Context = TestContext

////////////////////////////////////////////////////////////////////////////////

type ContextHolder struct {
	CC context.Context
	TC *TestContext
	UC *UnitContext
}

////////////////////////////////////////////////////////////////////////////////

func GetContext(c context.Context) (*TestContext, error) {

	holder, err := GetContextHolder(c)
	if err != nil {
		return nil, err
	}

	tc := holder.TC
	if tc == nil {
		return nil, fmt.Errorf("units.Context is nil")
	}

	return tc, nil
}

func GetUnitContext(c context.Context) (*UnitContext, error) {

	holder, err := GetContextHolder(c)
	if err != nil {
		return nil, err
	}

	uc := holder.UC
	if uc == nil {
		return nil, fmt.Errorf("units.UnitContext is nil")
	}

	return uc, nil
}

////////////////////////////////////////////////////////////////////////////////

const theContextHolderBindingKey = "bind://github.com/starter-go/units.ContextHolder"

func GetContextHolder(c context.Context) (*ContextHolder, error) {

	const key = theContextHolderBindingKey
	obj := c.Value(key)
	ch, ok := obj.(*ContextHolder)

	if ok && (ch != nil) {
		return ch, nil
	}

	return nil, fmt.Errorf("units.GetContextHolder: no ContextHolder, use 'SetupContextHolder()' to setup")
}

func SetupContextHolder(c context.Context) (*ContextHolder, error) {

	const key = theContextHolderBindingKey
	obj := c.Value(key)
	ch, ok := obj.(*ContextHolder)

	if ok && (ch != nil) {
		return ch, nil
	}

	holder := new(ContextHolder)
	c2 := context.WithValue(c, key, holder)

	holder.CC = c2
	return holder, nil
}

////////////////////////////////////////////////////////////////////////////////
