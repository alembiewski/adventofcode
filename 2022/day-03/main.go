package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var priorities string
	for ch := 'a'; ch <= 'z'; ch++ {
		priorities += string(ch)
	}
	for ch := 'A'; ch <= 'Z'; ch++ {
		priorities += string(ch)
	}

	filePath := os.Args[1]

	fmt.Printf("Reading input from %s\n", filePath)
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	var sumPriorities int
	for _, items := range strings.Split(string(file), "\n") {
		items = strings.TrimSpace(items)
		first := items[0 : len(items)/2]
		second := items[len(items)/2:]
		for _, item := range first {
			itemStr := string(item)
			if strings.Index(second, itemStr) != -1 {
				sumPriorities += strings.Index(priorities, itemStr) + 1
				break
			}
		}
	}
	fmt.Printf("Total sum of priorities: %d", sumPriorities)

	// Part 2
	rucksacks := strings.Split(string(file), "\n")
	groupSize := 3
	sumPriorities = 0
	for i := 0; i < len(rucksacks); i += groupSize {
		j := i + groupSize
		if j > len(rucksacks) {
			j = len(rucksacks)
		}
		group := rucksacks[i:j]
		if len(group) > 0 {
			sort.Slice(group, func(i, j int) bool {
				return len(group[i]) < len(group[j])
			})
			for _, item := range group[0] {
				itemStr := string(item)
				if strings.Index(group[1], itemStr) != -1 && strings.Index(group[2], itemStr) != -1 {
					sumPriorities += strings.Index(priorities, itemStr) + 1
					break
				}
			}
		}
	}
	fmt.Printf("Total sum: %d", sumPriorities)
}
