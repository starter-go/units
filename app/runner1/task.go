package runner1

import (
	"time"

	"github.com/starter-go/units"
)

type innerTask struct {
	Index int

	Info units.Registration

	Error error

	Done bool // 表示是否已经执行

	OK bool // 表示是否成功运行

	State units.TaskState

	Selected bool // 表示是否被 namelist 选中

	StartedAt time.Time

	StoppedAt time.Time
}
