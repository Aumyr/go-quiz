package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func loadFile() *os.File {
	csvfile, err := os.Open("problems.csv")
	if len(os.Args[:]) == 2 {
		csvfile, err = os.Open(os.Args[1])
	}
	if err != nil {
		log.Fatalln("Cannot load the file", err)
	}
	return csvfile
}

func main() {
	loader := loadFile()
	reader := csv.NewReader(loader)
	scanner := bufio.NewScanner(os.Stdin)
	questions, correctAnswers := 0, 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Error reading line")
		}
		questions++
		fmt.Println(record[0])
		scanner.Scan()
		if record[1] == scanner.Text() {
			correctAnswers++
			fmt.Println("Correct!")
		}
	}
	fmt.Println("Your score is", correctAnswers, "out of", questions)
}
