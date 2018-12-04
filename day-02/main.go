package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
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

	fmt.Printf("Solution for Day 2: %d\n", day02(in))
	fmt.Printf("Solution for Day 2 Part 2: %s\n", day02Part2(in))
}

func day02(in []string) int {
	var t2, t3 int
	for _, s := range in {
		m := map[rune]int{}

		for _, c := range s {
			if _, ok := m[c]; !ok {
				m[c] = 0
			}

			m[c]++
		}

		for _, t := range m {
			if t == 2 {
				t2++
				break
			}
		}

		for _, t := range m {
			if t == 3 {
				t3++
				break
			}
		}
	}

	return t2 * t3
}

func day02Part2(in []string) string {
	for _, s1 := range in {
		for _, s2 := range in {
			if s1 == s2 {
				continue
			}

			diff := []int{}
			for i := 0; i < len(s1); i++ {

				if s1[i] != s2[i] {
					diff = append(diff, i)
				}
			}

			if len(diff) == 1 {
				var b bytes.Buffer
				b.WriteString(s1[:diff[0]])
				b.WriteString(s1[diff[0]+1:])
				return b.String()
			}
		}
	}

	return ""
}
