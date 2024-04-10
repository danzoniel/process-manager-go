package algorithms

import (
	"fmt"
)

type Fcfs struct {
	Processes        []Process
	TotalProcessTime int
}

func (f *Fcfs) FirstComeFirtServerd() {
	fmt.Println("\nFIRST COME FIRST SERVED")

	p := make([]Process, len(f.Processes))
	finishedExecutingAt := p[0].ArrivedTime
	actualInstant := 0

	for i := range f.Processes {
		p[i] = f.Processes[i].NewProcess()

		finishedExecutingAt += p[i].ServiceTime

		p[i].ProcessTime.finishedExecutingAt = finishedExecutingAt

		p[i].ProcessTime.startedExecutingAt = actualInstant

		p[i].ProcessTime.totalExecutionTime = p[i].ProcessTime.finishedExecutingAt - p[i].ArrivedTime

		actualInstant = p[i].ProcessTime.finishedExecutingAt
		if i > 0 {
			p[i].ProcessTime.totalWaitingTime = p[i-1].ProcessTime.finishedExecutingAt - p[i].ArrivedTime
		}
	}

	f.TotalProcessTime = p[len(f.Processes)-1].ProcessTime.finishedExecutingAt

	Graph(p)
	CalculateAverageProcessTime(p)
	CalculateAverageWaitTime(p)
}
