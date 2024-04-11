package algorithms

import (
	"fmt"
)

type Rr struct {
	Processes        []Process
	TotalProcessTime int
	Quantum          int
}

func (s *Rr) RoundRobin() {
	fmt.Println("\nROUND ROBIN")

	s.Quantum = 2
	p := make([]Process, len(s.Processes))
	available := make([]Process, 0)
	nextJob := Process{}
	actualInstant := 0
	tempNextJob := Process{}
	waitingProcessess := make([]Process, 0)
	numberOfProcesses := len(s.Processes) - 1
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

		//Procura o menor tempo entre os disponíveis
		index := 0
		//Processo não pode olhar para o tempo de chegada todas às vezes. Encontrar a excessão.
		for i, job := range available {
			if i == 0 {
				index = i
				nextJob = job
			}
		}

		nextJob.ProcessTime.startedExecutingAt = actualInstant

		//Executando 2 quantum do processo que tá sendo executado

		nextJob.ServiceTime -= s.Quantum

		if nextJob.ServiceTime < 0 {
			aux := s.Quantum + nextJob.ServiceTime
			nextJob.ProcessTime.finishedExecutingAt = actualInstant + aux

		} else {
			nextJob.ProcessTime.finishedExecutingAt = actualInstant + s.Quantum
		}
		if nextJob.ServiceTime <= 0 {
			nextJob.Done = true
			nextJob.ProcessTime.totalExecutionTime = nextJob.ProcessTime.finishedExecutingAt - nextJob.ArrivedTime
			aux = append(aux, nextJob)
		}

		//Incrementa o instante atual do algoritmo
		actualInstant = nextJob.ProcessTime.finishedExecutingAt

		//Guarda o next job como um processo temporário
		tempNextJob = nextJob

		//Limpa o processo
		nextJob = Process{}

		// res = addOrUpdateJob(res, tempNextJob)
		res = append(res, tempNextJob)
		//Remove o processo da lista de processos disponíveis
		available = append(available[:index], available[index+1:]...)
	}

	Graph(res)
	CalculateAverageProcessTime(aux)
	CalculateAverageWaitTime(p)
}
