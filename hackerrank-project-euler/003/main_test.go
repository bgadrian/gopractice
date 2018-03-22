package main

import "testing"

var table = map[int]int{
	13195: 29,
	10:    5,
	17:    17,
}

func TestGetSumsMultipliers(t *testing.T) {
	for in, out := range table {
		r := getResult(in)
		if out != r {
			t.Errorf("for %v exp %v got %v", in, out, r)
		}
	}
}
