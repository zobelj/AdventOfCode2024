package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	tot := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		valsStr := strings.Split(scanner.Text(), " ")
		vals := make([]int, len(valsStr))

		// cast each val to an int
		for idx, val := range valsStr {
			num, _ := strconv.Atoi(val)
			vals[idx] = num
		}

		var safe int = 0
		var dir int = 0

		// check if vals is montonically increasing or decreasing
		diff1 := vals[0] - vals[1]
		if diff1 == 0 {
			safe = -1
		} else if diff1 > 0 {
			dir = 1
		} else {
			dir = -1
		}

		// early skip if this is already not safe
		if safe == -1 {
			continue
		}

		// iterate over each val. check each comparison matches `dir`
		for idx, _ := range vals[0:len(vals)-1] {
			diff := vals[idx] - vals[idx + 1]

			// if the diff is not between dir * 1 and dir * 3, continue (not safe)
			if dir == 1 && diff >= 1 && diff <= 3 {
				safe = 1
			} else if dir == -1 && diff <= -1 && diff >= -3 {
				safe = 1
			} else {
				safe = -1
				break
			}
		}

		// if its safe, increment tot
		if safe == 1 {
			tot++
		}

	}

	fmt.Println("total:", tot)


}
