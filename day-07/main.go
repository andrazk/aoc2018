package main

import (
	"bytes"
	"fmt"
	"log"
	"sort"

	"github.com/andrazk/aoc2018/helpers"
)

func main() {
	in, err := helpers.ReadLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Solution for Day 7: %v\n", day0701(in))
	fmt.Printf("Solution for Day 7 Part 2: %d\n", day0702(in))
}

type node struct {
	in, out []string
	visited bool
	time    int
}

type nmap map[string]*node

func day0701(in []string) string {
	nodes := parseNmap(in)

	path, ts := work(nodes, 1)
	fmt.Println("Time used", ts)

	b := bytes.Buffer{}
	for _, p := range path {
		b.WriteString(p)
	}

	return b.String()
}

func day0702(in []string) int {
	nodes := parseNmap(in)
	_, ts := work(nodes, 5)
	return ts
}

func sorts(in []string) {
	sort.Strings(in)
}

func parseNmap(in []string) nmap {

	newNode := func(k string) *node {
		return &node{
			in:      []string{},
			out:     []string{},
			visited: false,
			time:    int(k[0]) - 64 + 60,
		}
	}

	nodes := nmap{}

	for _, l := range in {
		f, b := l[5:6], l[36:37]

		if _, ok := nodes[f]; !ok {
			nodes[f] = newNode(f)
		}
		nodes[f].out = append(nodes[f].out, b)

		if _, ok := nodes[b]; !ok {
			nodes[b] = newNode(b)
		}
		nodes[b].in = append(nodes[b].in, f)
	}

	return nodes
}

func work(nodes nmap, workersLen int) ([]string, int) {
	clock := -1
	path := []string{}

	workers := make([]string, workersLen)

	// Go over nodes
	for {
		clock++

		// Check if all nodes were visited
		all := 0
		for _, n := range nodes {
			if n.visited {
				all++
			}
		}

		if all == len(nodes) {
			return path, clock
		}

		// fmt.Printf("%2d ", clock)

		// Check what can be added to next path
		next := []string{}
	NextLoop:
		for k, n := range nodes {
			if n.visited {
				continue
			}

			// working on
			for _, kk := range workers {
				if kk == k {
					continue NextLoop
				}
			}

			// All deps are done
			req := 0
			for _, prev := range n.in {
				if nodes[prev].visited {
					req++
				}
			}

			if req == len(n.in) {
				next = append(next, k)
			}
		}

		// Order by alphabet
		sorts(next)

		for k, w := range workers {
			// Is worker empty?
			if w == "" {
				if len(next) == 0 {
					// fmt.Print(".")
					continue
				}

				w = next[0]
				workers[k] = w
				next = next[1:]
			}

			// fmt.Print(w)

			nodes[w].time--

			if nodes[w].time > 0 {
				continue
			}

			path = append(path, w)
			nodes[w].visited = true
			workers[k] = ""
		}
		// fmt.Print("\n")
	}
}
