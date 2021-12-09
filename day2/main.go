package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// First get all the depths from
	// https://adventofcode.com/2021/day/1/input
	file, err := os.ReadFile("../day1/depths.txt")
	if err != nil {
		panic(err)
	}

	// Break the data into clean lines
	lines := strings.Split(string(file), "\n")
	depths := make([]int, 0, len(lines))

	// Range over lines and extract ints
	var larger, total int
	for _, line := range lines {
		// count++
		if len(line) == 0 {
			continue
		}

		number, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			panic(err)
		}

		depths = append(depths, number)

		// skip the first 3
		if len(depths) < 3 {
			continue
		}

		if total > 0 && total < sum(depths[len(depths)-3:len(depths)-0]) {
			fmt.Printf("LARGER: (%d, %d, %d) %d > %d\n", depths[len(depths)-3], depths[len(depths)-2], depths[len(depths)-1], sum(depths[len(depths)-3:len(depths)-0]), total)
			larger++
		} else {
			fmt.Printf("SMALLER: (%d, %d, %d) %d < %d\n", depths[len(depths)-3], depths[len(depths)-2], depths[len(depths)-1], sum(depths[len(depths)-3:len(depths)-0]), total)
		}

		total = sum(depths[len(depths)-3 : len(depths)-0])
	}

	fmt.Println(larger)
}

func sum(slice []int) int {
	total := 0
	for _, i := range slice {
		total += i
	}
	return total
}
