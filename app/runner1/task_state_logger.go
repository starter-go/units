package runner1

import (
	"strconv"
	"strings"

	"github.com/starter-go/vlog"
)

type taskStateLogger struct {
	buffer strings.Builder
}

func (inst *taskStateLogger) Log(tasks []*innerTask) {

	const bar = "--------------------------------------------------------------------------------"

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

func (inst *taskStateLogger) innerLogRow(task *innerTask, asHeader bool) {

	inst.innerLogFieldIndex("Index", task, asHeader)
	inst.innerLogFieldNameAlias("-", task, asHeader)

	inst.innerLogFieldState("State", task, asHeader)
	inst.innerLogFieldIsDone("Done", task, asHeader)
	inst.innerLogFieldIsOK("OK", task, asHeader)

	inst.innerLogFieldOnErrMethod("OnErrFn", task, asHeader)
	inst.innerLogFieldError("Error", task, asHeader)

	inst.buffer.WriteString("\n")
}

func (inst *taskStateLogger) innerLogFieldIndex(name string, task *innerTask, asHeader bool) {

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

func (inst *taskStateLogger) innerLogFieldState(name string, task *innerTask, asHeader bool) {

	text := ""
	width := 8

	if asHeader {
		text = name
	} else {
		state := task.State
		text = string(state)
	}

	inst.innerWriteStringWithWidth(width, text)
}

func (inst *taskStateLogger) innerLogFieldIsDone(name string, task *innerTask, asHeader bool) {

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

func (inst *taskStateLogger) innerLogFieldIsOK(name string, task *innerTask, asHeader bool) {

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

func (inst *taskStateLogger) innerLogFieldOnErrMethod(name string, task *innerTask, asHeader bool) {

	text := ""
	width := 8

	if asHeader {
		text = name
	} else {
		text = string(task.Info.OnError)
	}

	inst.innerWriteStringWithWidth(width, text)
}

func (inst *taskStateLogger) innerLogFieldNameAlias(name string, task *innerTask, asHeader bool) {

	text := ""
	width := 20

	if asHeader {
		text = "Name(Alias)"
	} else {
		n1 := task.Info.Name
		n2 := task.Info.Alias
		if n2 == "" {
			text = n1
		} else {
			text = n1 + "(" + n2 + ")"
		}
	}

	inst.innerWriteStringWithWidth(width, text)

}

func (inst *taskStateLogger) innerLogFieldError(name string, task *innerTask, asHeader bool) {

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
