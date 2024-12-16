package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
	"bufio"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var list1 []int
	var list2 []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splits := strings.Split(scanner.Text(), "   ")

		int1, err := strconv.Atoi(splits[0])
		int2, err := strconv.Atoi(splits[1])
		if err != nil {
			log.Fatal(err)
		}

		list1 = append(list1, int1)
		list2 = append(list2, int2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	tot := 0
	for _, val1 := range list1 {
		occurrences := 0
		for _, val2 := range list2 {
			if val1 == val2 {
				occurrences++
			}
		}

		tot += val1 * occurrences
	}

	fmt.Println(tot)

	
}



