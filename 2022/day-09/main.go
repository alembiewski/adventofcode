package main

import (
	"fmt"
	"math"
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

var movements = map[string][]int{
	"L": {-1, 0},
	"R": {1, 0},
	"U": {0, -1},
	"D": {0, 1},
}

func snake(numberOfKnots int, input []string) int {
	// initialize knots
	var knots [][2]int
	for i := 0; i < numberOfKnots; i++ {
		knots = append(knots, [2]int{0, 0})
	}
	// store visited points in a set
	visited := map[[2]int]struct{}{
		{0, 0}: {},
	}
	for _, line := range input {
		if line == "" {
			continue
		}

		command := strings.Split(line, " ")
		direction := command[0]
		steps := stringToInt(command[1])
		move := movements[direction]
		for i := 0; i < steps; i++ {
			// head moves
			knots[0] = [2]int{knots[0][0] + move[0], knots[0][1] + move[1]}
			// tail follow
			for j := 1; j < numberOfKnots; j++ {
				dx, dy := knots[j-1][0]-knots[j][0], knots[j-1][1]-knots[j][1]
				absDx := math.Abs(float64(dx))
				absDy := math.Abs(float64(dy))
				if absDx > 1 || absDy > 1 {
					tx, ty := knots[j][0], knots[j][1]
					if absDx > 0 {
						tx = knots[j][0] + int(math.Copysign(1, float64(dx)))
					}
					if absDy > 0 {
						ty = knots[j][1] + int(math.Copysign(1, float64(dy)))
					}
					knots[j] = [2]int{tx, ty}
				}
			}
			visited[knots[len(knots)-1]] = struct{}{}
		}
	}
	return len(visited)
}

func main() {

	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	input := strings.Split(string(file), "\n")
	// Part 1
	fmt.Println(snake(2, input))
	// Part 2
	fmt.Println(snake(10, input))
}
