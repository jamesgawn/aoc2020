package main

import (
	"aoc2020/day1"
	"aoc2020/day2"
	"aoc2020/day3"
	"aoc2020/day4"
	"aoc2020/pkg/utils"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	defer utils.TimeTrack(time.Now(), "Total execution")
	day := "4"
	if len(os.Args) > 1 {
		day = string(os.Args[1])
	}
	inputFileLocation := "./day" + day + "/input.txt"
	if len(os.Args) > 2 {
		inputFileLocation = string(os.Args[1])
	}

	fmt.Println("Using input file: " + inputFileLocation)
	file, err := os.Open(inputFileLocation)
	if err != nil {
		log.Fatal(err)
	}

	FindSolution(day, file)

	cErr := file.Close()
	if cErr != nil {
		fmt.Println(cErr)
	}
}

func FindSolution(day string, input io.Reader) {
	switch day {
	case "1":
		fmt.Println("Day " + day)
		day1.ExecuteSolution(input)
	case "2":
		fmt.Println("Day " + day)
		day2.ExecuteSolution(input)
	case "3":
		fmt.Println("Day " + day)
		day3.ExecuteSolution(input)
	case "4":
		fmt.Println("Day " + day)
		day4.ExecuteSolution(input)
	}
}