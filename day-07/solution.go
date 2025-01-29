package main

import (
	"bufio"
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

func main() {
	solvePartOne()

}
