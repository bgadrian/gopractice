package test

import "testing"

var good = []string{
	"",
	"abcd",
	"abcd012efg",
}

var bad = []string{
	"  ",
	"abcda",
	"abcd012efg0",
}

func TestSolution(t *testing.T) {
	for _, i := range good {
		r := solution(i)
		if r == false {
			t.Errorf("failed for %v", i)
		}
	}
	for _, i := range bad {
		r := solution(i)
		if r {
			t.Errorf("failed for %v", i)
		}
	}
}
