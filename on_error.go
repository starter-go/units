package units

import "strings"

const (
	OnErrorNone   OnErrorMethod = "none"
	OnErrorMute   OnErrorMethod = "mute"
	OnErrorWarn   OnErrorMethod = "warn"
	OnErrorPanic  OnErrorMethod = "panic"
	OnErrorError  OnErrorMethod = "error"
	OnErrorAbort  OnErrorMethod = "abort"
	OnErrorIgnore OnErrorMethod = "ignore"

	OnErrorDefault = OnErrorError
)

////////////////////////////////////////////////////////////////////////////////

type OnErrorMethod string

func (oem OnErrorMethod) Normalize() OnErrorMethod {
	str := oem.String()
	str = strings.TrimSpace(str)
	str = strings.ToLower(str)
	if str == "" {
		return OnErrorDefault
	}
	return OnErrorMethod(str)
}

func (oem OnErrorMethod) String() string {
	return string(oem)
}

////////////////////////////////////////////////////////////////////////////////
