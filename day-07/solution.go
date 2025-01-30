package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type EquationSet map[int][]int

func parseLine(line string) (int, []int, error) {
	parts := strings.Split(line, ":")
	if len(parts) != 2 {
		return 0, nil, fmt.Errorf("Invalid format: expected 'result: numbers' in line %q", line)
	}
	result, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, nil, fmt.Errorf("Invalid result %q: %v", parts[0], err)
	}

	numbersString := strings.Fields(parts[1])
	numbers := make([]int, 0, len(numbersString))

	for _, numberString := range numbersString {
		number, err := strconv.Atoi(numberString)
		if err != nil {
			return 0, nil, fmt.Errorf("Invalid number %q: %v", numberString, err)
		}

		numbers = append(numbers, number)
	}

	return result, numbers, nil
}

func readInput() EquationSet {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal("Couldn't open 'input.txt' file ", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	equations := EquationSet{}

	for scanner.Scan() {
		result, numbers, err := parseLine(scanner.Text())
		if err != nil {
			log.Fatal("Error parsing line: ", err)
		}

		equations[result] = numbers
	}

	return equations
}

func canGetResultFromNumbers(result int, numbers []int) bool {
	if len(numbers) == 0 {
		return false
	}

	if len(numbers) == 1 {
		return numbers[0] == result
	}

	lastNumber := numbers[len(numbers)-1]
	if result%lastNumber == 0 {
		return canGetResultFromNumbers(result-lastNumber, numbers[:len(numbers)-1]) || canGetResultFromNumbers(result/lastNumber, numbers[:len(numbers)-1])
	} else {
		return canGetResultFromNumbers(result-lastNumber, numbers[:len(numbers)-1])
	}
}

func solvePartOne() {
	equations := readInput()
	sum := 0
	for result, numbers := range equations {
		if canGetResultFromNumbers(result, numbers) {
			sum += result
		}
	}
	fmt.Println(sum)
}

func concat(a, b int) int {
	if a == 0 {
		return b
	}
	aStr := strconv.Itoa(a)
	bStr := strconv.Itoa(b)

	number, _ := strconv.Atoi(aStr + bStr)

	return number
}

func canGetResultFromNumbers2(result int, numbers []int, acc int) bool {
	if len(numbers) == 0 {
		return acc == result
	}
	firstNumber := numbers[0]
	if len(numbers) == 1 {
		return acc*firstNumber == result || acc+firstNumber == result || concat(acc, firstNumber) == result
	}

	addition := canGetResultFromNumbers2(result, numbers[1:], acc+firstNumber)
	multiplication := false
	if acc == 0 {
		multiplication = canGetResultFromNumbers2(result, numbers[1:], firstNumber)
	} else {
		multiplication = canGetResultFromNumbers2(result, numbers[1:], acc*firstNumber)
	}
	concat := canGetResultFromNumbers2(result, numbers[1:], concat(acc, firstNumber))

	return addition || multiplication || concat
}

func solvePartTwo() {
	equations := readInput()
	sum := 0

	for result, numbers := range equations {
		if canGetResultFromNumbers2(result, numbers, 0) {
			sum += result
		}
	}
	fmt.Println(sum)

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
