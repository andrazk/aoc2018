package main

import (
	"testing"
)

var in = []string{
	"1, 1",
	"1, 6",
	"8, 3",
	"3, 4",
	"5, 5",
	"8, 9",
}

func TestDay0601(t *testing.T) {
	want := 17
	got := day0601(in)
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestDay0602(t *testing.T) {
	want := 16
	got := day0602(in, 32)
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}
