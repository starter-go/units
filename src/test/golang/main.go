package main

import (
	"os"

	"github.com/starter-go/units"
	"github.com/starter-go/units/gen/test4units"
)

func main() {

	m := units.ModuleT()
	m.Components(test4units.ExportConfig)

	r := units.NewRunner()
	r.ModuleT(m)
	r.EnablePanic(false)
	r.Run(os.Args)
}
