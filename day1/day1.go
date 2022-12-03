package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open("./day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	max := 0
	sum := 0

	for scanner.Scan() {
		val := scanner.Text()
		if val == "" {
			if sum > max {
				max = sum
			}
			sum = 0
		} else {
			intVar, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			sum += intVar
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1: %d\n", max)
}

func part2() {
	f, err := os.Open("./day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	top := []int{0, 0, 0}
	sum := 0

	for scanner.Scan() {
		val := scanner.Text()
		if val == "" {
			if sum > top[0] {
				top[0] = sum
				sort.Ints(top)
			}
			sum = 0
		} else {
			intVar, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			sum += intVar
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	tot_sum := top[0] + top[1] + top[2]

	fmt.Printf("Part 2:\n%d\n%d\n%d\n%d\n", top[0], top[1], top[2], tot_sum)

}
