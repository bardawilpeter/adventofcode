package main

import (
    "fmt"
    "os"
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
        fmt.Printf("Part 1 Result (%s): %s\n", file, resultPart1)

        resultPart2 := solvePart2(string(input))
        fmt.Printf("Part 2 Result (%s): %s\n", file, resultPart2)
    }
}

func solvePart1(input string) int {
    // Implement the solution for part 1 here
    return "Not implemented yet"
}

func solvePart2(input string) int {
    // Implement the solution for part 2 here
    return "Not implemented yet"
}
