package algorithms

import (
	"fmt"
	"sort"
)

type RR struct {
	Processes []Process
	Quantum   int
}

func (a *RR) RoundRobin() []Process {
	fmt.Println("\nROUND ROBIN")

	if len(a.Processes) == 0 {
		fmt.Println("Nenhum processo na fila.")
		return a.Processes
	}

	sort.SliceStable(a.Processes, func(i, j int) bool {
		return a.Processes[i].ArrivedTime < a.Processes[j].ArrivedTime
	})

	var currentTime int
	var queue []Process
	for {
		for _, process := range a.Processes {
			if process.ArrivedTime <= currentTime && !contains(queue, process) {
				queue = append(queue, process)
			}
		}

		if len(queue) == 0 {
			break
		}

		process := queue[0]
		queue = queue[1:]

		executionTime := min(a.Quantum, process.ServiceTime)
		process.ServiceTime -= executionTime
		currentTime += executionTime

		if process.ServiceTime == 0 {
			process.WaitTime = currentTime - process.ArrivedTime - process.ServiceTime
		} else {
			queue = append(queue, process)
		}
	}

	return a.Processes
}

func (a *RR) AverageExecutionTime() {
	var totalExecutionTime int
	for _, process := range a.Processes {
		totalExecutionTime += process.ServiceTime
	}

	averageExecutionTime := float32(totalExecutionTime) / float32(len(a.Processes))
	fmt.Printf("\nTempo médio de execução: %.1f s\n", averageExecutionTime)
}

func (a *RR) AverageWaitingTime() {
	var totalWaitTime int
	for _, process := range a.Processes {
		totalWaitTime += process.WaitTime
	}

	averageWaitTime := float32(totalWaitTime) / float32(len(a.Processes))
	fmt.Printf("\nTempo médio de espera: %.1f s\n", averageWaitTime)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func contains(slice []Process, process Process) bool {
	for _, p := range slice {
		if p.ProcessId == process.ProcessId {
			return true
		}
	}
	return false
}
