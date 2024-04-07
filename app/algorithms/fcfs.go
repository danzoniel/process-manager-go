package algorithms

import (
	"fmt"
)

type Fcfs struct {
	Processes []Process
}

func (f Fcfs) FirstComeFirtServerd() []Process {
	fmt.Println("\nFIRST COME FIRST SERVED")

	p := make([]Process, len(f.Processes))
	finishedExecutingAt := p[0].ArrivedTime

	for i := range f.Processes {
		p[i] = f.Processes[i].NewProcess()

		// totalExecutionTime(finishedExecutingAt, &p[i])
		finishedExecutingAt += p[i].ServiceTime

		p[i].ProcessTime.finishedExecutingAt = finishedExecutingAt

		p[i].ProcessTime.totalExecutionTime = p[i].ProcessTime.finishedExecutingAt - p[i].ArrivedTime

		if i > 0 {
			p[i].ProcessTime.totalWaitingTime = p[i-1].ProcessTime.finishedExecutingAt - p[i].ArrivedTime
		}

		fmt.Println("total execution time for process ", p[i].ProcessId, p[i].ProcessTime.finishedExecutingAt)
		fmt.Println("total waiting time for process ", p[i].ProcessId, p[i].ProcessTime.totalWaitingTime)
	}

	fmt.Println(p)

	return nil
}

// func totalExecutionTime(finishedExecutingAt int, p *Process) {
// 	finishedExecutingAt += p.ServiceTime

// 	p.ProcessTime.finishedExecutingAt = finishedExecutingAt

// 	p.ProcessTime.totalExecutionTime = p.ProcessTime.finishedExecutingAt - p.ArrivedTime
// }

// func executeProcess(process *Process) *Process {
// 	fmt.Println("arrived time:", process.ArrivedTime)

// 	p := process.ProcessTime
// 	p.totalExecutionTime = process.ArrivedTime
// 	fmt.p
// 	return &Process{
// 		ProcessTime: p,
// 	}
// }
