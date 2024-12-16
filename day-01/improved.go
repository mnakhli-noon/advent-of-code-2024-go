package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseNumbersFromFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var firstList, secondList []int

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())

		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("Invalid line format: %s", scanner.Text())
		}

		firstNumber, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("Invalid number if first list: %s\n", parts[0])
		}

		secondNumber, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("Invalid number if second list: %s\n", parts[1])
		}

		firstList = append(firstList, firstNumber)
		secondList = append(secondList, secondNumber)
	}

	return firstList, secondList, nil
}

func abs(number int) int {
	if number < 0 {
		return -number
	} else {
		return number
	}
}

func solvePartOneImproved() {
	firstList, secondList, err := parseNumbersFromFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	slices.Sort(firstList)
	slices.Sort(secondList)

	distance := 0

	for i := 0; i < len(firstList); i++ {
		distance += abs(firstList[i] - secondList[i])
	}
	fmt.Println("Distance: ", distance)
}

func solvePartTwoImproved() {
	firstList, secondList, err := parseNumbersFromFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	occurrences := make(map[int]int)

	for _, num := range secondList {
		occurrences[num]++
	}

	similarity := 0

	for _, num := range firstList {
		similarity += num * occurrences[num]
	}

	fmt.Println("Similarity: ", similarity)
}
