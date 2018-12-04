package main

import "testing"

func TestDay02(t *testing.T) {
	in := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}

	if got, want := day02(in), 12; got != want {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestDay02Part2Example(t *testing.T) {
	in := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}

	if got, want := day02Part2(in), "fgij"; got != want {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}
