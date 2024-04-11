package algorithms

import (
	"fmt"
	"strconv"
	"strings"
)

type Srtf struct {
	Processes        []Process
	TotalProcessTime int
}

func (s *Srtf) ShortestRemainingJobFirst() {
	fmt.Println("\nSHORTEST REMANING TIME FIRST")

	p := make([]Process, len(s.Processes))
	available := make([]Process, 0)
	nextJob := Process{}
	actualInstant := 0
	tempNextJob := Process{}
	waitingProcessess := make([]Process, 0)
	numberOfProcesses := len(s.Processes) - 1
	moment := 1
	res := make([]Process, 0)
	aux := make([]Process, 0)

	//Cria um novo array com todos os processos e conta o tempo total de todos os processos
	for i := 0; i <= numberOfProcesses; i++ {
		p[i] = s.Processes[i].NewProcess()
		s.TotalProcessTime += p[i].ServiceTime
	}

	for actualInstant < s.TotalProcessTime {

		//Cria um novo array com os processos disponíveis para o instante atual
		waitingProcessess = make([]Process, 0)
		for i := range p {
			if p[i].ArrivedTime <= actualInstant {
				available = append(available, p[i])
			} else {
				waitingProcessess = append(waitingProcessess, p[i])
			}
		}

		if tempNextJob.ServiceTime > 0 {
			available = append(available, tempNextJob)
		}

		p = nil
		p = waitingProcessess

		//Procura a maior prioridade entre os disponíveis
		index := 0

		for i, job := range available {
			if job.ServiceTime < nextJob.ServiceTime || i == 0 {
				index = i
				nextJob = job
			} else if job.ServiceTime == nextJob.ServiceTime {
				jobId, _ := strconv.Atoi(strings.TrimPrefix(job.ProcessId, "t"))
				nextJobId, _ := strconv.Atoi(strings.TrimPrefix(nextJob.ProcessId, "t"))

				if jobId < nextJobId {
					index = i
					nextJob = job
				}
			}
		}

		nextJob.ProcessTime.startedExecutingAt = actualInstant

		//Execução do processo
		nextJob.ServiceTime -= moment

		nextJob.ProcessTime.finishedExecutingAt = actualInstant + moment

		nextJob.ProcessTime.totalWaitingTime = nextJob.ProcessTime.startedExecutingAt - nextJob.ArrivedTime

		if nextJob.ServiceTime <= 0 {
			nextJob.Done = true
			nextJob.ProcessTime.totalExecutionTime = nextJob.ProcessTime.finishedExecutingAt - nextJob.ArrivedTime
			aux = append(aux, nextJob)
		}

		//Checa para ver se há algum processo disponível que merece interromper

		//Encontra o instante de término do processo

		//Incrementa o instante atual do algoritmo
		actualInstant = nextJob.ProcessTime.finishedExecutingAt

		//Guarda o next job como um processo temporário
		tempNextJob = nextJob

		//Limpa o processo
		nextJob = Process{}

		tempNextJob.ProcessTime.totalExecutionTime = tempNextJob.ProcessTime.finishedExecutingAt - tempNextJob.ArrivedTime

		res = append(res, tempNextJob)

		//Remove o processo da lista de processos disponíveis
		available = append(available[:index], available[index+1:]...)

	}

	Graph(res)
	CalculateAverageProcessTime(aux)
	CalculateAverageWaitTime(p)
}
