package main

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	Inventory       []*big.Int
	InspectedTotal  int
	Operation       func(x *big.Int, y *big.Int)
	OperationNumber *big.Int
	Divisor         *big.Int
	ThrowIf         map[bool]int64
}

func stringToInt(s string) int64 {
	i, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		fmt.Printf("error converting string to int: %v", err)
		os.Exit(1)
	}
	return i
}

func Add(x *big.Int, y *big.Int) {
	x.Add(x, y)
}

func Mul(x *big.Int, y *big.Int) {
	x.Mul(x, y)
}

func Square(x *big.Int, y *big.Int) {
	x.Mul(x, x)
}

func processResults(monkeys []*Monkey) {
	var inspectedTotal []int
	for _, m := range monkeys {
		inspectedTotal = append(inspectedTotal, m.InspectedTotal)
	}
	sort.Ints(inspectedTotal)
	fmt.Println("Answer:", inspectedTotal[len(inspectedTotal)-1]*inspectedTotal[len(inspectedTotal)-2])
}

// Parse input data to populate Monkey objects
func initMonkeys(lines []string) ([]*Monkey, error) {
	// initialize monkeys' inventories
	var monkeys []*Monkey
	for i := 1; i < len(lines); i += 7 {
		monkey := &Monkey{}
		// Parsing inventory
		line := strings.TrimSpace(lines[i])
		re := regexp.MustCompile(`Starting items: ([0-9, ]+)`)
		s := re.FindStringSubmatch(line)
		if len(s) > 0 {
			inventory := re.FindStringSubmatch(line)[1]
			for _, item := range strings.Split(inventory, ",") {
				bi := new(big.Int)
				bi.SetString(strings.TrimSpace(item), 10)
				monkey.Inventory = append(monkey.Inventory, bi)
			}
		} else {
			return nil, errors.New("error during input data parsing")
		}

		// Operation
		re = regexp.MustCompile(`Operation: new = old ([0-9+* old]+)`)
		match := re.FindStringSubmatch(lines[i+1])
		if len(match) > 0 {
			operation := strings.Split(match[1], " ")
			num := new(big.Int)
			if operation[0] == "+" {
				monkey.Operation = Add
			} else {
				monkey.Operation = Mul
			}
			if operation[1] == "old" {
				monkey.Operation = Square
			} else {
				monkey.OperationNumber = num.SetInt64(stringToInt(operation[1]))
			}
		} else {
			return nil, errors.New("error during input data parsing")
		}

		// Test
		re = regexp.MustCompile(`Test: divisible by ([0-9]+)`)
		div := re.FindStringSubmatch(lines[i+2])
		if len(div) > 0 {
			monkey.Divisor = big.NewInt(stringToInt(strings.TrimSpace(div[1])))
			re = regexp.MustCompile(`[0-9+]`)
			monkey.ThrowIf = map[bool]int64{
				true:  stringToInt(re.FindStringSubmatch(lines[i+3])[0]),
				false: stringToInt(re.FindStringSubmatch(lines[i+4])[0]),
			}
		} else {
			return nil, errors.New("error during input data parsing")
		}

		monkeys = append(monkeys, monkey)
	}
	return monkeys, nil
}

func doRounds(rounds int, monkeys []*Monkey, f func(n *big.Int)) {
	for i := 1; i <= rounds; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.Inventory {
				monkey.Operation(item, monkey.OperationNumber)
				// Worry level normalization function
				if f != nil {
					f(item)
				}
				// Test item is dividable
				mod := new(big.Int)
				testResult := mod.Mod(item, monkey.Divisor).BitLen() == 0
				// Throw item to another monkey
				monkeys[monkey.ThrowIf[testResult]].Inventory = append(monkeys[monkey.ThrowIf[testResult]].Inventory, item)
				// Remove from the inventory
				monkey.Inventory = monkey.Inventory[1:]
				monkey.InspectedTotal += 1
			}
		}

		fmt.Printf(" === Round %d finished ===\n", i)
		for k, m := range monkeys {
			fmt.Printf("Monkey %d inspected items %v times\n", k, m.InspectedTotal)
		}
	}
}

func main() {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	lines := strings.Split(string(file), "\n")
	// Part 1
	monkeys, err := initMonkeys(lines)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	rounds := 20
	doRounds(rounds, monkeys, func(n *big.Int) {
		n.Div(n, big.NewInt(3))
	},
	)
	processResults(monkeys)

	// Part 2
	monkeys, err = initMonkeys(lines)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// To prevent numbers from growing up, we will use
	// a shared common denominator
	commonDivisor := big.NewInt(1)
	for _, m := range monkeys {
		commonDivisor.Mul(commonDivisor, m.Divisor)
	}
	rounds = 10000
	doRounds(rounds, monkeys, func(n *big.Int) {
		n.Mod(n, commonDivisor)
	})
	processResults(monkeys)
}
