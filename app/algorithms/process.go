package algorithms

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type Process struct {
	ProcessId   string
	Priority    int
	ServiceTime int
	ArrivedTime int
	ProcessTime ProcessTime
}

type ProcessTime struct {
	// startedExecutingAt  int
	finishedExecutingAt int
	// startedWaitingAt    int
	// finishedWaitingAt   int
	totalExecutionTime int
	totalWaitingTime   int
}

func (p *Process) NewProcess() Process {
	return Process{
		ProcessId:   p.ProcessId,
		Priority:    p.Priority,
		ServiceTime: p.ServiceTime,
		ArrivedTime: p.ArrivedTime,
		ProcessTime: p.ProcessTime,
	}
}

func PrintTable(processess []Process) {
	fmt.Printf("| %-12s | %-9s | %-15s | %-15s |\n", "ID Processo", "Prioridade", "Tempo de Serviço", "Tempo de Chegada")
	fmt.Println("+--------------+-----------+-----------------+-----------------+")
	for _, processo := range processess {
		fmt.Printf("| %-12s | %-9s | %-15s | %-15s |\n", processo.ProcessId, strconv.Itoa(int(processo.Priority)), strconv.Itoa(int(processo.ServiceTime)), strconv.Itoa(int(processo.ArrivedTime)))
	}
}

func (p *Process) RandomizeProcesses(n uint) []Process {
	fmt.Println("GERANDO", n, "PROCESSOS ALEATÓRIOS")
	processes := make([]Process, n)
	rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := uint(0); i < n; i++ {
		process := Process{
			Priority:    rand.Intn(int(n)) + 1,
			ServiceTime: rand.Intn(int(n)) + 1,
			ArrivedTime: rand.Intn(int(n)),
		}
		processes[i] = process
	}

	sort.Slice(processes, func(i, j int) bool {
		return processes[i].ArrivedTime < processes[j].ArrivedTime
	})

	for i := range processes {
		idStr := strconv.Itoa(i + 1)
		processes[i].ProcessId = "T" + idStr
	}

	return processes
}

func CalculateAverageProcessTime(processes []Process) {
	totalProcessTime := 0

	for i := range processes {
		totalProcessTime += processes[i].ProcessTime.totalExecutionTime
	}

	averageTotalProcessTime := float64(totalProcessTime) / float64(len(processes))

	fmt.Printf("\nTempo médio de processo: %.1fs\n", averageTotalProcessTime)
}

func CalculateAverageWaitTime(processes []Process) {
	totalWaitTime := 0

	for i := range processes {
		totalWaitTime += processes[i].ProcessTime.totalWaitingTime
	}

	averageTotalWaitTime := float64(totalWaitTime) / float64(len(processes))

	fmt.Printf("\nTempo médio de espera: %.1fs\n", averageTotalWaitTime)
}
