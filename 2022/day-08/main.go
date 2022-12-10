package main

import (
	"fmt"
	"os"
	"sort"
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
	trees, done := parseInput()
	if done {
		return
	}
	var count int
	var scenicScores []int
	for i := 1; i < len(trees)-1; i++ {
		for j := 1; j < len(trees[i])-1; j++ {
			tree := trees[i][j]
			var surroundingTrees [][]int

			treesOnTheLeft := trees[i][0:j]
			treesOnTheRight := trees[i][j+1 : len(trees[i])]

			var treesAbove []int
			for k := 0; k < i; k++ {
				treesAbove = append(treesAbove, trees[k][j])
			}
			var treesBelow []int
			for k := i + 1; k < len(trees); k++ {
				treesBelow = append(treesBelow, trees[k][j])
			}

			surroundingTrees = append(surroundingTrees, treesOnTheLeft, treesAbove, treesOnTheRight, treesBelow)

			for _, side := range surroundingTrees {
				if isVisible(tree, side) {
					count += 1
					break
				}
			}

			scenicScore := calculateScore(tree, reverse(treesOnTheLeft))
			scenicScore *= calculateScore(tree, treesOnTheRight)
			scenicScore *= calculateScore(tree, reverse(treesAbove))
			scenicScore *= calculateScore(tree, treesBelow)

			scenicScores = append(scenicScores, scenicScore)
		}
	}
	fmt.Println(count + 4*len(trees) - 4)
	sort.Ints(scenicScores)
	fmt.Println(scenicScores[len(scenicScores)-1])
}

func parseInput() ([][]int, bool) {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return nil, true
	}

	var trees [][]int
	for _, line := range strings.Split(string(file), "\n") {
		if line == "" {
			continue
		}
		var row []int
		for _, tree := range strings.Split(line, "") {
			row = append(row, stringToInt(tree))
		}
		trees = append(trees, row)
	}
	return trees, false
}

func isVisible(treeHeight int, trees []int) bool {
	for _, tree := range trees {
		if tree >= treeHeight {
			return false
		}
	}
	return true
}

func calculateScore(tree int, trees []int) int {
	counter := 0
	for i := 0; i < len(trees); i++ {
		counter += 1
		if trees[i] >= tree {
			break
		}
	}
	return counter
}

func reverse(arr []int) []int {
	var reversed []int
	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}
	return reversed
}
