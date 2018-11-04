package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

type Problem struct {
	question string
	answer   string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format 'question,answer', e.g. '1+1,2'")
	timeLimit := flag.Int("time", 30, "the time limit in seconds")
	flag.Parse()

	problems := readCsv(*csvFilename)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correctAnswers := 0
	for i, problem := range problems {
		select {
		case <-timer.C:
			endQuiz(correctAnswers, len(problems))
		default:
			fmt.Printf("Problem %d: %s = \n", i+1, problem.question)
			var answer string
			fmt.Scanf("%s\n", &answer)
			if answer == problem.answer {
				correctAnswers++
			}
		}
	}
	endQuiz(correctAnswers, len(problems))
}

func readCsv(filename string) []Problem {
	file, err := os.Open(filename)
	if err != nil {
		exit(1, fmt.Sprintf("Failed to open csv file %s", filename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(1, "Failed to parse csv")
	}

	return parseLines(lines)
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

func endQuiz(score, total int) {
	exit(0, fmt.Sprintf("You scored %d out of %d \n", score, total))
}

func exit(code int, msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
