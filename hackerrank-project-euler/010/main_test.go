package main

import "testing"

//2,3,5,7,11,13,17
var table = map[int]int{
	5:  10,
	10: 17,
	8:  17,
	11: 28,
}

func TestResult(t *testing.T) {
	for in, out := range table {
		r := getResult(in)
		if out != r {
			t.Errorf("for %v exp %v got %v", in, out, r)
		}
	}
}
