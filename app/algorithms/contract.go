package algorithms

type AverageProcessTimer interface {
	AverageExecutionTimer
	AverageWaitingTimer
}

type AverageExecutionTimer interface {
	AverageExecutionTime()
}

type AverageWaitingTimer interface {
	AverageWaitingTime()
}
