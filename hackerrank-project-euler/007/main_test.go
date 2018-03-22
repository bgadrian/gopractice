package main

import "testing"

var table = map[int]int{
	3: 5,
	6: 13,
}

func TestResult(t *testing.T) {
	for in, out := range table {
		r := getResult(in)
		if out != r {
			t.Errorf("for %v exp %v got %v", in, out, r)
		}
	}
}
