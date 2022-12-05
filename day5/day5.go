package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gopkg.in/karalabe/cookiejar.v2/collections/stack"
)

func main() {
	part1()
	part2()

}
func part1() {
	f, err := os.Open("./day5.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	// Containers
	containersInput := stack.New() //Stack[string]
	for scanner.Scan() {
		if scanner.Text() == "" {
			containersInput.Pop()
			break
		}
		containersInput.Push(scanner.Text())
	}

	// Parse the containers into Stacks - returns stacks = []Stack[rune]
	stacks := parseContainers(containersInput)

	// Instructions
	for scanner.Scan() {
		instructions := parseInstruction(scanner.Text())
		for c := 0; c < instructions["number"]; c++ {
			val := stacks[instructions["source"]].Pop().(rune)
			stacks[instructions["target"]].Push(val)
		}

	}

	output("Part 1", stacks)

}

func part2() {
	f, err := os.Open("./day5.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	// Containers
	containersInput := stack.New() //Stack[string]
	for scanner.Scan() {
		if scanner.Text() == "" {
			containersInput.Pop()
			break
		}
		containersInput.Push(scanner.Text())
	}

	// Parse the containers into Stacks - returns stacks = []Stack[rune]
	stacks := parseContainers(containersInput)

	// Instructions
	for scanner.Scan() {
		instructions := parseInstruction(scanner.Text())
		clutch := stack.New()
		for c := 0; c < instructions["number"]; c++ {
			val := stacks[instructions["source"]].Pop().(rune)
			clutch.Push(val)
		}
		for clutch.Empty() == false {
			stacks[instructions["target"]].Push(clutch.Pop())
		}

	}

	output("Part 2", stacks)

}
func output(title string, stacks []*stack.Stack) {
	// Output
	var outString []rune

	for _, stack := range stacks {
		outString = append(outString, stack.Pop().(rune))
	}

	fmt.Printf("%s: %s\n", title, string(outString))
}

func parseContainers(input *stack.Stack) []*stack.Stack {
	// for each layer in the diagram
	// turn string into array of runes
	// while idx < len:
	// chunk by 4s. push char onto map[idx//4]
	stacks := []*stack.Stack{stack.New()}
	for input.Empty() != true {

		chars := []rune(input.Pop().(string))
		idx := 0
		for idx < len(chars) {
			if strings.TrimSpace(string(chars[idx:idx+3])) == "" {
				idx += 4
				continue
			}
			if idx/4 >= len(stacks) {
				stacks = append(stacks, stack.New())
			}
			stacks[idx/4].Push(chars[idx+1])
			idx += 4
		}
	}
	return stacks
}

func parseInstruction(instr string) map[string]int {
	results := make(map[string]int)
	words := strings.Fields(instr)
	source, err := strconv.Atoi(words[3])
	target, err := strconv.Atoi(words[5])
	number, err := strconv.Atoi(words[1])
	if err != nil {
		log.Fatal(err)
	}
	results["source"] = source - 1
	results["target"] = target - 1
	results["number"] = number
	return results
}
