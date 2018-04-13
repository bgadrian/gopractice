package main

import (
	"math/big"
	"testing"
)

var table = map[int]*big.Int{
	5: big.NewInt(12),
}

func TestResult(t *testing.T) {
	for in, out := range table {
		r := getResult(in)
		if out.Cmp(r) != 0 {
			t.Errorf("for %v exp %v got %v", in, out, r)
		}
	}
}
