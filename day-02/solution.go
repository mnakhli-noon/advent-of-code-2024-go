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

		inc := reportInt[0] < reportInt[1]
		shouldInc := true

		for i := 0; i < len(reportInt)-1; i++ {
			dif := reportInt[i+1] - reportInt[i]
			if abs(dif) < 1 || abs(dif) > 3 {
				shouldInc = false
				break
			} else if ((dif > 0) && !inc) || ((dif < 0) && inc) {
				shouldInc = false
				break
			}
		}
		if shouldInc {
			safe++
		}

	}

	fmt.Println("Safe: ", safe)

}
