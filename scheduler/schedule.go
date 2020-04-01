package scheduler

import "fmt"

type Schedule struct {
	Thread int
	Jobs   []Job
}

func PrintSchedule(schedules []Schedule) {
	for _, s := range schedules {
		fmt.Printf("Thread: %d -> ", s.Thread)
		for _, j := range s.Jobs {
			fmt.Printf("%v ,", j.Name)
		}
		fmt.Println()
	}
}
