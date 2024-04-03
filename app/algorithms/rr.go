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

	// Ordena os processos pelo tempo de chegada
	sort.SliceStable(a.Processes, func(i, j int) bool {
		return a.Processes[i].ArrivedTime < a.Processes[j].ArrivedTime
	})

	// Inicializa o tempo de espera e o tempo de execução
	// for i := range a.Processes {
	//     a.Processes[i].WaitTime = 0
	//     a.Processes[i].ServiceTime = a.Processes[i].ServiceTime
	// }

	// Executa o algoritmo Round Robin
	var currentTime int
	var queue []Process
	for {
		// Adiciona processos que chegam na fila
		for _, process := range a.Processes {
			if process.ArrivedTime <= currentTime && !contains(queue, process) {
				queue = append(queue, process)
			}
		}

		// Se a fila está vazia, termina o algoritmo
		if len(queue) == 0 {
			break
		}

		// Executa o próximo processo na fila
		process := queue[0]
		queue = queue[1:]

		// Executa o processo pelo quantum ou até que ele termine
		executionTime := min(a.Quantum, process.ServiceTime)
		process.ServiceTime -= executionTime
		currentTime += executionTime

		// Se o processo terminou, atualiza o tempo de espera
		if process.ServiceTime == 0 {
			process.WaitTime = currentTime - process.ArrivedTime - process.ServiceTime
		} else {
			// Se o processo não terminou, adiciona de volta à fila
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
