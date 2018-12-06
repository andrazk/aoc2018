package main

import "testing"

var in = []string{
	"#1 @ 1,3: 4x4",
	"#2 @ 3,1: 4x4",
	"#3 @ 5,5: 2x2",
	"#4 @ 3,3: 2x2",
}

func TestDay03(t *testing.T) {
	if want, got := 4, day03(in); want != got {
		t.Errorf("Wanted %d, got %d", want, got)
	}
}

func TestDay0302(t *testing.T) {
	if want, got := "3", day0302(in); want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}
