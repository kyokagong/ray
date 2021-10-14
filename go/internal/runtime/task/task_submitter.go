package task

type TaskSubmitter interface {
	SubmitTask() string
}

type LocalTaskSubmitter struct {
	taskExecutor TaskExecutor
	// waitingTasks Map[api.ObjectRef]TaskSpec
}

func (localTaskSubmitter LocalTaskSubmitter) SubmitTask() string {
	return "test"
}

func NewLocalTaskSubmitter() LocalTaskSubmitter {
	return LocalTaskSubmitter{TaskExecutor{}}
}
