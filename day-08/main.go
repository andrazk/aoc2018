package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/andrazk/aoc2018/helpers"
)

func main() {
	in, err := helpers.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Solution for Day 8: %v\n", day0801(in[0]))
	fmt.Printf("Solution for Day 8 Part 2: %d\n", day0802(in[0]))
}

func day0801(in string) int {
	_, sum, _ := treverse(parseTree(in), 0)

	return sum
}

func day0802(in string) int {
	_, _, ms := treverse(parseTree(in), 0)
	return ms
}

func parseTree(in string) []int {
	tree := []int{}
	for _, v := range strings.Split(in, " ") {
		n, _ := strconv.Atoi(v)
		tree = append(tree, n)
	}
	return tree
}

func treverse(tree []int, sum int) ([]int, int, int) {
	childs := tree[0]
	metas := tree[1]
	tree = tree[2:]
	metaSum := 0
	childMetaSums := map[int]int{}

	for c := 0; c < childs; c++ {
		tree, sum, childMetaSums[c+1] = treverse(tree, sum)
	}

	tmpSum := 0
	for i := 0; i < metas; i++ {
		tmpSum += tree[i]
	}

	// meta sums
	if childs == 0 {
		metaSum = tmpSum
	} else {
		for _, m := range tree[:metas] {
			s, ok := childMetaSums[m]
			if !ok {
				continue
			}

			metaSum += s
		}
	}

	tree = tree[metas:]
	return tree, sum + tmpSum, metaSum
}
