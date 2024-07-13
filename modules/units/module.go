package units

import (
	"github.com/starter-go/application"
	"github.com/starter-go/starter"
	"github.com/starter-go/units"
	"github.com/starter-go/units/gen/main4units"
	"github.com/starter-go/units/gen/test4units"
)

// Module  ...
func Module() application.Module {
	mb := units.ModuleMainT()
	mb.Components(main4units.ExportConfig)
	mb.Depend(starter.Module())
	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {
	mb := units.ModuleTestT()
	mb.Components(test4units.ExportConfig)
	mb.Depend(Module())
	return mb.Create()
}
