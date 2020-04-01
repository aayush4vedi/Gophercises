package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Gophercises/scheduler"
)

const ThreadCount int = 2

func main() {
	//input file
	filename := flag.String("file", "data.json", "The json file with list of jobs to be scheduled")
	flag.Parse()
	fmt.Printf("Using the inputs in %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	jobs, err := scheduler.JsonParser(f)
	if err != nil {
		panic(err)
	}
	for _, j := range jobs {
		fmt.Printf("%+v\n", j)
	}

	scheduler.SchedulerController(ThreadCount, jobs)
	os.Exit(1)
}
