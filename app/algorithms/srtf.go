package algorithms

import (
	"fmt"
)

type SRTF struct {
	Processes []Process
}

func (s *SRTF) ShortestRemainingTimeFirst() []Process {
	fmt.Println("\nSHORTEST REMAINING TIME FIRST")

	if len(s.Processes) == 0 {
		fmt.Println("Nenhum processo na fila.")
		return s.Processes
	}

	var currentTime int
	var queue []Process
	var executedProcesses []Process

	for len(s.Processes) > 0 || len(queue) > 0 {
		for i := 0; i < len(s.Processes); i++ {
			if s.Processes[i].ArrivedTime <= currentTime {
				queue = append(queue, s.Processes[i])
				s.Processes = append(s.Processes[:i], s.Processes[i+1:]...)
				i--
			}
		}

		if len(queue) == 0 {
			currentTime++
			continue
		}

		shortestIndex := 0
		for i := 1; i < len(queue); i++ {
			if queue[i].ServiceTime < queue[shortestIndex].ServiceTime {
				shortestIndex = i
			}
		}

		shortestProcess := &queue[shortestIndex]

		shortestProcess.ServiceTime--

		if shortestProcess.ServiceTime == 0 {
			shortestProcess.WaitTime = currentTime - shortestProcess.ArrivedTime
			queue = append(queue[:shortestIndex], queue[shortestIndex+1:]...)
			executedProcesses = append(executedProcesses, *shortestProcess)
		}

		currentTime++
	}

	for _, process := range queue {
		s.Processes = append(s.Processes, process)
	}

	return executedProcesses
}

func (a *SRTF) AverageExecutionTime() {
	var totalExecutionTime int
	for _, process := range a.Processes {
		totalExecutionTime += process.ServiceTime
	}

	averageExecutionTime := float32(totalExecutionTime) / float32(len(a.Processes))
	fmt.Printf("\nTempo médio de execução: %.1f s\n", averageExecutionTime)
}

func (s *SRTF) AverageWaitingTime() {
	totalWaitTime := 0
	totalProcesses := len(s.Processes)

	for _, p := range s.Processes {
		totalWaitTime += p.WaitTime
	}

	averageWaitTime := float32(totalWaitTime) / float32(totalProcesses)
	fmt.Printf("\nTempo médio de espera: %.1f s\n", averageWaitTime)
}
