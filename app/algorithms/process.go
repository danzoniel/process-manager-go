package algorithms

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Process struct {
	ProcessId   string
	Priority    uint
	ServiceTime uint
	ArrivedTime uint
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
	counter := 1

	for range n {
		idStr := strconv.Itoa(int(counter))
		p.ProcessId = "T" + idStr
		p.Priority = uint(rand.Intn(int(n)))
		p.ServiceTime = uint(rand.Intn(int(n)) + 1)
		p.ArrivedTime = uint(rand.Intn(int(n)))
		processes[counter-1] = p.NewProcess()
		counter++
	}

	return processes
}
