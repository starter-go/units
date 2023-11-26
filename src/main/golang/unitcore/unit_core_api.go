package unitcore

// RunnerHolder ...
type RunnerHolder struct {
	Run func(ctx *Context) error
}

// AttributeName ...
func (inst *RunnerHolder) AttributeName() string {
	return "uri:attr:unitcore.RunnerHolder#binding"
}
