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
	availabe := make([]Process, 0)
	shortestJob := Process{}

	actualInstant := 0

	for i := range s.Processes {
		p[i] = s.Processes[i].NewProcess()
	}

	for i := range s.Processes {
		if p[i].ArrivedTime <= actualInstant {
			availabe = append(availabe, p[i])
		}
	}
	index := 0
	for y := 0; y <= len(availabe)-1; y++ {
		if availabe[0].ArrivedTime == availabe[y].ArrivedTime {
			if availabe[0].ServiceTime > availabe[y].ServiceTime {
				shortestJob = availabe[y]
				index = y
			}
		}
	}

	shortestJob.ProcessTime.finishedExecutingAt = actualInstant + shortestJob.ServiceTime

	actualInstant = shortestJob.ProcessTime.finishedExecutingAt

	fmt.Println(availabe)
	fmt.Println(shortestJob)

	shortestJob = Process{}
	availabe = append(availabe[:index], availabe[index+1:]...)

	fmt.Println(availabe)
	fmt.Println(shortestJob)
	fmt.Println(actualInstant)

	// for i := range s.Processes {

	// 	if p[i].ArrivedTime <= actualInstant {
	// 		availabe = append(availabe, p[i])
	// 	}
	// }

	// for y := 0; y <= len(availabe)-1; y++ {
	// 	if availabe[0].ArrivedTime == availabe[y].ArrivedTime {
	// 		if availabe[0].ServiceTime > availabe[y].ServiceTime {
	// 			shortestJob = append(shortestJob, availabe[y])
	// 		}
	// 	}
	// }

	// fmt.Println(availabe)
	// fmt.Println(shortestJob)

	return nil
}

func removeMatchedElements(slice []Process, match int) []int {
	result := []int{}

	for _, element := range slice {
		if element != match {
			result = append(result, element)
		}
	}

	return result
}
