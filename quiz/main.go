package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

// TODO - remove globals where possible
var score int
var totalQuestions int

func main() {
	// TODO - Welcome message, enter name or something
	fmt.Println("Hello...")
	csvFileName := flag.String("csv", "problems.csv", "CSV file in 'question, answer' format.")
	flag.Parse()

	f, err := os.Open(*csvFileName)
	if err != nil {
		log.Fatalln("Err opening CSV", err)
	}

	reader := csv.NewReader(f)
	questions, err := reader.ReadAll()
	if err != nil {
		log.Fatalln("Err reading problems", err)
	}

	// TODO:
	//		While we are above EOF:
	//			Print question to the user
	//			Wait for user input
	//			Read input and compare to answer
	//			Inc/Dec user's score
	//			Repeat until EOF

	//		After EOF:
	//			Display user score and name
	//			Terminate

	fmt.Printf("You scored %v out of %v questions.\n", score, totalQuestions)
}
