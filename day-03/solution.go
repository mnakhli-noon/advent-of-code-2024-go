package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type acceptedChar struct {
	value            string
	nextAcceptedChar string
	shouldEndAfter   bool
}

func checkDigit(character string) bool {
	if _, err := strconv.ParseInt(character, 10, 64); err == nil {
		return true
	} else {
		return false
	}

}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	acceptMap := map[string]acceptedChar{
		"m": {
			value:            "m",
			nextAcceptedChar: "u",
			shouldEndAfter:   false,
		},
		"u": {
			value:            "u",
			nextAcceptedChar: "l",
			shouldEndAfter:   false,
		},
		"l": {
			value:            "l",
			nextAcceptedChar: "(",
			shouldEndAfter:   false,
		},
		"(": {
			value:            "(",
			nextAcceptedChar: "digit1",
			shouldEndAfter:   false,
		},
		"digit1": {
			value:            "digit1",
			nextAcceptedChar: ",",
			shouldEndAfter:   false,
		},
		"digit2": {
			value:            "digit2",
			nextAcceptedChar: ")",
			shouldEndAfter:   false,
		},
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
