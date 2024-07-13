package main

import (
	"os"

	"github.com/starter-go/units"
	u2 "github.com/starter-go/units/modules/units"
)

func main() {

	m := u2.ModuleForTest()
	// m.Components(test4units.ExportConfig)

	r := units.NewRunner()
	r.Module(m)
	r.EnablePanic(false)
	r.Run(os.Args)
}
