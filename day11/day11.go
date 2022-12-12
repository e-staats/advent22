package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	f, err := os.Open("./day11.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	lines = append(lines, "")

	start := time.Now()
	fmt.Printf("Part 1: %d\n", part1(lines))
	duration := time.Since(start)
	fmt.Println(duration)

	fmt.Printf("\n")

	start = time.Now()
	result := part2(lines)
	fmt.Printf("Part 2: %d\n", result)
	duration = time.Since(start)
	fmt.Println(duration)
}

func part1(lines []string) int {
	var chunk []string
	var monkeys []*Monkey
	for _, line := range lines {
		if line != "" {
			chunk = append(chunk, line)
		} else {
			monkey := parseMonkey(chunk)
			monkeys = append(monkeys, monkey)
			chunk = nil
		}
	}

	for round := 1; round <= 20; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.inventory {
				newWorry := calcWorry(item, monkey.op, monkey.opVal) / 3
				if newWorry%monkey.testVal == 0 {
					monkeys[monkey.trueDest].inventory = append(monkeys[monkey.trueDest].inventory, newWorry)
				} else {
					monkeys[monkey.falseDest].inventory = append(monkeys[monkey.falseDest].inventory, newWorry)
				}
				monkey.inspectCount++
			}
			monkey.inventory = nil
		}
		fmt.Printf("After round %d:\n", round)
		for _, m := range monkeys {
			fmt.Print(m.inventory)
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}

	max1 := 0
	max2 := 0
	for _, monkey := range monkeys {
		fmt.Printf("Monkey %d inspected %d items\n", monkey.id, monkey.inspectCount)
		if monkey.inspectCount >= max1 {
			max2 = max1
			max1 = monkey.inspectCount
		} else if monkey.inspectCount > max2 {
			max2 = monkey.inspectCount
		}
	}

	fmt.Printf("%d * %d = %d\n", max1, max2, max1*max2)

	return max1 * max2
}

func part2(lines []string) int {
	var chunk []string
	var monkeys []*Monkey
	for _, line := range lines {
		if line != "" {
			chunk = append(chunk, line)
		} else {
			monkey := parseMonkey(chunk)
			monkeys = append(monkeys, monkey)
			chunk = nil
		}
	}

	modConst := 1
	for _, monkey := range monkeys {
		modConst *= monkey.testVal
	}
	for round := 1; round <= 10000; round++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.inventory {
				newWorry := calcWorry(item, monkey.op, monkey.opVal)
				testResult := newWorry % monkey.testVal
				if testResult == 0 {
					monkeys[monkey.trueDest].inventory = append(monkeys[monkey.trueDest].inventory, newWorry%modConst)
				} else {
					monkeys[monkey.falseDest].inventory = append(monkeys[monkey.falseDest].inventory, newWorry%modConst)
				}
				monkey.inspectCount++
			}
			monkey.inventory = nil
		}
		if round%1000 == 0 {
			fmt.Printf("After round %d:\n", round)
			for _, m := range monkeys {
				fmt.Print(m.inspectCount)
				fmt.Printf("\n")
			}
			fmt.Printf("\n")
		}
		if round == 1 || round == 20 {
			fmt.Printf("After round %d:\n", round)
			for _, m := range monkeys {
				fmt.Print(m.inspectCount)
				fmt.Printf("\n")
			}
			fmt.Printf("\n")
		}
	}

	max1 := 0
	max2 := 0
	for _, monkey := range monkeys {
		fmt.Printf("Monkey %d inspected %d items\n", monkey.id, monkey.inspectCount)
		if monkey.inspectCount >= max1 {
			max2 = max1
			max1 = monkey.inspectCount
		} else if monkey.inspectCount > max2 {
			max2 = monkey.inspectCount
		}
	}

	fmt.Printf("%d * %d = %d\n", max1, max2, max1*max2)

	return max1 * max2
}

type Monkey struct {
	id           int
	inventory    []int
	op           string
	opVal        int
	testVal      int
	trueDest     int
	falseDest    int
	inspectCount int
}

func parseMonkey(lines []string) *Monkey {
	monkey := new(Monkey)
	var err error
	// ID
	id, err := strconv.Atoi(strings.TrimRight(strings.Fields(lines[0])[1], ":"))
	monkey.id = id

	//Inventory
	temp := strings.Split(strings.Split(lines[1], ":")[1], ",")
	var inventory []int
	for _, str := range temp {
		s := strings.TrimSpace(str)
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		inventory = append(inventory, i)
	}
	monkey.inventory = inventory

	//Op and OpVal
	temp = strings.Fields(lines[2])
	if temp[5] == "old" {
		monkey.op = "square"
		monkey.opVal = -1
	} else {
		monkey.op = temp[4]
		i, _ := strconv.Atoi(temp[5])
		monkey.opVal = i
	}

	//TestVal
	temp = strings.Fields(lines[3])
	i, err := strconv.Atoi(temp[3])
	monkey.testVal = i

	//Dests
	temp = strings.Fields(lines[4])
	i, err = strconv.Atoi(temp[5])
	monkey.trueDest = i

	temp = strings.Fields(lines[5])
	i, err = strconv.Atoi(temp[5])
	monkey.falseDest = i

	if err != nil {
		log.Fatal(err)
	}

	//Inspect Count
	monkey.inspectCount = 0

	return monkey
}

func calcWorry(input int, op string, opVal int) int {
	if op == "square" {
		return input * input
	}
	if op == "+" {
		return input + opVal
	}
	if op == "*" {
		return input * opVal
	}
	return input
}
