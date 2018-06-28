package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	q, a string
}

func main() {
	timeLimit := flag.Int("timer", 30, "Quiz time limit (seconds)")
	csvFileName := flag.String("csv", "problems.csv", "CSV file in 'question, answer' format.")
	shuffle := flag.Bool("shuffle", false, "If shuffle is true, problems will be presented in a random order.")
	flag.Parse()

	f, err := os.Open(*csvFileName)
	if err != nil {
		log.Fatalln("Err opening CSV", err)
	}

	reader := csv.NewReader(f)
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Err reading problems", err)
	}

	questions := parseQuestions(lines)

	if *shuffle {
		rand.Shuffle(len(questions), func(i, j int) {
			questions[i], questions[j] = questions[j], questions[i]
		})
	}

	fmt.Println("What is your name? (Timer will start after you hit 'return')")
	var user string
	fmt.Scanf("%s", &user)

	var score int
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	go func() {
		<-timer.C
		fmt.Println("\nTimer expired")
		printScore(user, score, len(questions))
	}()

	for i, p := range questions {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		var answer string
		fmt.Scanf("%s", &answer)
		if answer == p.a {
			score++
		} else {
			fmt.Println("Incorrect :/")
		}
	}
	printScore(user, score, len(questions))
}

func parseQuestions(q [][]string) []problem {
	problems := make([]problem, len(q))
	for i, line := range q {
		problems[i] = problem{
			line[0],
			strings.TrimSpace(line[1]),
		}
	}
	return problems
}

func printScore(u string, s, q int) {
	fmt.Printf("%s scored %d out of %d questions.\n", u, s, q)
	os.Exit(0)
}
