package algorithms

import (
	"fmt"
)

type Rr struct {
	Processes        []Process
	TotalProcessTime int
	Quantum          int
}

func (s *Rr) RoundRobin() []Process {
	fmt.Println("\nROUND ROBIN")

	s.Quantum = 2
	p := make([]Process, len(s.Processes))
	available := make([]Process, 0)
	nextJob := Process{}
	actualInstant := 0

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

		fmt.Println("Processos disponíveis para executar", available)
		fmt.Println("Processos na lista de espera", p)

		//Procura o menor tempo entre os disponíveis
		index := 0
		// fmt.Println("Valor de chegada do próximo processo a ser executado: ", nextJob.ArrivedTime)
		for i, job := range available {
			if job.ArrivedTime < nextJob.ArrivedTime || i == 0 {
				index = i
				nextJob = job
			}
		}

		fmt.Println("Processo que será executado: ", nextJob)

		fmt.Println("Index do processo que será removido da lista de disponíveis: ", index)

		fmt.Println("Instante atual antes do processo: ", actualInstant)

		//Executando 2 quantum do processo que tá sendo executado

		//Tira 2 quantum do processo, remove ele da fila e o coloca ao final da fila se tempo de serviço - 2 > 0
		fmt.Println("Next Job antes do processa o quantum", nextJob)

		nextJob.ServiceTime -= s.Quantum

		fmt.Println("Next Job depois do processa o quantum", nextJob)

		//Encontra o instante de término do processo
		nextJob.ProcessTime.finishedExecutingAt = actualInstant + nextJob.ServiceTime

		//Incrementa o instante atual do algoritmo
		actualInstant = nextJob.ProcessTime.finishedExecutingAt

		if nextJob.ServiceTime > 0 {

		}

		fmt.Println("Instante que o processo atual terminou de executar: ", nextJob.ProcessTime.finishedExecutingAt)
		fmt.Println("Instante atual: ", actualInstant)

		//Limpa o processo
		nextJob = Process{}

		//Remove o processo da lista de processos disponíveis
		available = append(available[:index], available[index+1:]...)
		fmt.Println("Lista de disponíveis depois da remoção: ", available)

	}

	return nil
}
