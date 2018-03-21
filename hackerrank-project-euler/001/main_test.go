package main

import "testing"

var table = map[int64]int64{
	10:  23,
	100: 2318,
}

func TestGetSumsMultipliers(t *testing.T) {
	for in, out := range table {
		r := getResult(in)
		if out != r {
			t.Errorf("for %v exp %v got %v", in, out, r)
		}
	}
}
