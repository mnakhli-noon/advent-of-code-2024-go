package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func readInput() (map[int][]int, [][]int) {
	rules := make(map[int][]int)
	reports := [][]int{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Can't open file: ", err)
	}

	scanner := bufio.NewScanner(file)

	parseReports := false

	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			parseReports = true
			continue
		}
		if parseReports {
			numbers := strings.Split(line, ",")
			row := make([]int, len(numbers))
			for i, v := range numbers {
				row[i], _ = strconv.Atoi(v)
			}
			reports = append(reports, row)

		} else {
			numbers := strings.Split(line, "|")
			firstNumber, _ := strconv.Atoi(numbers[0])
			secondNumber, _ := strconv.Atoi(numbers[1])
			if len(rules[secondNumber]) > 0 {
				rules[secondNumber] = append(rules[secondNumber], firstNumber)
			} else {
				rules[secondNumber] = []int{firstNumber}
			}

		}

	}
	return rules, reports
}

func check(index int, report []int, numberRule []int) bool {
	for _, v := range report[:index] {
		if !slices.Contains(numberRule, v) {
			return false
		}
	}

	return true
}

func solvePartOne() {
	rules, reports := readInput()
	output := 0

	for _, report := range reports {
		for index, number := range report {

			if check(index, report, rules[number]) {
				if index == len(report)-1 {
					output += report[(index / 2)]
				}
			} else {
				break
			}
		}
	}
	fmt.Println(output)
}

func solvePartTwo() {
	rules, reports := readInput()

	sortFunc := func(a int, b int) int {
		if slices.Contains(rules[a], b) {
			return 1
		}
		if slices.Contains(rules[b], a) {
			return -1
		}
		return 0
	}

	output := 0
	for _, report := range reports {
		for index, number := range report {
			if !check(index, report, rules[number]) {
				slices.SortFunc(report, sortFunc)
				output += report[((len(report) - 1) / 2)]
				break
			}
		}
	}

	fmt.Println(output)

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
