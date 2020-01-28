package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

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

	correct := 0
	totalQuestions := len(records)
	var inputAns string
	for i, data := range records {
		fmt.Printf("Question #%v : %v ?\n", i, data[0])
		fmt.Scan(&inputAns)

		if inputAns == data[1] {
			correct++
		}
	}
	fmt.Printf("Score: %v / %v\n", correct, totalQuestions)
}
