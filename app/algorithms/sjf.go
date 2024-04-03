package algorithms

// import (
// 	"fmt"
// 	"sort"
// )

// // ShortestJobFirst implementa o algoritmo de escalonamento SJF.
// func ShortestJobFirst(p []Process) []Process {
// 	fmt.Println("\nSHORTEST JOB FIRST")

// 	// Ordena os processos primeiro por tempo de chegada e, em caso de empate, por tempo de serviço.
// 	sort.SliceStable(p, func(i, j int) bool {
// 		if p[i].ArrivedTime == p[j].ArrivedTime {
// 			return p[i].ServiceTime < p[j].ServiceTime
// 		}
// 		return p[i].ArrivedTime < p[j].ArrivedTime
// 	})

// 	return p
// }

// // AverageExecutionTime calcula o tempo médio de execução dos processos.
// func AverageExecutionTime(p []Process) {
// 	var totalExecutionTime int
// 	for _, process := range p {
// 		totalExecutionTime += int(process.ServiceTime)
// 	}
// 	AverageExecutionTime := float32(totalExecutionTime) / float32(len(p))
// 	fmt.Printf("\nTempo médio de execução: %.1f s\n", AverageExecutionTime)
// }

// // AverageWaitingTime calcula o tempo médio de espera dos processos.
// func AverageWaitingTime(p []Process) {
// 	var totalWaitTime int
// 	for i, process := range p {
// 		if i != 0 {
// 			totalWaitTime += int(p[i-1].ServiceTime)
// 		}
// 		totalWaitTime += process.ArrivedTime
// 	}
// 	averageWaitingTime := float32(totalWaitTime) / float32(len(p))
// 	fmt.Printf("\nTempo médio de espera: %.1f s\n", averageWaitingTime)
// }
