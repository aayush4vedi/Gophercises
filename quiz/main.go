package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func parseRecords(records [][]string) []problem {
	p := make([]problem, len(records))
	for i, r := range records {
		p[i].q = r[0]
		p[i].a = strings.TrimSpace(r[1])
	}
	return p
}

func shuffled(questions []problem) []problem {
	q := make([]problem, len(questions))
	rand.Seed(time.Now().UnixNano())
	perm := rand.Perm(len(questions))
	for i, v := range perm {
		q[v] = questions[i]
	}
	return q
}

func main() {
	csvFile := flag.String("fileName", "problems.csv", "File containing questions in the format 'question, answer'")
	shuffle := flag.Bool("shuffle", false, "Pass 'true' to shuffle the question order")
	flag.Parse()
	file, err := os.Open(*csvFile)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", *csvFile)
		os.Exit(1)
	}
	reader := csv.NewReader(file)

	records, _ := reader.ReadAll()

	questions := parseRecords(records)
	fmt.Printf("-shuffle=> &= %v\n", *shuffle)
	if *shuffle == true {
		questions = shuffled(questions)
	}
	correct := 0
	var inputAns string
	for i, q := range questions {
		fmt.Printf("Question #%v : %v ?\n", i, q.q)
		fmt.Scanf("%s", &inputAns)
		if inputAns == q.a {
			correct++
		}
	}
	fmt.Printf("Score: %v / %v\n", correct, len(records))
}
