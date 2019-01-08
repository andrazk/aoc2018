package main

import (
	"testing"
)

var testdata = "2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2"

func Test_Day0801(t *testing.T) {
	want := 138
	got := day0801(testdata)

	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func Test_Day0802(t *testing.T) {
	want := 66
	got := day0802(testdata)

	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}
