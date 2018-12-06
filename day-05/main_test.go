package main

import "testing"

var in = "dabAcCaCBAcCcaDA"

func TestDay0501(t *testing.T) {
	want := 10
	got := day0501(in)

	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}

func TestDay0502(t *testing.T) {
	want := 4
	got := day0502(in)

	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}
}
