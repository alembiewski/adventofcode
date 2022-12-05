package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	filePath := os.Args[1]

	fmt.Printf("Reading input from %s\n", filePath)
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	// Total calories for each elf
	var totalCalories []int
	// Calories counter for the current elf
	var calories int

	for _, line := range strings.Split(string(file), "\n") {
		// Empty string appeared - record calculated calories for the current Elf
		if line == "" {
			totalCalories = append(totalCalories, calories)
			calories = 0
			continue
		}
		caloriesInt, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			fmt.Println("Error during conversion", err)
			return
		}
		calories += caloriesInt
	}
	// Sort calories
	sort.Ints(totalCalories)
	totalElves := len(totalCalories)
	fmt.Printf("Answer (part 1): %d\n", totalCalories[totalElves-1])
	// Calculate total calories for top 3 elves
	topElves := 3
	topElvesCalories := totalCalories[totalElves-topElves:]
	topElvesCaloriesTotal := 0
	for i := range topElvesCalories {
		topElvesCaloriesTotal += topElvesCalories[i]
	}
	fmt.Printf("Answer (part 2): %d\n", topElvesCaloriesTotal)
}
