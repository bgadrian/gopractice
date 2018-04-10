package main

import "testing"

type inputM struct {
	tips   [][]rune
	result []rune
}

var table = []*inputM{
	{
		tips:   [][]rune{[]rune("SMH"), []rune("TON"), []rune("RNG"), []rune("WRO"), []rune("THG")},
		result: []rune("SMTHWRONG"),
	}, { //fail test
		tips:   [][]rune{[]rune("an0"), []rune("n/."), []rune(".#a")},
		result: []rune("SMTH WRONG"),
	},
}

func TestResult(t *testing.T) {
	for _, in := range table {
		r := getResult(in.tips)
		if compareRunes(r, in.result) == false {
			t.Errorf("exp %s got %s", string(in.result), string(r))
		}
	}
}

func compareRunes(r1, r2 []rune) bool {
	if len(r1) != len(r2) {
		return false
	}

	for i, r := range r1 {
		if r != r2[i] {
			return false
		}
	}
	return true
}
