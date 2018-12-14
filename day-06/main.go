package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/andrazk/aoc2018/helpers"
)

var (
	top, bottom, left, right int
	centers                  = map[int]*point{}
)

type point struct {
	x        int
	y        int
	sum      int
	infinity bool
}

func main() {
	in, err := helpers.ReadLines("data")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Solution for Day 6: %d\n", day0601(in))
	fmt.Printf("Solution for Day 6 Part 2: %d\n", day0602(in, 10000))
}

func setup(in []string) {
	for k, v := range in {
		x, y := splitToInt(v)
		centers[k] = &point{x, y, 0, false}

		if x > right {
			right = x
		}

		if y > bottom {
			bottom = y
		}
	}
}

func day0601(in []string) int {
	setup(in)
	// Iterate over every field inside limits
	// find closest center point
	iterate(func(x, y int, inf bool) {
		if p := closest(centers, x, y); p != nil {
			p.sum++
			// Check if closest point goes to infinity
			if inf {
				p.infinity = true
			}
		}
	})

	// Count for the win
	max := 0
	for _, p := range centers {
		if p.infinity {
			continue
		}

		if max < p.sum {
			max = p.sum
		}
	}

	return max
}

func day0602(in []string, distLimit int) int {
	setup(in)
	fields := 0
	// Iterate over every field inside limits
	// find fields with total distance less then given
	iterate(func(x, y int, inf bool) {
		if totalDist(x, y) < distLimit {
			fields++
		}
	})

	return fields
}

func totalDist(x, y int) int {
	sum := 0
	for _, c := range centers {
		sum += dist(x, c.x, y, c.y)
	}

	return sum
}

func closest(centers map[int]*point, x, y int) *point {
	var c *point
	min := -1
	for _, p := range centers {
		// Check if x,y is one of center points
		if x == p.x && y == p.y {
			return p
		}

		d := dist(x, p.x, y, p.y)
		// new closest point
		if min == -1 || min > d {
			c = p
			min = d
			continue
		}
		// another closest point
		if min == d {
			c = nil
		}
	}

	return c
}

func dist(x1, x2, y1, y2 int) int {
	return abs(x1-x2) + abs(y1-y2)
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return int(math.Abs(float64(n)))
}

func splitToInt(s string) (x, y int) {
	splits := strings.Split(s, ", ")
	x, _ = strconv.Atoi(splits[0])
	y, _ = strconv.Atoi(splits[1])
	return x, y
}

// infinity checks if point touches the limit border
func infinity(x, y int) bool {
	return x == left ||
		x == right ||
		y == top ||
		y == bottom
}

// iterate function over every field inside limits
// third parameter defines if
func iterate(f func(int, int, bool)) {
	for x := left; x <= right; x++ {
		for y := top; y <= bottom; y++ {
			f(x, y, infinity(x, y))
		}
	}
}
