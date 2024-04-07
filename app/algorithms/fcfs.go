package algorithms

import (
	"fmt"
)

type Fcfs struct {
	Processes        []Process
	TotalProcessTime int
}

func (f *Fcfs) FirstComeFirtServerd() []Process {
	fmt.Println("\nFIRST COME FIRST SERVED")

	p := make([]Process, len(f.Processes))
	finishedExecutingAt := p[0].ArrivedTime

	for i := range f.Processes {
		p[i] = f.Processes[i].NewProcess()

		finishedExecutingAt += p[i].ServiceTime

		p[i].ProcessTime.finishedExecutingAt = finishedExecutingAt

		p[i].ProcessTime.totalExecutionTime = p[i].ProcessTime.finishedExecutingAt - p[i].ArrivedTime

		if i > 0 {
			p[i].ProcessTime.totalWaitingTime = p[i-1].ProcessTime.finishedExecutingAt - p[i].ArrivedTime
		}
	}

	f.TotalProcessTime = p[len(f.Processes)-1].ProcessTime.finishedExecutingAt

	CalculateAverageProcessTime(p)
	CalculateAverageWaitTime(p)

	fmt.Println(p)
	return nil
}

func (f *Fcfs) PrintTable() {

	fmt.Println("\nGráfico First Come, First Served\n")

	maxYScale := len(f.Processes) * 2
	maxXScale := len(f.Processes) * 18

	scale := NewAxisScale(0, maxXScale, 0, maxYScale)

	DrawYAxis(scale, f.Processes)

	DrawXAxis(scale, f.TotalProcessTime, len(f.Processes))

}
