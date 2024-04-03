package algorithms

type ProcessHandler interface {
	AverageExecutionTime(p []Process)
	AverageWaitingTime(p []Process)
}
