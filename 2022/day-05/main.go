package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	TotalStacks = 9
)

type Stack struct {
	elements []string
}

func (s *Stack) Add(item string) {
	s.elements = append(s.elements, item)
}

func (s *Stack) Pop(n int) ([]string, error) {
	if n > len(s.elements) {
		return nil, errors.New("not enough items")
	}
	a := s.elements[len(s.elements)-n:]
	s.elements = s.elements[:len(s.elements)-n]
	return a, nil
}

type Instruction struct {
	Count int
	From  int
	To    int
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("error converting string to int: %v", err)
		os.Exit(1)
	}
	return i
}

func GetInput() ([]Stack, []Instruction) {
	filePath := os.Args[1]

	fmt.Printf("Reading input from %s\n", filePath)
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	var cratesConfiguration []string
	var movingInstructions []string
	for i, line := range strings.Split(string(file), "\n") {
		if i < 8 {
			cratesConfiguration = append(cratesConfiguration, line)
		}
		if i >= 10 {
			movingInstructions = append(movingInstructions, line)
		}
	}

	var stacks []Stack
	for i := 0; i < TotalStacks; i++ {
		s := &Stack{}
		for j := len(cratesConfiguration) - 1; j >= 0; j-- {
			startPos := 4 * i
			endPos := startPos + 3
			if len(cratesConfiguration[j]) < endPos {
				break
			}
			elem := cratesConfiguration[j][startPos:endPos]
			if strings.TrimSpace(elem) != "" {
				s.Add(cratesConfiguration[j][startPos:endPos])
			}
		}
		stacks = append(stacks, *s)
	}
	fmt.Println("Stacks configuration:")
	for i, stack := range stacks {
		fmt.Printf("%d %v\n", i+1, stack.elements)
	}
	var instructions []Instruction
	for _, i := range movingInstructions {
		if len(strings.TrimSpace(i)) > 0 {
			a := strings.Split(i, " ")
			instructions = append(instructions, Instruction{
				Count: stringToInt(a[1]),
				From:  stringToInt(a[3]) - 1,
				To:    stringToInt(a[5]) - 1},
			)
		}
	}
	return stacks, instructions
}

func GetAnswer(stacks []Stack) {
	fmt.Printf("Answer: ")
	for _, stack := range stacks {
		fmt.Printf("%v", stack.elements[len(stack.elements)-1][1:2])
	}
	fmt.Println()
}

func main() {
	// Part 1
	stacks, instructions := GetInput()
	for _, instruction := range instructions {
		a, err := stacks[instruction.From].Pop(instruction.Count)
		if err != nil {
			fmt.Printf("error during instruction execution %v: %v", instruction, err)
			os.Exit(1)
		}

		for i := len(a) - 1; i >= 0; i-- {
			stacks[instruction.To].Add(a[i])
		}
	}
	GetAnswer(stacks)

	fmt.Println("-------------------------")
	// Part 2
	stacks, instructions = GetInput()
	for _, instruction := range instructions {
		a, err := stacks[instruction.From].Pop(instruction.Count)
		if err != nil {
			fmt.Printf("error during instruction execution %v: %v", instruction, err)
			os.Exit(1)
		}

		for _, i := range a {
			stacks[instruction.To].Add(i)
		}
	}
	GetAnswer(stacks)
}
