package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"strconv"
	"sort"
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

	// sort the two lists
	sort.Ints(list1)
	sort.Ints(list2)

	// iterate over each pair to calculate distance
	var tot int = 0
	for idx, _ := range list1 {
		distance := list1[idx] - list2[idx]
		tot += max(distance, -distance)
	}
	
	fmt.Println("total distance:", tot)

}

