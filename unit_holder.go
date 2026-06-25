package units

import (
	"github.com/starter-go/base/lang"
)

type UnitHolder struct {
	Ref *Registration

	Info Registration

	State TaskState

	StartedAt lang.Time

	StoppedAt lang.Time

	Index int

	Error error

	Message string

	Done bool // 表示是否已经执行

	OK bool // 表示是否成功运行

	Selected bool // 表示是否被 namelist 选中

}
