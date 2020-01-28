package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	q string
	a string
}

func parseRecords(records [][]string) []problem {
	p := make([]problem, len(records))
	for i, r := range records {
		p[i].q = r[0]
		p[i].a = r[1]
	}
	return p
}


func main() {
	csvFile := flag.String("fileName", "problems.csv", "File containing questions in the format 'question, answer'")
	flag.Parse()
	file, err := os.Open(*csvFile)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", *csvFile)
		os.Exit(1)
	}
	reader := csv.NewReader(file)

	records, _ := reader.ReadAll()

	questions := parseRecords(records)

	correct := 0
	var inputAns string
	for i, q := range questions {
		fmt.Printf("Question #%v : %v ?\n", i, q.q)
		fmt.Scanf("%s",&inputAns)
		if inputAns == q.a {
			correct++
		}
	}
	fmt.Printf("Score: %v / %v\n", correct, len(records))
}

