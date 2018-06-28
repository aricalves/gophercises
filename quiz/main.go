package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

// TODO - remove globals where possible
var score int
var totalQuestions int

func main() {
	// TODO - Welcome message, enter name or something
	fmt.Println("Hello...")
	// TODO -

	f := openCSV("./problems.csv")
	reader := csv.NewReader(bufio.NewReader(f))

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

	printMatrix(readCSV(reader))
	rest()
	fmt.Printf("You scored %v out of %v questions.\n", score, totalQuestions)
}

func openCSV(target string) *os.File {
	f, err := os.Open(target)
	if err != nil {
		log.Fatalln("Err opening CSV", err)
	}
	return f
}

func readCSV(reader *csv.Reader) [][]string {
	data, err := reader.ReadAll()
	totalQuestions = len(data)
	if err != nil {
		log.Fatalln("Error reading problems", err)
	}
	return data
}

func printMatrix(m [][]string) {
	for _, s := range m {
		for _, item := range s {
			fmt.Println(item)
		}
	}
}

func rest() {
	time.Sleep(time.Duration(500) * time.Millisecond)
}
