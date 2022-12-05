package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("error converting string to int: %v", err)
		os.Exit(1)
	}
	return i
}

func main() {

	filePath := os.Args[1]

	fmt.Printf("Reading input from %s\n", filePath)
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	// Part 1 & 2
	var fullOverlapCount int
	var partialOverlapCount int
	for _, line := range strings.Split(string(file), "\n") {
		if line != "" {
			pairs := strings.Split(line, ",")
			a, b := strings.Split(pairs[0], "-"), strings.Split(pairs[1], "-")
			a0, a1, b0, b1 := stringToInt(a[0]), stringToInt(a[1]), stringToInt(b[0]), stringToInt(b[1])
			if (a0 >= b0 && a1 <= b1) || (b0 >= a0 && b1 <= a1) {
				fullOverlapCount += 1
			}
			if (a0 >= b0 && a0 <= b1) || (b0 >= a0 && b0 <= a1) {
				partialOverlapCount += 1
			}
		}
	}
	fmt.Printf("Ranges fully overlap: %d\n", fullOverlapCount)
	fmt.Printf("Ranges partial overlap: %d\n", partialOverlapCount)
}
