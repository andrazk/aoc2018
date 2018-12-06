package main

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/andrazk/aoc2018/helpers"
)

func main() {
	in, err := helpers.ReadLines("data")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Solution for Day 5: %d\n", day0501(in[0]))
	fmt.Printf("Solution for Day 5 Part 2: %d\n", day0502(in[0]))
}

func day0501(data string) int {
	return react([]byte(data))
}

func day0502(in string) int {
	min := len(in)

	for a := 65; a <= 90; a++ {
		su := string(a)
		sl := string(a + 32)
		left := strings.Replace(in, su, "", -1)
		left = strings.Replace(left, sl, "", -1)
		if m := react([]byte(left)); m < min {
			min = m
		}
	}

	return min
}

func react(in []byte) int {
	// Find reaction and explode
	i := 0
	for i < len(in)-1 {
		diff := math.Abs(float64(in[i]) - float64(in[i+1]))
		// 32 is difference between up and low case ascii chars.
		if diff == 32 {
			in = append(in[:i], in[i+2:]...)
			// start over again
			i = 0
			continue
		}

		i++
	}

	return len(in)
}
