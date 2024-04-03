package algorithms

import (
	"fmt"
	"sort"
)

type Fcfs struct {
	Processes []Process
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

	for i = range a.Processes {
		res += a.Processes[i].ServiceTime
		resMinusArrivedTime := res - a.Processes[i].ArrivedTime
		finalRes = float32(finalRes) + float32(resMinusArrivedTime)

		if i > 0 {
			fmt.Print(" + ")
		}
		fmt.Printf("Processo %s: %d - %d = %d", a.Processes[i].ProcessId, res, a.Processes[i].ArrivedTime, resMinusArrivedTime)
	}

	processesLength := float32(len(a.Processes))
	averageWaitTime := finalRes / processesLength

	fmt.Printf(" / %d = %.1f s\n", len(a.Processes), averageWaitTime)
}

func (a *Fcfs) AverageWaitingTime() {
	fmt.Println("\n\nTempo médio de espera:")

	if len(a.Processes) == 0 {
		fmt.Println("Nenhum processo na fila.")
		return
	}

	var totalWaitTime int
	var finalRes float32

	fmt.Printf("Processo %s: %d - %d = %d", a.Processes[0].ProcessId, 0, 0, 0)
	totalWaitTime += 0
	finalRes += 0

	for i := 0; i < len(a.Processes); i++ {

		if i == 0 {

		} else {

			waitTime := totalWaitTime + a.Processes[i-1].ServiceTime - a.Processes[i].ArrivedTime
			totalWaitTime += a.Processes[i-1].ServiceTime
			finalRes += float32(waitTime)
			fmt.Printf(" + Processo %s: %d - %d = %d", a.Processes[i].ProcessId, totalWaitTime, a.Processes[i].ArrivedTime, waitTime)
		}

	}

	processesLength := float32(len(a.Processes))
	averageWaitTime := finalRes / processesLength

	fmt.Printf(" / %d = %.1f s\n", len(a.Processes), averageWaitTime)
}
