package main

import (
	"fmt"
	"testing"
)

func TestDay01Part02(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{
			in:   []string{"+1", "-1"},
			want: 0,
		},
		{
			in:   []string{"+3", "+3", "+4", "-2", "-4"},
			want: 10,
		},
		{
			in:   []string{"-6", "+3", "+8", "+5", "-6"},
			want: 5,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := day01Part02(c.in)
			if got != c.want {
				t.Errorf("Wanted %d, got %d", c.want, got)
			}
		})
	}
}

func TestDay01(t *testing.T) {
	// +1, +1, +1 results in  3
	// +1, +1, -2 results in  0
	// -1, -2, -3 results in -6
	cases := []struct {
		in   []string
		want int
	}{
		{
			in:   []string{"+1", "+1", "+1"},
			want: 3,
		},
		{
			in:   []string{"+1", "+1", "-2"},
			want: 0,
		},
		{
			in:   []string{"-1", "-2", "-3"},
			want: -6,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			got := day01(c.in)
			if got != c.want {
				t.Errorf("Wanted %d, got %d", c.want, got)
			}
		})
	}
}
