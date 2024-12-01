#!/bin/bash

YEAR=$(date +%Y)

# Create the base directory for the year
BASE_DIR="aoc-${YEAR}"
mkdir -p "$BASE_DIR"

# Create folders and files for each day
for day in {1..25}; do
    DAY_DIR="$BASE_DIR/day$(printf "%02d" $day)"
    mkdir -p "$DAY_DIR"
    
    # Generate the main.go file with scaffolding for both parts
    cat > "$DAY_DIR/main.go" <<EOL
package main

import (
    "fmt"
    "os"
)

func main() {
    files := []string{"example_input.txt", "input.txt"}

    for _, file := range files {
        fmt.Printf("Processing file: %s\\n", file)
        input, err := os.ReadFile(file)
        if err != nil {
            fmt.Printf("Error reading %s: %s\\n", file, err)
            continue
        }

        resultPart1 := solvePart1(string(input))
        fmt.Printf("Part 1 Result (%s): %s\\n", file, resultPart1)

        resultPart2 := solvePart2(string(input))
        fmt.Printf("Part 2 Result (%s): %s\\n", file, resultPart2)
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
EOL

    echo "// Add example input here" > "$DAY_DIR/example_input.txt"
    echo "// Add real input here" > "$DAY_DIR/input.txt"
done
