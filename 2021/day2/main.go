package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// First read the movements from
	// https://adventofcode.com/2021/day/2/input
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// Break the data into clean lines
	lines := strings.Split(string(file), "\n")
	// movements := make([]int, 0, len(lines))

	// Range over lines
	var horizontal, depth, aim int
	for _, line := range lines {
		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		// Extract the direction and value for reach line
		var direction, valueStr string
		var value int
		splitLine := strings.Split(strings.TrimSpace(line), " ")
		direction = splitLine[0]
		valueStr = splitLine[len(splitLine)-1]

		value, err := strconv.Atoi(valueStr)
		if err != nil {
			panic(err)
		}

		// Interpret each line as forward, down or up
		switch direction {
		case "forward":
			fmt.Printf("\nFORWARD\nhorizontal: %d + %d = %d\naim: %d\ndepth: %d + %d = %d\n\n", horizontal, value, horizontal+value, aim, depth, aim*value, depth+aim*value)
			horizontal += value
			depth += (aim * value)
		case "up":
			fmt.Printf("\nUP\ndepth: %d - %d = %d\naim: %d - %d = %d\n", depth, value, depth-value, aim, value, aim-value)
			// depth -= value
			aim -= value
		case "down":
			fmt.Printf("\nDOWN\ndepth: %d + %d = %d\naim: %d + %d = %d\n", depth, value, depth+value, aim, value, aim+value)
			// depth += value
			aim += value
		}
	}

	fmt.Printf("TOTALS\nhorizontal: %d\ndepth: %d\naim: %d\nsum: %d", horizontal, depth, aim, horizontal*depth)
}
