package main

import (
	"bufio"
	"fmt"
	"golang.org/x/exp/maps"
	"log"
	"os"
)

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open("./day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		halves := halve_string(scanner.Text())
		half1 := count_freq(halves[0])
		for _, r := range halves[1] {
			_, isPresent := half1[r]
			if isPresent == true {
				sum += score_letter(r)
				break
			}
		}
	}
	fmt.Printf("%d\n", sum)
}

func part2() {
	f, err := os.Open("./day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		elf1 := count_freq([]rune(scanner.Text()))
		scanner.Scan()
		elf2 := count_freq([]rune(scanner.Text()))
		scanner.Scan()
		elf3 := count_freq([]rune(scanner.Text()))
		elf1keys := maps.Keys(elf1)
		for _, r := range elf1keys {
			_, isPresent2 := elf2[r]
			_, isPresent3 := elf3[r]
			if isPresent2 && isPresent3 {
				sum += score_letter(r)
				fmt.Printf("%q %d\n", r, score_letter(r))
				break
			}
		}

	}
	fmt.Printf("%d\n", sum)
}

func score_letter(letter rune) int {
	ascii := int(letter)
	if ascii < 97 { //uppercase
		return ascii - 38
	}
	return ascii - 96
}

func halve_string(str string) [][]rune {
	str_arr := []rune(str)
	length := len(str_arr)
	return [][]rune{str_arr[:length/2], str_arr[length/2:]}
}

func count_freq(runes []rune) map[rune]int {
	freq := make(map[rune]int)
	for _, r := range runes {
		freq[r] += 1
	}
	return freq
}
