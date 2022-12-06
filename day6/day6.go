package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	fmt.Printf("Part 1: %d\n", part1())
	duration := time.Since(start)
	fmt.Println(duration)

	start = time.Now()
	fmt.Printf("Part 2: %d\n", part2())
	duration = time.Since(start)
	fmt.Println(duration)
}

func part1() int {
	f, err := os.Open("./day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(f)
	index := 1
	var buffer []rune
	seen := make(map[rune]int)
	for {
		if c, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			dup := false
			if len(buffer) < 3 {
				buffer = append(buffer, c)
				seen[c]++
				index++
				continue
			}
			if len(buffer) == 3 {
				buffer = append(buffer, c)
				seen[c]++
			} else {
				seen[buffer[0]] -= 1
				seen[c]++
				buffer = append(buffer[1:], c)
			}
			for _, r := range buffer {
				if seen[r] > 1 {
					dup = true
				}
			}
			if dup == true {
				index++
			} else {
				return index
			}

		}

	}
	return index
}

func part2() int {
	f, err := os.Open("./day6.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(f)
	index := 1
	var buffer []rune
	seen := make(map[rune]int)
	for {
		if c, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			dup := false
			if len(buffer) < 13 {
				buffer = append(buffer, c)
				seen[c]++
				index++
				continue
			}
			if len(buffer) == 13 {
				buffer = append(buffer, c)
				seen[c]++
			} else {
				seen[buffer[0]] -= 1
				seen[c]++
				buffer = append(buffer[1:], c)
			}
			for _, r := range buffer {
				if seen[r] > 1 {
					dup = true
				}
			}
			if dup == true {
				index++
			} else {
				return index
			}

		}

	}
	return index
}
