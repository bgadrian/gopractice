package main

import "testing"

type inputChars struct {
	encoded []rune
	key     []rune
	decoded []rune
}

var table = []*inputChars{
	{
		encoded: []rune{32, 66, 50, 20, 11, 0, 42, 66, 33, 19, 13, 20, 47, 66, 37, 14, 58, 67, 43, 23, 14, 17, 49, 67, 46, 20, 6, 51, 66, 55, 9, 39, 67, 45, 3, 25, 56, 66, 39, 14, 37, 34, 65, 51, 22, 8, 1, 40, 65, 32, 17, 14, 21, 45, 65, 36, 12, 57, 66, 41, 20, 15, 19, 50, 66, 44, 23, 7, 49, 65, 54, 11, 36, 66, 47, 0, 24, 58, 65, 38, 12, 38},
		key:     []rune{'a', 'b', 'c'},
	}}

func TestResult(t *testing.T) {
	for _, inputChar := range table {
		r := getResult(inputChar.encoded)
		if compareRunes(inputChar.key, r) == false {
			t.Errorf("exp %s got %s", string(inputChar.key), string(r))
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
