package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func readCSV(csvFile string) *csv.Reader {
	fileBytes, err := os.ReadFile(csvFile)

	if err != nil {
		os.Exit(1)
	}

	fileContent := string(fileBytes)

	csvParsed := csv.NewReader(strings.NewReader(fileContent))
	csvParsed.Comma = ','
	return csvParsed
}

func main() {
	var csv string
	var limit int
	flag.StringVar(&csv, "csv", "quiz.csv", "Default path to quiz csv file")
	flag.IntVar(&limit, "limit", 30, "The time limit for the quiz in seconds")
	flag.Parse()

	csvParsed := readCSV(csv)
	records, err := csvParsed.ReadAll()

	if err != nil {
		os.Exit(1)
	}

	reader := bufio.NewReader(os.Stdin)

	wrongTotal := 0
	correctTotal := 0

	for _, record := range records {
		fmt.Printf("Say to me the result of %s\n", record[0])
		result, _ := reader.ReadString('\n')
		answer := record[1]
		if strings.Trim(result, "\n") == strings.Trim(answer, "\n") {
			fmt.Println("Correct!")
			correctTotal += 1
		} else {
			fmt.Println("Wrong")
			wrongTotal += 1
		}

	}

	fmt.Printf("\n\nCorrect Answers: %d\n", correctTotal)
	fmt.Printf("\n\nWrong Answers: %d\n", wrongTotal)

}
