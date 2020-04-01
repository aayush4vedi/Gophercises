package scheduler

import (
	"fmt"
	"sort"
)

func SchedulerController(threadCount int, jobs []Job) {
	fmt.Printf(" number of threads: %d\n", threadCount)
	for _, j := range jobs {
		fmt.Printf("%+v\n", j)
	}
	// for FCFS
	fmt.Println("******* FCFS **********")
	s1 := make([]Schedule, 0)
	t1 := 0
	for _, j := range jobs {
		t := t1%threadCount
		s1 = append(s1, {(t1%threadCount), j})
		t1++
	}
	// for SJF
	fmt.Println("******* SJF **********")
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].Duration < jobs[j].Duration
	})
	for _, j := range jobs {
		fmt.Printf("%+v\n", j)
	}
	// for FPS
	// for EDF

}
