package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	/*
		  A  B  C
		X 3  0  6
		Y 6  3  0
		Z 0  6  0

	*/
	scores := [][]int{
		[]int{3, 0, 6},
		[]int{6, 3, 0},
		[]int{0, 6, 3},
	}
	us := map[string]int{"X": 0, "Y": 1, "Z": 2}
	them := map[string]int{"A": 0, "B": 1, "C": 2}
	values := map[string]int{"X": 1, "Y": 2, "Z": 3}
	/*
		  L  T  W
		A Y  X  Z
		B Z  Y  X
		C X  Z  Y

	*/
	strats := [][]string{
		[]string{"Z", "X", "Y"},
		[]string{"X", "Y", "Z"},
		[]string{"Y", "Z", "X"},
	}

	part1(scores, us, them, values)
	part2(scores, us, them, values, strats)

}

func part1(scores [][]int, us map[string]int, them map[string]int, values map[string]int) {
	f, err := os.Open("./day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		plays := strings.Fields(scanner.Text())
		score := scores[us[plays[1]]][them[plays[0]]] + values[plays[1]]
		total += score
	}
	fmt.Printf("%d\n", total)
}
func part2(scores [][]int, us map[string]int, them map[string]int, values map[string]int, strats [][]string) {
	f, err := os.Open("./day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		plays := strings.Fields(scanner.Text())
		outcome_idx := us[plays[1]]
		our_play := strats[them[plays[0]]][outcome_idx]
		score := scores[us[our_play]][them[plays[0]]] + values[our_play]
		total += score
	}
	fmt.Printf("%d\n", total)

}
