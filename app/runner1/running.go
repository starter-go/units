package runner1

import (
	"context"

	"github.com/starter-go/units"
)

type runningContext struct {
	config    innerRunnerV1config
	tasks     []*units.UnitHolder
	abort     bool
	lastError error
	cc        context.Context
	tc        *units.TestContext
	uc        *units.UnitContext
}
