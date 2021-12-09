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
	file, err := os.ReadFile("depths.txt")
	if err != nil {
		panic(err)
	}

	// Break the data into clean lines
	lines := strings.Split(string(file), "\n")
	depths := make([]int, 0, len(lines))

	// Range over lines and extract ints
	var larger int
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		number, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			panic(err)
		}

		if len(depths) > 0 && depths[len(depths)-1] < number {
			larger++
		}

		depths = append(depths, number)
	}

	fmt.Println(larger)
}
