package algorithms

import (
	"fmt"
)

type Sjf struct {
	Processes        []Process
	TotalProcessTime int
}

func (s *Sjf) ShortestJobFirst() []Process {
	fmt.Println("\nSHORTEST JOB FIRST")

	p := make([]Process, len(s.Processes))
	available := make([]Process, 0)
	shortestJob := Process{}

	actualInstant := 0

	//Cria um novo array com todos os processos
	for i := range s.Processes {
		p[i] = s.Processes[i].NewProcess()
	}

	n := 0
	for n <= 4 {	

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
			fmt.Println("Valor do shortest job arrived time: ", shortestJob.ArrivedTime)
			for i, job := range available {
				if job.ServiceTime < shortestJob.ServiceTime || i == 0{
					index = i
					shortestJob = job
				}
			}

			fmt.Println("Processo com o menor tempo de serviço que será executado: ", shortestJob)

			fmt.Println("Index do processo que será removido da lista de disponíveis: ", index)

			fmt.Println("Instante atual antes do processo: ", actualInstant)

			//Encontra o instante de término do processo 
			shortestJob.ProcessTime.finishedExecutingAt = actualInstant + shortestJob.ServiceTime

			//Incrementa o instante atual do algoritmo
			actualInstant = shortestJob.ProcessTime.finishedExecutingAt

			fmt.Println("Instante que o processo atual terminou de executar: ", shortestJob.ProcessTime.finishedExecutingAt)
			fmt.Println("Instante atual: ", actualInstant)

			//Limpa o processo
			shortestJob = Process{}

			//Remove o processo da lista de processos disponíveis
			available = append(available[:index], available[index+1:]...)
			fmt.Println("Lista de disponíveis depois da remoção: ", available)

			n++
		}

		return nil
	}

// func removeMatchedElements(slice []Process, match int) []int {
// 	result := []int{}

// 	for _, element := range slice {
// 		if element != match {
// 			result = append(result, element)
// 		}
// 	}

// 	return result
// }
