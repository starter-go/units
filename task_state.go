package units

type TaskState string

const (
	TaskStateInit    TaskState = "init"
	TaskStateOK      TaskState = "ok"
	TaskStateError   TaskState = "error"
	TaskStatePanic   TaskState = "panic"
	TaskStateAbort   TaskState = "abort"
	TaskStateIgnored TaskState = "ignored"
	TaskStateRunning TaskState = "running"
)
