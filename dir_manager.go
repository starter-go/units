package units

import (
	"context"

	"github.com/starter-go/afs"
)

////////////////////////////////////////////////////////////////////////////////

type DirKey string

const (
	DirKeyTemp   DirKey = "temp"
	DirKeyData   DirKey = "data"
	DirKeyConfig DirKey = "config"
	DirKeyInput  DirKey = "input"
	DirKeyOutput DirKey = "output"
)

////////////////////////////////////////////////////////////////////////////////

type DirScope string

const (
	DirScopeRuntime DirScope = "runtime"
	DirScopeUnit    DirScope = "unit"
	DirScopeModule  DirScope = "module"
)

////////////////////////////////////////////////////////////////////////////////

type DirHolder struct {
	Context context.Context

	Path afs.Path

	Key DirKey

	Scope DirScope

	Unit *Registration
}

type DirManager interface {
	GetDir(holder *DirHolder) (*DirHolder, error)
}

////////////////////////////////////////////////////////////////////////////////
