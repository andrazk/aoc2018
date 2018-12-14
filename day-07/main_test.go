package main

import (
	"bytes"
	"fmt"
	"testing"
)

var testdata = []string{
	"Step C must be finished before step A can begin.",
	"Step C must be finished before step F can begin.",
	"Step A must be finished before step B can begin.",
	"Step A must be finished before step D can begin.",
	"Step B must be finished before step E can begin.",
	"Step D must be finished before step E can begin.",
	"Step F must be finished before step E can begin.",
}

func Test_worker(t *testing.T) {
	cases := []struct {
		workers int
		time    int
		path    string
	}{
		{
			workers: 1,
			time:    21,
			path:    "CABDFE",
		},
		{
			workers: 2,
			time:    15,
			path:    "CABFDE",
		},
		{
			workers: 3,
			time:    14,
			path:    "CABDFE",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {

			gotPath, gotTime := work(parseNmap(testdata), c.workers)

			b := bytes.Buffer{}
			for _, p := range gotPath {
				b.WriteString(p)
			}

			if gotPath := b.String(); c.path != gotPath {
				t.Errorf("Wanted path %v, got %v", c.path, gotPath)
			}

			if c.time != gotTime {
				t.Errorf("Wanted time %v, got %v", c.time, gotTime)
			}
		})
	}
}
