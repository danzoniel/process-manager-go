package algorithms

import (
	"fmt"
)

type PrioC struct {
	Processes        []Process
	TotalProcessTime int
}

func (s *PrioC) PrioC() {
	fmt.Println("\nPRIOc")

	p := make([]Process, len(s.Processes))
	available := make([]Process, 0)
	nextJob := Process{}
	actualInstant := 0
	res := make([]Process, 0)

	numberOfProcesses := len(s.Processes) - 1
	//Cria um novo array com todos os processos
	for i := 0; i <= numberOfProcesses; i++ {
		p[i] = s.Processes[i].NewProcess()
	}

	for i := 0; i <= numberOfProcesses; i++ {

		//Cria um novo array com os processos disponíveis para o instante atual
		waitingProcessess := make([]Process, 0)
		for i := range p {
			if p[i].ArrivedTime <= actualInstant {
				available = append(available, p[i])
			} else {
				waitingProcessess = append(waitingProcessess, p[i])
			}
		}

		p = nil
		p = waitingProcessess

		//Procura a maior prioridade entre os disponíveis
		index := 0
		// fmt.Println("Valor da prioridade: ", shortestJob.Priority)
		for i, job := range available {
			if job.Priority > nextJob.Priority || i == 0 {
				index = i
				nextJob = job
			}
		}

		nextJob.ProcessTime.startedExecutingAt = actualInstant

		//Encontra o instante de término do processo
		nextJob.ProcessTime.finishedExecutingAt = actualInstant + nextJob.ServiceTime

		//Incrementa o instante atual do algoritmo
		actualInstant = nextJob.ProcessTime.finishedExecutingAt

		nextJob.ProcessTime.totalWaitingTime = nextJob.ProcessTime.startedExecutingAt - nextJob.ArrivedTime

		tempNextJob := nextJob

		//Limpa o processo
		nextJob = Process{}

		tempNextJob.ProcessTime.totalExecutionTime = tempNextJob.ProcessTime.finishedExecutingAt - tempNextJob.ArrivedTime

		res = append(res, tempNextJob)

		//Remove o processo da lista de processos disponíveis
		available = append(available[:index], available[index+1:]...)

	}

	Graph(res)
	CalculateAverageProcessTime(res)
	CalculateAverageWaitTime(res)
}
