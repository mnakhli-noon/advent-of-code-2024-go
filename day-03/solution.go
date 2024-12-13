package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type acceptedChar struct {
	value            string
	nextAcceptedChar string
}
type condtionalChar struct {
	value             string
	doNext            string
	dontNext          string
	shouldChangeState bool
}

var acceptMap = map[string]acceptedChar{
	"m": {
		value:            "m",
		nextAcceptedChar: "u",
	},
	"u": {
		value:            "u",
		nextAcceptedChar: "l",
	},
	"l": {
		value:            "l",
		nextAcceptedChar: "(",
	},
	"(": {
		value:            "(",
		nextAcceptedChar: "digit1",
	},
	"digit1": {
		value:            "digit1",
		nextAcceptedChar: ",",
	},
	"digit2": {
		value:            "digit2",
		nextAcceptedChar: ")",
	},
}

var conditionalMap = map[string]condtionalChar{
	"d": {
		value:             "d",
		doNext:            "o",
		dontNext:          "o",
		shouldChangeState: false,
	},
	"o": {
		value:             "o",
		doNext:            "(",
		dontNext:          "n",
		shouldChangeState: false,
	},
	"(": {
		value:             "(",
		doNext:            ")",
		dontNext:          ")",
		shouldChangeState: false,
	},
	")": {
		value:             ")",
		doNext:            "d",
		dontNext:          "d",
		shouldChangeState: true,
	},
	"n": {
		value:             "n",
		doNext:            "d",
		dontNext:          "'",
		shouldChangeState: false,
	},
	"'": {
		value:             "'",
		doNext:            "d",
		dontNext:          "t",
		shouldChangeState: false,
	},
	"t": {
		value:             "t",
		doNext:            "d",
		dontNext:          "(",
		shouldChangeState: false,
	},
}

func checkDigit(character string) bool {
	if _, err := strconv.ParseInt(character, 10, 64); err == nil {
		return true
	} else {
		return false
	}

}

func solvePartOne() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanBytes)
	nextAcceptedChar := acceptMap["m"]
	firstNumber := ""
	secondNumber := ""
	result := 0

	for scanner.Scan() {
		character := scanner.Text()
		if nextAcceptedChar.value == "digit1" {
			if checkDigit(character) {
				firstNumber += character
				continue
			} else if character == nextAcceptedChar.nextAcceptedChar {
				nextAcceptedChar = acceptMap["digit2"]
				continue
			}
		} else if nextAcceptedChar.value == "digit2" {
			if checkDigit(character) {
				secondNumber += character
				continue
			} else if character == nextAcceptedChar.nextAcceptedChar {
				firstInt, err := strconv.Atoi(firstNumber)
				if err != nil {
					log.Fatal("Something is wrong with first number", firstNumber)
				}

				secondInt, err := strconv.Atoi(secondNumber)
				if err != nil {
					log.Fatal("Something is wrong with secondNumber number", secondNumber)
				}
				result += firstInt * secondInt
			}
		}
		if character != nextAcceptedChar.value {
			firstNumber = ""
			secondNumber = ""
			nextAcceptedChar = acceptMap["m"]
		} else {
			nextAcceptedChar = acceptMap[nextAcceptedChar.nextAcceptedChar]
		}
	}

	fmt.Println("Result is: ", result)
}

func solvePartTwo() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanBytes)
	nextAcceptedChar := acceptMap["m"]
	nextConditionalChar := conditionalMap["d"]
	firstNumber := ""
	secondNumber := ""
	result := 0
	enabled := true

	for scanner.Scan() {
		character := scanner.Text()
		if enabled {
			if nextAcceptedChar.value == "digit1" {
				if checkDigit(character) {
					firstNumber += character
					continue
				} else if character == nextAcceptedChar.nextAcceptedChar {
					nextAcceptedChar = acceptMap["digit2"]
					continue
				}
			} else if nextAcceptedChar.value == "digit2" {
				if checkDigit(character) {
					secondNumber += character
					continue
				} else if character == nextAcceptedChar.nextAcceptedChar {
					firstInt, err := strconv.Atoi(firstNumber)
					if err != nil {
						log.Fatal("Something is wrong with first number", firstNumber)
					}

					secondInt, err := strconv.Atoi(secondNumber)
					if err != nil {
						log.Fatal("Something is wrong with secondNumber number", secondNumber)
					}
					result += firstInt * secondInt
				}
			}
			if character != nextAcceptedChar.value {
				firstNumber = ""
				secondNumber = ""
				nextAcceptedChar = acceptMap["m"]
			} else {
				nextAcceptedChar = acceptMap[nextAcceptedChar.nextAcceptedChar]
			}
		}

		if nextAcceptedChar.value == "m" {
			if character == nextConditionalChar.value {

				if enabled {
					nextConditionalChar = conditionalMap[nextConditionalChar.dontNext]
				} else {
					nextConditionalChar = conditionalMap[nextConditionalChar.doNext]
				}
				if nextConditionalChar.shouldChangeState {
					enabled = !enabled
				}
			} else {
				nextConditionalChar = conditionalMap["d"]
			}
		}
	}

	fmt.Println("Result is: ", result)
}

func main() {
	isPartTwo := flag.Bool("partTwo", false, "Solve part two")
	flag.Parse()
	if *isPartTwo {
		solvePartTwo()
	} else {
		solvePartOne()
	}

}
