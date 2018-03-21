package main

import "testing"

var table = map[int]int{
	10:  10,
	100: 44,
	155: 188,
}

func TestGetSumsMultipliers(t *testing.T) {
	for in, out := range table {
		r := getResult(in)
		if out != r {
			t.Errorf("for %v exp %v got %v", in, out, r)
		}
	}
}
