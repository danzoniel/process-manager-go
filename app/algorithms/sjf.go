package algorithms

import (
	"fmt"
	"sort"
)

type SJF struct {
	Processes []Process
}

func (s *SJF) ShortestJobFirst() []Process {
	fmt.Println("\nSHORTEST JOB FIRST")

	sort.SliceStable(s.Processes, func(i, j int) bool {
		return s.Processes[i].ServiceTime < s.Processes[j].ServiceTime
	})

	return s.Processes
}

func (s *SJF) AverageExecutionTime() {
	var totalExecutionTime int

	fmt.Println("\nTempo médio de execução:")

	for _, p := range s.Processes {
		totalExecutionTime += p.ServiceTime
	}

	averageExecutionTime := float32(totalExecutionTime) / float32(len(s.Processes))

	fmt.Printf("%.1f s\n", averageExecutionTime)
}

func (s *SJF) AverageWaitingTime() {
	fmt.Println("\nTempo médio de espera:")

	if len(s.Processes) == 0 {
		fmt.Println("Nenhum processo na fila.")
		return
	}

	var totalWaitingTime int

	totalWaitingTime += 0

	fmt.Printf("Processo %s: %d - %d = %d", s.Processes[0].ProcessId, 0, 0, 0)

	for i := 1; i < len(s.Processes); i++ {
		waitTime := totalWaitingTime + s.Processes[i-1].ServiceTime

		totalWaitingTime += s.Processes[i-1].ServiceTime

		fmt.Printf(" + Processo %s: %d - %d = %d", s.Processes[i].ProcessId, totalWaitingTime, s.Processes[i].ArrivedTime, waitTime)
	}

	averageWaitingTime := float32(totalWaitingTime) / float32(len(s.Processes))
	fmt.Printf(" / %d = %.1f s\n", len(s.Processes), averageWaitingTime)
}
