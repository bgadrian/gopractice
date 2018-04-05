package main

import "testing"

var table = map[int]string{
	1: "abcdefghijklm",
	2: "abcdefghijkml",
}

func TestResult(t *testing.T) {
	for in, out := range table {
		r := getResult("abcdefghijklm", in)
		if out != r {
			t.Errorf("for %v exp '%v' got '%v'", in, out, r)
		}
	}
}
