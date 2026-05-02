package units

type OnErrorMethod string

const (
	OnErrorNone   OnErrorMethod = "none"
	OnErrorMute   OnErrorMethod = "mute"
	OnErrorWarn   OnErrorMethod = "warn"
	OnErrorPanic  OnErrorMethod = "panic"
	OnErrorError  OnErrorMethod = "error"
	OnErrorAbort  OnErrorMethod = "abort"
	OnErrorIgnore OnErrorMethod = "ignore"
)
