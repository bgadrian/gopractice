package anagrams

import "testing"

func TestHash(t *testing.T) {
	var v = map[string]int{
		"a":    1,
		"b":    2,
		"c":    4,
		"d":    8,
		"abcd": 15,
	}

	for input, expected := range v {
		res := hashToNum(input)
		if res == int32(expected) {
			continue
		}
		t.Errorf("for '%v', got '%v', exp '%v'",
			input, res, expected)
	}
}
