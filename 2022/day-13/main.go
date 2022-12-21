package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"sort"
	"strings"
)

func parseList(list string) []interface{} {
	var nums []interface{}

	if err := json.Unmarshal([]byte(list), &nums); err != nil {
		log.Println(err)
	}
	return nums
}

func compare(l1 interface{}, l2 interface{}) int {
	fmt.Println("Comparing: ", l1, "vs", l2)
	var result int
	if reflect.TypeOf(l1).Kind() == reflect.Float64 && reflect.TypeOf(l2).Kind() == reflect.Float64 {
		if l1.(float64) < l2.(float64) {
			return -1
		}
		if l1.(float64) == l2.(float64) {
			return 0
		}
		return 1
	}
	if reflect.TypeOf(l1).Kind() == reflect.Float64 {
		l1 = []interface{}{l1}
	}

	if reflect.TypeOf(l2).Kind() == reflect.Float64 {
		l2 = []interface{}{l2}
	}
	min := math.Min(float64(len(l1.([]interface{}))), float64(len(l2.([]interface{}))))
	for i := 0; i < int(min); i++ {
		result = compare(l1.([]interface{})[i], l2.([]interface{})[i])
		if result == 0 {
			continue
		}
		return result
	}
	if len(l1.([]interface{})) < len(l2.([]interface{})) {
		return -1
	} else if len(l1.([]interface{})) > len(l2.([]interface{})) {
		return 1
	} else {
		return 0
	}
}

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	pairs := strings.Split(string(file), "\n\n")

	var pairsInOrder int
	for i, pair := range pairs {
		split := strings.Split(pair, "\n")
		l1, l2 := parseList(split[0]), parseList(split[1])
		fmt.Println(l1)
		fmt.Println(l2)
		if compare(l1, l2) == -1 {
			pairsInOrder += i + 1
		}

		fmt.Println("==========")
	}

	fmt.Println(pairsInOrder)

	// Part 2
	dividerPacket1 := []interface{}{[]interface{}{float64(2)}}
	dividerPacket2 := []interface{}{[]interface{}{float64(6)}}
	lines := []interface{}{
		dividerPacket1,
		dividerPacket2,
	}
	for _, line := range strings.Split(string(file), "\n") {
		if line == "" {
			continue
		}
		lines = append(lines, parseList(line))
	}
	sort.SliceStable(lines, func(i, j int) bool {
		if compare(lines[i], lines[j]) == -1 {
			return true
		}
		return false
	})
	decoderKey := 1
	for i, line := range lines {
		if reflect.DeepEqual(line, dividerPacket1) || reflect.DeepEqual(line, dividerPacket2) {
			decoderKey *= i + 1
		}
	}
	fmt.Println(decoderKey)
}
