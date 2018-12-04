package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	in := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		in = append(in, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Solution for Day 1: %d\n", day01(in))
	fmt.Printf("Solution for Day 1 Part 2: %d\n", day01Part02(in))
}

func day01Part02(in []string) int {
	m := make(map[int]int, 0)
	sum := 0
	for true {
		for _, s := range in {
			i, _ := strconv.Atoi(s)

			if _, ok := m[sum]; !ok {
				m[sum] = 0
			}

			m[sum]++
			if m[sum] == 2 {
				return sum
			}

			sum += i
		}
	}

	return 0
}

func day01(in []string) int {
	var sum int
	for _, s := range in {
		i, _ := strconv.Atoi(s)
		sum += i
	}

	return sum
}
