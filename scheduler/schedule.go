package scheduler

import "fmt"

type Schedule map[int][]Job

func PrintSchedule(s Schedule) {
	for k,v := range s{
		fmt.Printf("#Thread %d : ",k)
		for _,j := range v{
			fmt.Printf("%+v  ",j.Name)
		}
		fmt.Println("")
	}
}

func AssignScheduler(threadCount int, jobs []Job){
	s1 := make(Schedule)
	t1 := 0
	for _, j := range jobs {
		t := t1 % threadCount
		s1[t] = append(s1[t], j)
		t1++
	}
	PrintSchedule(s1)
}


