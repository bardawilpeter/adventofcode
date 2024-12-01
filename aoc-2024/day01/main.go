package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	files := []string{"example_input.txt", "input.txt"}

	for _, file := range files {
		fmt.Printf("Processing file: %s\n", file)
		input, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading %s: %s\n", file, err)
			continue
		}

		resultPart1 := solvePart1(string(input))
		fmt.Printf("Part 1 Result (%s): %d\n", file, resultPart1)

		resultPart2 := solvePart2(string(input))
		fmt.Printf("Part 2 Result (%s): %d\n", file, resultPart2)
	}
}

func solvePart1(input string) int {
	column1, column2 := parseInput(input)

	sort.Ints(column1)
	sort.Ints(column2)

	totalDifference := 0
	for i := 0; i < len(column1) && i < len(column2); i++ {
		totalDifference += abs(column1[i] - column2[i])
	}

	return totalDifference
}

func solvePart2(input string) int {
    column1, column2 := parseInput(input)

    frequency := make(map[int]int)
    for _, num := range column2 {
        frequency[num]++
    }

    totalScore := 0
    for _, num := range column1 {
        if count, exists := frequency[num]; exists {
            totalScore += num * count
        }
    }

    return totalScore
}

func parseInput(input string) ([]int, []int) {
	var column1, column2 []int
	lines := strings.Split(strings.TrimSpace(input), "\n")

	for _, line := range lines {
		parts := strings.Fields(line)

		num1, err1 := strconv.Atoi(parts[0])
		num2, err2 := strconv.Atoi(parts[1])
		if err1 == nil && err2 == nil {
			column1 = append(column1, num1)
			column2 = append(column2, num2)
		}
	}

	return column1, column2
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}
