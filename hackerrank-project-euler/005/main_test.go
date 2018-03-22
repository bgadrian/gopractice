package main

import "testing"

var table = map[int]int{
	3:  6, //6/1=6,6/2=3,6/3=2
	10: 2520,
}

func TestResult(t *testing.T) {
	for in, out := range table {
		r := getResult(in)
		if out != r {
			t.Errorf("for %v exp %v got %v", in, out, r)
		}
	}
}
