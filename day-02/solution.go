package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func convertToArrayOfInt(array []string) ([]int, error) {
	newArray := make([]int, len(array))
	for i, v := range array {
		intV, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("Couldn't convert value to int %s - Reason: %s", v, err)
		}

		newArray[i] = intV
	}

	return newArray, nil
}

func abs(number int) int {
	if number >= 0 {
		return number
	} else {
		return -number
	}
}

func isItSafe(x int, y int, inc bool) bool {
	dif := y - x

	if (abs(dif) < 1 || abs(dif) > 3) || ((dif > 0) && !inc) || ((dif < 0) && inc) {
		return false
	}

	return true
}

func checkIfReportSafe(report []int) bool {
	inc := report[1] > report[0]
	for i := 0; i < len(report)-1; i++ {
		if isItSafe(report[i], report[i+1], inc) {
			continue
		} else {
			return false
		}
	}
	return true
}

func removeItemAt(array []int, index int) []int {
	result := make([]int, 0, len(array)-1)
	result = append(result, array[:index]...)
	result = append(result, array[index+1:]...)

	return result
}
func checkIfReportSafeWithTolerence(report []int) bool {
	for i := 0; i < len(report); i++ {
		newArray := removeItemAt(report, i)
		if checkIfReportSafe(newArray) {
			return true
		}
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	safe := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		report := strings.Fields(scanner.Text())
		reportInt, err := convertToArrayOfInt(report)
		if err != nil {
			log.Fatal(err)
		}

		if checkIfReportSafe(reportInt) {
			safe++
		} else if checkIfReportSafeWithTolerence(reportInt) {
			safe++
		}

	}

	fmt.Println("Safe: ", safe)

}
