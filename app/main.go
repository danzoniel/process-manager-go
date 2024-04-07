package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/danzoniel/process-manager-go/algorithms"
)

func main() {
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
			fmt.Println("Nenhum número de processos fornecido. Gerando 5 processos aleatórios.")
			// processes := p.RandomizeProcesses(5)
			handleProcesses(processes)
		} else {
			numProcesses, err := strconv.Atoi(input)
			if err != nil || numProcesses > 10 {
				fmt.Println("Entrada inválida ou excede o limite de 10 processos. Gerando 5 processos aleatórios.")
				// processes := p.RandomizeProcesses(5)
				handleProcesses(processes)
			} else {
				// processes := p.RandomizeProcesses(uint(numProcesses))
				handleProcesses(processes)
			}
		}
	case <-time.After(10 * time.Second):
		fmt.Println("Tempo limite atingido. Gerando 5 processos aleatórios.")
		// processes := p.RandomizeProcesses(5)
		handleProcesses(processes)
	}
}

func handleProcesses(processes []algorithms.Process) {
	fmt.Println("PROCESSOS")

	algorithms.PrintTable(processes)

	callFcfs(processes)
}

func callFcfs(processes []algorithms.Process) {
	fcfs := algorithms.Fcfs{Processes: processes}

	algorithms.PrintTable(fcfs.FirstComeFirtServerd())
}
