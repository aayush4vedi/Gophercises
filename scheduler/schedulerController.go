package scheduler

import "fmt"

func SchedulerController(threadCount int, jobs []Job) {
	fmt.Printf(" number of threads: %d\n", threadCount)
	for _, j := range jobs {
		fmt.Printf("%+v\n", j)
	}
	// for SJF	
	sort.Slice(jobs, func(i,j int)bool{
		return jobs[i].Duration < jobs[j].Duration
	})

	// for FCFS
	// for FPS	
	// for EDF	




}