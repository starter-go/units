package main

import (
	"os"
	"testing"

	"github.com/starter-go/units"
	u2 "github.com/starter-go/units/modules/units"
	"github.com/starter-go/vlog"
)

func TestRunFunc(t *testing.T) {

	props := map[string]string{
		"debug.enabled":        "1",
		"debug.log-properties": "1",
		"vlog.level":           "info",
	}

	units.Run(&units.Config{
		Args:       os.Args,
		Cases:      "test-4",
		Module:     u2.ModuleForTest(),
		T:          t,
		Properties: props,
		UsePanic:   false,
	})

	vlog.Debug("done")
}
