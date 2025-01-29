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

type Equations map[int][]int

func readInput() Equations {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal("Couldn't open 'input.txt' file ", err)
	}

	scanner := bufio.NewScanner(file)
	equations := Equations{}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")

		result, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal("Couldn't convert result ", parts[0], " to int")
		}

		numberStriped := strings.TrimSpace(parts[1])
		numbersStr := strings.Split(numberStriped, " ")

		array := []int{}
		for _, num := range numbersStr {
			number, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal("Couldn't convert ", num, " to int")
			}
			array = append(array, number)
		}
		equations[result] = array
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
	aStr := ""
	if a != 0 {
		aStr = strconv.Itoa(a)
	}
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
