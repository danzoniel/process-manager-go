package algorithms

import (
	"fmt"
	"sort"
)

type Fcfs struct {
	Processes []Process
	ProcessHandler
}

func (a *Fcfs) FirstComeFirtServerd() []Process {
	fmt.Println("\nFIRST COME FIRST SERVED")

	processessQuantity := len(a.Processes)

	if processessQuantity <= 1 {
		return a.Processes
	}

	sort.SliceStable(a.Processes, func(i, j int) bool {
		return a.Processes[i].ArrivedTime < a.Processes[j].ArrivedTime
	})
	return a.Processes
}

func (a *Fcfs) AverageExecutionTime() {
	var res int
	var finalRes float32
	var i int

	fmt.Print("\nTempo médio de execução\n")
	fmt.Print("(")

	for i = range a.Processes {
		res = int(a.Processes[i].ServiceTime) + res
		resMinusArrivedTime := res - int(a.Processes[i].ArrivedTime)
		finalRes = float32(finalRes) + float32(resMinusArrivedTime)

		fmt.Printf("(%d - %d)", res, a.Processes[i].ArrivedTime)
		if i < len(a.Processes)-1 {
			fmt.Print(" + ")
		}
	}

	processesLength := float32(len(a.Processes))

	fmt.Print(")")
	fmt.Printf(" / %d", len(a.Processes))
	fmt.Printf(" = %.1f s", finalRes/processesLength)

}

func (a *Fcfs) AverageWaitingTime() {
	var totalWaitTime uint

	fmt.Print("\nTempo médio de espera\n")
	fmt.Print("(")

	for i := range a.Processes {
		if i != 0 {
			totalWaitTime += a.Processes[i-1].ServiceTime
		}
		totalWaitTime += a.Processes[i].ArrivedTime

		fmt.Printf("(%d + %d)", totalWaitTime, a.Processes[i].ArrivedTime)
		if i < len(a.Processes)-1 {
			fmt.Print(" + ")
		}
	}

	processesLength := float32(len(a.Processes))

	fmt.Print(")")
	fmt.Printf(" / %d", len(a.Processes))
	fmt.Printf(" = %.1f s", float32(totalWaitTime)/processesLength)
}
