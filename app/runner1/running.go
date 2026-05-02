package runner1

type runningContext struct {
	config    innerRunnerV1config
	tasks     []*innerTask
	abort     bool
	lastError error
}
