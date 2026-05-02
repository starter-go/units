package main

import (
	"testing"

	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

func TestRunFunc(t *testing.T) {

	props := map[string]string{
		"debug.enabled":        "1",
		"debug.log-properties": "1",
		"vlog.level":           "info",
	}

	ctx := units.NewContext()

	ctx.Module = nil
	ctx.T = t
	ctx.Properties = props

	units.Run(ctx)
	vlog.Debug("done")
}
