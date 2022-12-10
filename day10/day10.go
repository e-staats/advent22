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
	f, err := os.Open("./day10.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	start := time.Now()
	result := part1(lines)
	fmt.Printf("Part 1: %d\n", result)
	duration := time.Since(start)
	fmt.Println(duration)

	fmt.Printf("\n")

	start = time.Now()
	result = part2(lines)
	fmt.Printf("Part 2: %d\n", result)
	duration = time.Since(start)
	fmt.Println(duration)
}

func part1(lines []string) int {
	cycle := 0
	register := 1
	var adding bool
	var value int
	var err error
	idx := 0
	output := 0
	for idx < len(lines) {
		cycle++
		if (cycle-20)%40 == 0 {
			output += (cycle * register)
			fmt.Printf("Cycle %d: adding (%d * %d) - new total is %d\n", cycle, cycle, register, output)
		}
		if adding == true {
			register += value
			adding = false
			idx++
			continue
		}
		fields := strings.Fields(lines[idx])
		if fields[0] == "noop" {
			idx++
			continue
		}
		if fields[0] == "addx" {
			value, err = strconv.Atoi(fields[1])
			if err != nil {
				log.Fatal(err)
			}
			adding = true
		}
	}
	return output
}

func part2(lines []string) int {
	cycle := 0
	register := 1
	var adding bool
	var value int
	var err error
	idx := 0
	output := 0
	for idx < len(lines) {
		cycle++
		if comp := register + 2 - (cycle % 40); comp >= 0 && comp < 3 {
			// fmt.Printf("Cycle %d: register is %d\n", cycle, register)
			print("#")
		} else {
			// fmt.Printf("Cycle %d: register is %d\n", cycle, register)
			print(".")
		}
		if cycle%40 == 0 {
			output += (cycle * register)
			fmt.Printf("\n")
		}
		if adding == true {
			register += value
			adding = false
			idx++
			continue
		}
		fields := strings.Fields(lines[idx])
		if fields[0] == "noop" {
			idx++
			continue
		}
		if fields[0] == "addx" {
			value, err = strconv.Atoi(fields[1])
			if err != nil {
				log.Fatal(err)
			}
			adding = true
		}
	}
	return output
}

type Directory struct {
	name     string
	size     int
	parent   *Directory
	children []*Directory
}

func new_dir(name string) *Directory {
	result := new(Directory)
	result.name = name
	result.size = 0
	return result
}
