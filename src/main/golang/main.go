package main

import (
	"github.com/starter-go/units"

	u2 "github.com/starter-go/units/modules/units"
)

func main() {

	m := u2.ModuleForTest()
	c := units.NewContext()
	r := units.NewRunner()

	c.Module = m

	r.Run(c)

}
