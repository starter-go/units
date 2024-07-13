package main

import (
	"os"
	"testing"

	"github.com/starter-go/units"
	u2 "github.com/starter-go/units/modules/units"
)

func TestRunFunc(t *testing.T) {

	props := map[string]string{
		"debug.enabled":        "1",
		"debug.log-properties": "1",
	}

	units.Run(&units.Config{
		Args:       os.Args,
		Cases:      "test-4",
		Module:     u2.ModuleForTest(),
		T:          t,
		Properties: props,
		UsePanic:   false,
	})

}
