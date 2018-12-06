package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/andrazk/aoc2018/helpers"
)

const (
	overlap = "x"
)

func main() {
	in, err := helpers.ReadLines("data")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Solution for Day 3: %d\n", day03(in))
	fmt.Printf("Solution for Day 3 Part 2: %v\n", day0302(in))
}

func day03(claims []string) int {
	field := map[int]map[int]string{}
	overlaps := 0

	re := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)

	for _, claim := range claims {
		// Parse claim string
		matches := re.FindStringSubmatch(claim)
		// id, _ := strconv.Atoi(matches[1])
		id := matches[1]
		l, _ := strconv.Atoi(matches[2])
		t, _ := strconv.Atoi(matches[3])
		w, _ := strconv.Atoi(matches[4])
		h, _ := strconv.Atoi(matches[5])

		l2 := l + w
		t2 := t + h
		for x := l; x < l2; x++ {
			if _, ok := field[x]; !ok {
				field[x] = map[int]string{}
			}
			for y := t; y < t2; y++ {
				if _, ok := field[x][y]; !ok {
					field[x][y] = id
				}

				// Dont count overlaps twice
				if field[x][y] != id && field[x][y] != overlap {
					field[x][y] = overlap
					overlaps++
				}

			}
		}
	}

	return overlaps
}

func day0302(claims []string) string {
	field := map[int]map[int][]string{}
	overlaps := map[string]int{}

	re := regexp.MustCompile(`^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$`)

	for _, claim := range claims {
		// Parse claim string
		matches := re.FindStringSubmatch(claim)
		// id, _ := strconv.Atoi(matches[1])
		id := matches[1]
		l, _ := strconv.Atoi(matches[2])
		t, _ := strconv.Atoi(matches[3])
		w, _ := strconv.Atoi(matches[4])
		h, _ := strconv.Atoi(matches[5])

		l2 := l + w
		t2 := t + h
		for x := l; x < l2; x++ {
			if _, ok := field[x]; !ok {
				field[x] = map[int][]string{}
			}
			for y := t; y < t2; y++ {
				if _, ok := field[x][y]; !ok {
					field[x][y] = []string{}
				}

				if _, ok := overlaps[id]; !ok {
					overlaps[id] = 0
				}

				// check for overlap
				if len(field[x][y]) > 0 {
					// mark self
					overlaps[id]++
					// mark the first id
					if len(field[x][y]) == 1 {
						overlaps[field[x][y][0]]++
					}
				}

				field[x][y] = append(field[x][y], id)
			}
		}
	}

	for k, v := range overlaps {
		if v == 0 {
			return k
		}
	}

	return overlap
}
