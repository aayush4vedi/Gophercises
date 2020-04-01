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
	AssignScheduler(threadCount,jobs)

	// for SJF
	fmt.Println("******* SJF **********")
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].Duration < jobs[j].Duration
	})
	AssignScheduler(threadCount,jobs)
	// for FPS
	fmt.Println("******* FPS **********")
	sort.Slice(jobs, func(i, j int) bool {
		if jobs[i].Priority == jobs[j].Priority{
			if jobs[i].UserType == jobs[j].UserType{
				return jobs[i].Duration<jobs[j].Duration
			}
			return jobs[i].UserType > jobs[j].UserType
		}
		return jobs[i].Duration < jobs[j].Duration
	})
	AssignScheduler(threadCount,jobs)
	// for EDF

}
