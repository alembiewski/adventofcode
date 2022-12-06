package main

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	contents, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	input := string(contents)
	// Part 1
	result, err := processInput(input, 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// Part 2
	result, err = processInput(input, 14)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

}

func processInput(input string, offset int) (int, error) {
	for i := 0; i < len(input)-offset; i++ {
		chunk := strings.Split(input[i:i+offset], "")
		sort.Strings(chunk)
		for j := 0; j < len(chunk)-1; j++ {
			if chunk[j] == chunk[j+1] {
				break
			} else {
				if j == len(chunk)-2 {
					return i + offset, nil
				} else {
					continue
				}
			}
		}
	}
	return -1, errors.New("no matches found")
}
