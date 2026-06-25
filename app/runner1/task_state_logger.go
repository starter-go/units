package runner1

import (
	"strconv"
	"strings"
	"time"

	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

type taskStateLogger struct {
	buffer strings.Builder
}

func (inst *taskStateLogger) Log(tasks []*units.UnitHolder) {

	bar := inst.innerGetHorzBar(128)

	inst.buffer.WriteString("\n")
	inst.innerLogRow(nil, true)
	inst.buffer.WriteString(bar + "\n")

	for idx, t := range tasks {
		t.Index = idx
		inst.innerLogRow(t, false)
	}

	str := inst.buffer.String()
	vlog.Info("%s", str)
}

func (inst *taskStateLogger) innerGetHorzBar(width int) string {
	const part = "--------"
	builder := new(strings.Builder)
	count := 0
	for count < width {
		count += len(part)
		builder.WriteString(part)
	}
	return builder.String()
}

func (inst *taskStateLogger) innerLogRow(task *units.UnitHolder, asHeader bool) {

	inst.innerLogFieldIndex("Index", task, asHeader)
	inst.innerLogFieldIDandName("ID", "Name", task, asHeader)

	inst.innerLogFieldState("State", task, asHeader)
	inst.innerLogFieldIsSelected("Selected", task, asHeader)
	inst.innerLogFieldIsDone("Done", task, asHeader)
	inst.innerLogFieldIsOK("OK", task, asHeader)

	inst.innerLogFieldPriority("Priority", task, asHeader)
	inst.innerLogFieldStartedAt("StartedAt", task, asHeader)
	inst.innerLogFieldTimeSpan("Span", task, asHeader)

	inst.innerLogFieldOnErrMethod("OnErrFn", task, asHeader)
	inst.innerLogFieldError("Error", task, asHeader)

	inst.buffer.WriteString("\n")
}

func (inst *taskStateLogger) innerLogFieldIndex(name string, task *units.UnitHolder, asHeader bool) {

	text := ""
	width := 6

	if asHeader {
		text = name
	} else {
		index := task.Index
		text = strconv.Itoa(index)
	}

	inst.innerWriteStringWithWidth(width, text)
}

func (inst *taskStateLogger) innerLogFieldState(name string, task *units.UnitHolder, asHeader bool) {

	text := ""
	width := 10

	if asHeader {
		text = name
	} else {
		state := task.State
		text = string(state)
	}

	inst.innerWriteStringWithWidth(width, text)
}

func (inst *taskStateLogger) innerLogFieldIsDone(name string, task *units.UnitHolder, asHeader bool) {

	text := ""
	width := 6

	if asHeader {
		text = name
	} else {
		if task.Done {
			text = "Yes"
		} else {
			text = "No"
		}
	}

	inst.innerWriteStringWithWidth(width, text)
}

func (inst *taskStateLogger) innerLogFieldIsOK(name string, task *units.UnitHolder, asHeader bool) {

	text := ""
	width := 6

	if asHeader {
		text = name
	} else {
		if task.OK {
			text = "Yes"
		} else {
			text = "No"
		}
	}

	inst.innerWriteStringWithWidth(width, text)
}

func (inst *taskStateLogger) innerLogFieldIsSelected(name string, task *units.UnitHolder, asHeader bool) {

	text := ""
	width := 10

	if asHeader {
		text = name
	} else {
		if task.Selected {
			text = "Yes"
		} else {
			text = "No"
		}
	}

	inst.innerWriteStringWithWidth(width, text)
}

func (inst *taskStateLogger) innerLogFieldOnErrMethod(name string, task *units.UnitHolder, asHeader bool) {

	text := ""
	width := 10

	if asHeader {
		text = name
	} else {
		text = string(task.Info.OnError)
	}

	inst.innerWriteStringWithWidth(width, text)
}

func (inst *taskStateLogger) innerLogFieldPriority(name string, task *units.UnitHolder, asHeader bool) {

	// log:优先级

	text := ""
	width := 10

	if asHeader {
		text = name
	} else {
		p := task.Info.Priority
		text = strconv.Itoa(p)
	}

	inst.innerWriteStringWithWidth(width, text)
}

func (inst *taskStateLogger) innerLogFieldStartedAt(name string, task *units.UnitHolder, asHeader bool) {

	// log:开始时间戳

	text := ""
	width := 28

	if asHeader {
		text = name
	} else {
		t0 := task.StartedAt.Time()
		text = t0.Format(time.DateTime)
	}

	inst.innerWriteStringWithWidth(width, text)
}

func (inst *taskStateLogger) innerLogFieldTimeSpan(name string, task *units.UnitHolder, asHeader bool) {

	// log:耗时

	text := ""
	width := 15

	if asHeader {
		text = name
	} else {
		t0 := task.StartedAt
		t1 := task.StoppedAt
		span := t1.Sub(t0)
		text = span.String()
	}

	inst.innerWriteStringWithWidth(width, text)
}

func (inst *taskStateLogger) innerLogFieldIDandName(id, name string, task *units.UnitHolder, asHeader bool) {
	text := ""
	width := 50
	if asHeader {
		text = id + "/" + name
	} else {
		str1 := strings.TrimSpace(task.Info.ID)
		str2 := strings.TrimSpace(task.Info.Name)
		if str2 == str1 {
			str2 = ""
		}
		text = str1 + "/" + str2
	}
	inst.innerWriteStringWithWidth(width, text)
}

// func (inst *taskStateLogger) innerLogFieldID(name string, task *units.UnitHolder, asHeader bool) {
// 	text := ""
// 	width := 16
// 	if asHeader {
// 		text = name
// 	} else {
// 		text = task.Info.ID
// 	}
// 	inst.innerWriteStringWithWidth(width, text)
// }

func (inst *taskStateLogger) innerLogFieldError(name string, task *units.UnitHolder, asHeader bool) {

	text := ""
	width := 32

	if asHeader {
		text = name
	} else {
		err := task.Error
		if err != nil {
			text = err.Error()
		}
	}

	inst.innerWriteStringWithWidth(width, text)
}

func (inst *taskStateLogger) innerWriteStringWithWidth(width int, str string) {

	count := len(str)
	inst.buffer.WriteString(str)

	for ; count < width; count++ {
		inst.buffer.WriteRune(' ')
	}
}
