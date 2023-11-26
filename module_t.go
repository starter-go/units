package units

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/starter"
)

const (
	theMainModuleName     = "github.com/starter-go/units"
	theMainModuleVersion  = "v0.0.1"
	theMainModuleRevision = 1
	theMainModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theMainModuleResFS embed.FS

// ModuleT 创建模块 [github.com/starter-go/units]
func ModuleT() *application.ModuleBuilder {
	mb := &application.ModuleBuilder{}
	mb.Name(theMainModuleName)
	mb.Version(theMainModuleVersion)
	mb.Revision(theMainModuleRevision)
	mb.EmbedResources(theMainModuleResFS, theMainModuleResPath)
	mb.Depend(starter.Module())
	return mb
}
