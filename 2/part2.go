package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
	"math"
)

func main() {
	vals := readInput("input.txt")

	safeCount := 0
	for _, row := range vals {
		if checkRow(row) {
			safeCount++
		}
	}

	fmt.Println("Safe count:", safeCount)

}

func readInput(filename string) (vals [][]int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		var row []int
		for _, val := range line {
			num, err := strconv.Atoi(val)
			if err != nil {
				log.Fatal(err)
			}
			row = append(row, num)
		}
		vals = append(vals, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return vals
}

// check if a pair of values is safe
// -- the difference is between 1 and 3
// -- the difference is moving in the right direction
func checkIsSafe(prev int, next int, diffSign int) bool {
	absDiff := math.Abs(float64(prev - next))
	if absDiff >= 1 && absDiff <= 3 && (prev - next) * diffSign >= 0 {
		return true
	}

	return false
}


func checkRowHelper(row []int) bool {
	diffSign := 0
	if row[0] - row[1] > 0 {
		diffSign = 1
	} else {
		diffSign = -1
	}

	isSafe := true
	for i := 0; i < len(row) - 1; i++ {
		if !checkIsSafe(row[i], row[i+1], diffSign) {
			isSafe = false
			break
		}
	}

	return isSafe
}

func checkRow(row []int) bool {
	isSafe := checkRowHelper(row)

	if !isSafe {
		for i, _ := range row {
			var newRow = make([]int, len(row))
			copy(newRow, row)

			newRow = append(newRow[:i], newRow[i+1:]...)
			if checkRowHelper(newRow) {
				isSafe = true
				break
			}
		}
	}

	return isSafe
}
