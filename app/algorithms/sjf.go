package algorithms

import (
	"fmt"
	"sort"
)

type Sjf struct {
	Processes []Process
}

// ShortestJobFirst implementa o algoritmo de escalonamento SJF.
func (a *Sjf) ShortestJobFirst() []Process {
	fmt.Println("\nSHORTEST JOB FIRST")

	// Ordena os processos primeiro por tempo de chegada e, em caso de empate, por tempo de serviço.
	sort.SliceStable(a.Processes, func(i, j int) bool {
		if a.Processes[i].ArrivedTime == a.Processes[j].ArrivedTime {
			return a.Processes[i].ServiceTime < a.Processes[j].ServiceTime
		}
		return a.Processes[i].ArrivedTime < a.Processes[j].ArrivedTime
	})

	return a.Processes
}

// AverageExecutionTime calcula o tempo médio de execução dos processos.
func (a *Sjf) AverageExecutionTime() {
	var totalExecutionTime int
	for _, process := range a.Processes {
		totalExecutionTime += int(process.ServiceTime)
	}
	AverageExecutionTime := float32(totalExecutionTime) / float32(len(a.Processes))
	fmt.Printf("\nTempo médio de execução: %.1f s\n", AverageExecutionTime)
}

// AverageWaitingTime calcula o tempo médio de espera dos processos.
func (a *Sjf) AverageWaitingTime() {
	var totalWaitTime int
	for i, process := range a.Processes {
		if i != 0 {
			totalWaitTime += int(a.Processes[i-1].ServiceTime)
		}
		totalWaitTime += int(process.ArrivedTime)
	}
	averageWaitingTime := float32(totalWaitTime) / float32(len(a.Processes))
	fmt.Printf("\nTempo médio de espera: %.1f s\n", averageWaitingTime)
}
