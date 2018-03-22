package main

import "testing"

var table = map[int]int{
	101110: 101101, //143 x 707
	800000: 793397, //869 x 913
}

func TestResult(t *testing.T) {
	for in, out := range table {
		r := getResult(in)
		if out != r {
			t.Errorf("for %v exp %v got %v", in, out, r)
		}
	}
}
