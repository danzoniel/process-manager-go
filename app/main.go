package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/danzoniel/process-manager-go/algorithms"
)

func main() {
	p := algorithms.Process{}
	processes := []algorithms.Process{
		{
			ProcessId:   "t1",
			Priority:    2,
			ServiceTime: 5,
			ArrivedTime: 0,
		},
		{
			ProcessId:   "t2",
			Priority:    3,
			ServiceTime: 2,
			ArrivedTime: 0,
		},
		{
			ProcessId:   "t3",
			Priority:    1,
			ServiceTime: 4,
			ArrivedTime: 1,
		},
		{
			ProcessId:   "t4",
			Priority:    4,
			ServiceTime: 1,
			ArrivedTime: 3,
		},
		{
			ProcessId:   "t5",
			Priority:    5,
			ServiceTime: 2,
			ArrivedTime: 5,
		},
	}

	userInput := make(chan string)

	go func() {
		var input string
		fmt.Print("Digite o número de processos (ou pressione Enter para 5 processos aleatórios): ")

		fmt.Scanln(&input)
		userInput <- input

	}()

	select {
	case input := <-userInput:
		if input == "" {
			fmt.Println("Nenhum número de processos fornecido. Gerando 5 processos default.")
			handleProcesses(processes)
		} else {
			numProcesses, err := strconv.Atoi(input)
			if err != nil || numProcesses > 10 {
				fmt.Println("Entrada inválida ou excede o limite de 10 processos. Gerando 5 processos aleatórios.")
				processes = p.RandomizeProcesses(5)
				handleProcesses(processes)
			} else {
				processes = p.RandomizeProcesses(uint(numProcesses))
				handleProcesses(processes)
			}
		}
	case <-time.After(5 * time.Second):
		fmt.Println("Tempo limite atingido. Gerando 5 processos aleatórios.")
		processes = p.RandomizeProcesses(5)
		handleProcesses(processes)
	}
}

func handleProcesses(processes []algorithms.Process) {
	fmt.Println("PROCESSOS")

	algorithms.PrintTable(processes)

	callFcfs(processes)
	callSjf(processes)
	callRr(processes)
	callPrioC(processes)
	callPrioP(processes)
	callSrtf(processes)
}

func callFcfs(processes []algorithms.Process) {
	fcfs := algorithms.Fcfs{Processes: processes}

	fcfs.FirstComeFirtServerd()
}

func callSjf(processes []algorithms.Process) {
	sjf := algorithms.Sjf{Processes: processes}

	sjf.ShortestJobFirst()
}

func callRr(processes []algorithms.Process) {
	rr := algorithms.Rr{Processes: processes}

	rr.RoundRobin()
}

func callPrioC(processes []algorithms.Process) {
	prioc := algorithms.PrioC{Processes: processes}

	prioc.PrioC()
}

func callPrioP(processes []algorithms.Process) {
	priop := algorithms.PrioP{Processes: processes}

	priop.PrioP()
}

func callSrtf(processes []algorithms.Process) {
	srtf := algorithms.Srtf{Processes: processes}

	srtf.ShortestRemainingJobFirst()
}
