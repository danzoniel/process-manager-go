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
}

func (p *Process) NewProcess() Process {
	return Process{
		ProcessId:   p.ProcessId,
		Priority:    p.Priority,
		ServiceTime: p.ServiceTime,
		ArrivedTime: p.ArrivedTime,
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
	rand.Seed(time.Now().UnixNano())

	for i := uint(0); i < n; i++ {
		process := Process{
			Priority:    rand.Intn(int(n)),
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
