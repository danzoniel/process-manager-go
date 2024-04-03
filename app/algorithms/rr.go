package algorithms

import (
	"fmt"
	"sort"
)

type RoundRobin struct {
	Processes []Process
	Quantum   int
}

func (r *RoundRobin) Schedule() []Process {
	fmt.Println("\nROUND ROBIN")

	// Ordena os processos pelo tempo de chegada.
	sort.SliceStable(r.Processes, func(i, j int) bool {
		return r.Processes[i].ArrivedTime < r.Processes[j].ArrivedTime
	})

	var waitingQueue []Process
	var completedQueue []Process
	var currentTime int

	for len(r.Processes) > 0 || len(waitingQueue) > 0 {
		// Move todos os processos que chegaram até agora para a fila de espera.
		for _, process := range r.Processes {
			if process.ArrivedTime <= currentTime {
				waitingQueue = append(waitingQueue, process)
			} else {
				break
			}
		}
		if len(r.Processes) >= len(waitingQueue) {
			r.Processes = r.Processes[len(waitingQueue):]
		} else {
			r.Processes = nil // ou r.Processes = []Process{}
		}

		// Se não houver processos na fila de espera, avança o tempo até o próximo processo chegar.
		if len(waitingQueue) == 0 {
			if len(r.Processes) > 0 {
				currentTime = r.Processes[0].ArrivedTime
			}
			continue
		}

		// Executa o próximo processo na fila de espera.
		process := waitingQueue[0]
		waitingQueue = waitingQueue[1:]

		// Se o processo não termina dentro do quantum, interrompe-o e coloca-o de volta na fila de espera.
		if process.ServiceTime > r.Quantum {
			process.ServiceTime -= r.Quantum
			waitingQueue = append(waitingQueue, process)
		} else {
			completedQueue = append(completedQueue, process)
		}

		// Avança o tempo pelo quantum.
		currentTime += r.Quantum
	}

	// Imprime os processos completados.
	fmt.Println("\nProcessos completados:")
	for _, process := range completedQueue {
		fmt.Printf("Tempo de chegada: %d, Tempo de serviço: %d\n", process.ArrivedTime, process.ServiceTime)
	}

	return completedQueue
}

// AverageExecutionTime calcula o tempo médio de execução dos processos completados.
func (r *RoundRobin) AverageExecutionTime() {
	completedProcesses := r.Schedule()
	var totalExecutionTime int
	for _, process := range completedProcesses {
		totalExecutionTime += process.ServiceTime
	}
	averageExecutionTime := float32(totalExecutionTime) / float32(len(completedProcesses))
	fmt.Printf("\nTempo médio de execução: %.1f s\n", averageExecutionTime)
}

// AverageWaitingTime calcula o tempo médio de espera dos processos completados.
func (r *RoundRobin) AverageWaitingTime() {
	completedProcesses := r.Schedule()
	var totalWaitTime int
	for i, process := range completedProcesses {
		if i != 0 {
			totalWaitTime += completedProcesses[i-1].ServiceTime
		}
		totalWaitTime += process.ArrivedTime
	}
	averageWaitingTime := float32(totalWaitTime) / float32(len(completedProcesses))
	fmt.Printf("\nTempo médio de espera: %.1f s\n", averageWaitingTime)
}
