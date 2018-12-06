package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/andrazk/aoc2018/helpers"
)

const (
	wu = "wakes up"
	fa = "falls asleep"
)

func main() {
	in, err := helpers.ReadLines("data")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Solution for Day 4: %d\n", day0401(in))
	fmt.Printf("Solution for Day 4 Part 2: %v\n", day0402(in))
}

func day0401(in []string) int {
	var fam, id int

	idre := regexp.MustCompile(`.* #(\d+) .*$`)
	m := map[int]map[int]int{}

	sort.Strings(in)

	for _, l := range in {

		a := l[19:]
		minute, _ := strconv.Atoi(l[15:17])

		if a == wu {
			for i := fam; i < minute; i++ {
				if _, ok := m[id][i]; !ok {
					m[id][i] = 0
				}
				m[id][i]++
				m[id][99]++
			}
			continue
		}
		if a == fa {
			fam = minute
			continue
		}

		mid := idre.FindStringSubmatch(a)
		if len(mid) != 2 {
			log.Fatalf("Error parsing log action %v\n", a)
		}
		id, _ = strconv.Atoi(mid[1])
		if _, ok := m[id]; !ok {
			m[id] = map[int]int{
				99: 0,
			}
		}
	}

	var maxid int
	for id, v := range m {
		if _, ok := m[maxid]; !ok {
			maxid = id
			continue
		}

		if v[99] > m[maxid][99] {
			maxid = id
		}
	}

	var maxm int
	for k, v := range m[maxid] {
		if k == 99 {
			continue
		}

		if m[maxid][maxm] < v {
			maxm = k
		}
	}

	return maxid * maxm
}

func day0402(in []string) int {
	var fam, id int

	idre := regexp.MustCompile(`.* #(\d+) .*$`)
	m := map[string]int{}

	sort.Strings(in)

	for _, l := range in {

		a := l[19:]
		minute, _ := strconv.Atoi(l[15:17])

		if a == wu {
			for i := fam; i < minute; i++ {
				key := fmt.Sprintf("%d-%d", id, i)
				if _, ok := m[key]; !ok {
					m[key] = 0
				}
				m[key]++
			}
			continue
		}

		if a == fa {
			fam = minute
			continue
		}

		mid := idre.FindStringSubmatch(a)
		if len(mid) != 2 {
			log.Fatalf("Error parsing log action %v\n", a)
		}
		id, _ = strconv.Atoi(mid[1])
	}

	var maxkey string
	for key, mm := range m {
		if m[maxkey] < mm {
			maxkey = key
		}
	}
	s := strings.Split(maxkey, "-")
	s0, _ := strconv.Atoi(s[0])
	s1, _ := strconv.Atoi(s[1])
	return s0 * s1
}
