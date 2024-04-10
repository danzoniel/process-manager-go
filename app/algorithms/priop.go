package algorithms

import (
	"fmt"
)

type PrioP struct {
	Processes        []Process
	TotalProcessTime int
}

func (s *PrioP) PrioP() []Process {
	fmt.Println("\nPRIOp")

	p := make([]Process, len(s.Processes))
	available := make([]Process, 0)
	nextJob := Process{}
	actualInstant := 0
	tempNextJob := Process{}
	waitingProcessess := make([]Process, 0)
	numberOfProcesses := len(s.Processes) - 1
	moment := 1

	//Cria um novo array com todos os processos e conta o tempo total de todos os processos
	for i := 0; i <= numberOfProcesses; i++ {
		p[i] = s.Processes[i].NewProcess()
		s.TotalProcessTime += p[i].ServiceTime
	}

	// fmt.Println("Tempo total de todos os processos somados: ", s.TotalProcessTime)

	// fmt.Println("Processos recebidos: ", s.Processes)

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

		// fmt.Println("Processo temporário que talvez volte pra fila de espera: ", tempNextJob)

		if tempNextJob.ServiceTime > 0 {
			// fmt.Println("Processo temporário que tá voltando pra lista de espera: ", tempNextJob)

			available = append(available, tempNextJob)
		}

		p = nil
		p = waitingProcessess

		// fmt.Println("Processos disponíveis para executar", available)
		// fmt.Println("Processos na lista de espera", p)

		//Procura a maior prioridade entre os disponíveis
		index := 0
		// fmt.Println("Valor da prioridade: ", nextJob.Priority)
		// fmt.Println("Valor do nextJob: ", nextJob)
		for i, job := range available {
			if job.Priority > nextJob.Priority || i == 0 {
				index = i
				nextJob = job
			}
		}

		fmt.Println("Processo que será executado: ", nextJob)

		// fmt.Println("Index do processo que será removido da lista de disponíveis: ", index)

		// fmt.Println("Instante atual antes do processo: ", actualInstant)

		//Execução do processo
		nextJob.ServiceTime -= moment

		nextJob.ProcessTime.finishedExecutingAt = actualInstant + moment

		//Checa para ver se há algum processo disponível que merece interromper

		fmt.Println("Next Job depois de processar o quantum", nextJob)

		//Encontra o instante de término do processo

		//Incrementa o instante atual do algoritmo
		actualInstant = nextJob.ProcessTime.finishedExecutingAt

		//Guarda o next job como um processo temporário
		tempNextJob = nextJob

		// fmt.Println("Instante que o processo atual terminou de executar: ", nextJob.ProcessTime.finishedExecutingAt)
		fmt.Println("Instante atual: ", actualInstant)

		//Limpa o processo
		nextJob = Process{}

		//Remove o processo da lista de processos disponíveis
		available = append(available[:index], available[index+1:]...)
		// fmt.Println("Lista de disponíveis depois da remoção: ", available)

	}

	return nil
}
