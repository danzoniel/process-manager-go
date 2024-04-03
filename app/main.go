package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/danzoniel/process-manager-go/algorithms"
)

func main() {
	p := algorithms.Process{}

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
			processes := p.RandomizeProcesses(5)
			handleProcesses(processes)
		} else {
			numProcesses, err := strconv.Atoi(input)
			if err != nil || numProcesses > 10 {
				fmt.Println("Entrada inválida ou excede o limite de 10 processos. Gerando 5 processos aleatórios.")
				processes := p.RandomizeProcesses(5)
				handleProcesses(processes)
			} else {
				processes := p.RandomizeProcesses(uint(numProcesses))
				handleProcesses(processes)
			}
		}
	case <-time.After(10 * time.Second):
		fmt.Println("Tempo limite atingido. Gerando 5 processos aleatórios.")
		processes := p.RandomizeProcesses(5)
		handleProcesses(processes)
	}
}

func handleProcesses(processes []algorithms.Process) {
	fmt.Println("PROCESSOS")

	algorithms.PrintTable(processes)

	callingFcfs(processes)
	callingSjf(processes)
}

func callingFcfs(processes []algorithms.Process) {
	fcfs := algorithms.Fcfs{Processes: processes}

	algorithms.PrintTable(fcfs.FirstComeFirtServerd())
	fcfs.AverageExecutionTime()
	fcfs.AverageWaitingTime()
}

func callingSjf(processes []algorithms.Process) {
	sjf := algorithms.Sjf{Processes: processes}

	algorithms.PrintTable(sjf.ShortestJobFirst())
	sjf.AverageExecutionTime()
	sjf.AverageWaitingTime()
}
