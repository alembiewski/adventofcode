package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	dr = [4]int{1, -1, 0, 0}
	dc = [4]int{0, 0, 1, -1}
)

type Queue struct {
	elements [][3]int
}

func (s *Queue) Enqueue(item [3]int) {

	s.elements = append(s.elements, item)
}

func (s *Queue) Dequeue() [3]int {
	if len(s.elements) == 0 {
		return [3]int{}
	}
	a := s.elements[0]
	s.elements = s.elements[1:len(s.elements)]
	return a
}

func main() {

	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	lines := strings.Split(string(file), "\n")

	var queue = &Queue{}
	// Part 1
	grid, target := buildGrid(lines, queue, "S")
	fmt.Println(findShortestPath(grid, queue, target))

	// Part 2
	grid, target = buildGrid(lines, queue, "a")
	fmt.Println(findShortestPath(grid, queue, target))
}

func buildGrid(lines []string, queue *Queue, startingPoint string) (map[[2]int]int, [2]int) {
	var grid = make(map[[2]int]int)
	var target [2]int
	for i, row := range lines {
		for j, column := range row {
			value := fmt.Sprintf("%c", column)
			if value == startingPoint {
				queue.Enqueue([3]int{i, j, 0})
				grid[[2]int{i, j}] = int('a')
			} else if value == "E" {
				target = [2]int{i, j}
				grid[[2]int{i, j}] = int('z')
			} else {
				grid[[2]int{i, j}] = int(column)
			}
		}
	}
	return grid, target
}

func findShortestPath(grid map[[2]int]int, queue *Queue, target [2]int) int {
	var visited = make(map[[2]int]struct{})
	for len(queue.elements) > 0 {
		position := queue.Dequeue()
		r, c, d := position[0], position[1], position[2]

		if [2]int{r, c} == target {
			return d
		}
		_, ok := visited[[2]int{r, c}]
		if ok {
			continue
		}
		visited[[2]int{r, c}] = struct{}{}

		for i := 0; i < 4; i++ {
			if grid[[2]int{r + dr[i], c + dc[i]}] <= grid[[2]int{r, c}]+1 {
				queue.Enqueue([3]int{r + dr[i], c + dc[i], d + 1})
			}
		}
	}
	return -1
}
