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
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	x := 1
	cycle := 1
	signalStrength := 0
	for _, line := range strings.Split(string(file), "\n") {

		if line == "" {
			continue
		}
		instruction := strings.Split(line, " ")
		switch instruction[0] {
		case "addx":
			value := stringToInt(instruction[1])
			for j := 0; j < 2; j++ {
				processCycle(&cycle, &signalStrength, x)
			}
			x += value
		case "noop":
			processCycle(&cycle, &signalStrength, x)
		}
	}
	fmt.Println(signalStrength)
	fmt.Println("========================")

	// Part 2
	x = 1
	cycle = 1
	for _, line := range strings.Split(string(file), "\n") {
		instruction := strings.Split(line, " ")
		switch instruction[0] {
		case "addx":
			value := stringToInt(instruction[1])
			for j := 0; j < 2; j++ {
				drawPixel(x, &cycle)
			}
			x += value
		case "noop":
			drawPixel(x, &cycle)
		}
	}
}

func processCycle(cycle *int, strength *int, x int) {
	if (*cycle+20)%40 == 0 {
		*strength += *cycle * x
	}
	*cycle++
}

func drawPixel(x int, cycle *int) {
	if *cycle >= x && *cycle <= x+2 {
		fmt.Printf("#")
	} else {
		fmt.Printf(".")
	}
	*cycle++
	if *cycle%41 == 0 {
		fmt.Printf("\n")
		*cycle = 1
	}
}
