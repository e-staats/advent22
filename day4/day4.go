package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open("./day4.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		field1 := strings.Split(fields[0], "-")
		field2 := strings.Split(fields[1], "-")

		var vals []int
		for _, char := range field1 {
			int_val, err := strconv.Atoi(char)
			if err != nil {
				log.Fatal(err)
			}
			vals = append(vals, int_val)
		}
		for _, char := range field2 {
			int_val, err := strconv.Atoi(char)
			if err != nil {
				log.Fatal(err)
			}
			vals = append(vals, int_val)
		}
		if vals[0] == vals[2] || vals[1] == vals[3] {
			count++
			continue
		}
		if vals[0] < vals[2] {
			if vals[1] >= vals[3] {
				count++
				continue
			}
		} else if vals[3] >= vals[1] {
			count++
			continue
		}
		// fmt.Printf("Neither contains the other: %d-%d, %d-%d\n", vals[0], vals[1], vals[2], vals[3])
	}

	fmt.Printf("Part 1: %d\n", count)
}

func part2() {
	f, err := os.Open("./day4.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	count := 0
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		field1 := strings.Split(fields[0], "-")
		field2 := strings.Split(fields[1], "-")

		var vals []int
		for _, char := range field1 {
			int_val, err := strconv.Atoi(char)
			if err != nil {
				log.Fatal(err)
			}
			vals = append(vals, int_val)
		}
		for _, char := range field2 {
			int_val, err := strconv.Atoi(char)
			if err != nil {
				log.Fatal(err)
			}
			vals = append(vals, int_val)
		}
		if vals[2] > vals[1] || vals[0] > vals[3] {
			continue
		}
		// fmt.Printf("Overlaps: %d-%d, %d-%d\n", vals[0], vals[1], vals[2], vals[3])
		count++
	}

	fmt.Printf("Part 2: %d\n", count)
}
