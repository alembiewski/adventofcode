package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

const (
	DrawPoints = 3
	WinPoints  = 6
	LossPoints = 0
)

var ab = []string{"A", "B"}
var bc = []string{"B", "C"}
var ac = []string{"A", "C"}

var scoreMapping = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
}
var decodeFigure = map[string]string{
	"X": "A",
	"Y": "B",
	"Z": "C",
}

// Value array elements store stronger and weaker hands respectively, e.g. "A" loses to "B" and beats "C"
var figuresMapping = map[string][]string{
	"A": {"B", "C"},
	"B": {"C", "A"},
	"C": {"A", "B"},
}

func Game(opponent string, player string) int {
	if player == opponent {
		return DrawPoints
	}

	var game = []string{player, opponent}
	sort.Strings(game)
	if (reflect.DeepEqual(game, ab) && player == "B") ||
		(reflect.DeepEqual(game, bc) && player == "C") ||
		(reflect.DeepEqual(game, ac) && player == "A") {
		return WinPoints
	}
	return LossPoints
}

func main() {
	var player string
	var opponent string

	filePath := os.Args[1]

	fmt.Printf("Reading input from %s\n", filePath)
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	// Part 1
	var totalScore int
	for _, line := range strings.Split(string(file), "\n") {
		combination := strings.Split(strings.TrimSpace(line), " ")
		if len(combination) == 2 {
			opponent = combination[0]
			player = decodeFigure[combination[1]]
			totalScore += Game(opponent, player) + scoreMapping[player]
		}
	}
	fmt.Printf("Total player score: %d\n", totalScore)
	// Part 2
	totalScore = 0
	for _, line := range strings.Split(string(file), "\n") {
		combination := strings.Split(strings.TrimSpace(line), " ")
		if len(combination) == 2 {
			opponent = combination[0]
			player = combination[1]
			// X - player needs to lose
			// Y - draw
			// Z - player needs to win
			switch player {
			case "Y":
				player = opponent
			case "X":
				player = figuresMapping[opponent][1]
			case "Z":
				player = figuresMapping[opponent][0]
			}

			totalScore += Game(opponent, player) + scoreMapping[player]
		}
	}
	fmt.Printf("Total player score: %d\n", totalScore)

}
