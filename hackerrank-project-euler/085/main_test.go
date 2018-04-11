package main

import (
	"math"
	"testing"
)

var table = map[int]int{
	18:      6,
	2:       2,
	33:      9,
	2000000: 2772,
}

func TestResult(t *testing.T) {
	for in, out := range table {
		r := getResult(in)
		if out != r {
			t.Errorf("for %v exp %v got %v", in, out, r)
		}

		//r = bruteForceR(in)
		//if out != r {
		//	t.Errorf("BF for %v exp %v got %v", in, out, r)
		//}
	}
}

func bruteForceR(L int) (area int) {
	min_diff := L

	for x := 2; x <= L; x++ {
		for y := x; y <= L; y++ {
			diff := int(math.Abs(float64(x*(x+1)*y*(y+1)/4 - L)))
			if diff < min_diff {
				area, min_diff = x*y, diff
			}
		}
	}
	return
}
