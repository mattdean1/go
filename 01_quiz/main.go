package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type Problem struct {
	question string
	answer   string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format 'question,answer', e.g. '1+1,2'")
	flag.Parse()

	problems := readCsv(*csvFilename)

	correctAnswers := 0
	for i, problem := range problems {
		fmt.Printf("Problem %d: %s = \n", i+1, problem.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			correctAnswers++
		}
	}

	fmt.Printf("You scored %d out of %d \n", correctAnswers, len(problems))
}

func parseLines(lines [][]string) []Problem {
	problems := make([]Problem, len(lines))
	for i, line := range lines {
		problems[i] = Problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return problems
}

func readCsv(filename string) []Problem {
	file, err := os.Open(filename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open csv file %s", filename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse csv")
	}

	return parseLines(lines)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
