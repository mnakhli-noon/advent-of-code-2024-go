// this improved version was inspired by https://www.youtube.com/@josh_ackland

package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readFile() (string, error) {
	content, err := os.ReadFile("input.txt")
	return strings.TrimSpace(string(content)), err
}

func solvePartOneImproved() {
	input, err := readFile()
	if err != nil {
		log.Fatal("Can't open input file ", err)
	}
	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	operations := mulRegex.FindAllStringSubmatch(input, -1)

	result := 0
	for _, value := range operations {
		firstNumber, _ := strconv.Atoi(value[1])
		secondNumber, _ := strconv.Atoi(value[2])

		result += firstNumber * secondNumber
	}

	fmt.Println(result)

}

func solvePartTwoImproved() {

	input, err := readFile()
	if err != nil {
		log.Fatal("Can't open input file ", err)
	}
	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)`)

	operations := mulRegex.FindAllStringSubmatch(input, -1)

	result := 0
	enabled := true
	for _, value := range operations {
		if value[0] == "don't()" {
			enabled = false
			continue
		} else if value[0] == "do()" {
			enabled = true
			continue
		} else if enabled {
			firstNumber, _ := strconv.Atoi(value[1])
			secondNumber, _ := strconv.Atoi(value[2])
			result += firstNumber * secondNumber
		}

	}

	fmt.Println(result)

}
