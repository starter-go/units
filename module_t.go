package units

import (
	"embed"

	"github.com/starter-go/application"
)

const (
	theModuleName     = "github.com/starter-go/units"
	theModuleVersion  = "v0.0.5"
	theModuleRevision = 5
)

////////////////////////////////////////////////////////////////////////////////

const (
	theSrcMainResPath = "src/main/resources"
	theSrcTestResPath = "src/test/resources"
)

//go:embed "src/main/resources"
var theSrcMainResFS embed.FS

//go:embed "src/test/resources"
var theSrcTestResFS embed.FS

////////////////////////////////////////////////////////////////////////////////

// ModuleMainT 创建模块 [github.com/starter-go/units]
func ModuleMainT() *application.ModuleBuilder {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theSrcMainResFS, theSrcMainResPath)
	// mb.Depend(starter.Module())
	return mb
}

// ModuleTestT 创建模块 [github.com/starter-go/units]
func ModuleTestT() *application.ModuleBuilder {
	mb := &application.ModuleBuilder{}
	mb.Name(theModuleName + "#test")
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)
	mb.EmbedResources(theSrcTestResFS, theSrcTestResPath)
	// mb.Depend(starter.Module())
	return mb
}
