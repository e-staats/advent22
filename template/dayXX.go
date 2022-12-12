package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.Open("./dayXX.txt")
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
	return 0
}

func part2(lines []string) int {
	return 0
}
