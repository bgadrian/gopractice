package longest_substr_distinct_chars

import "testing"

var tableTwo = map[string]string{
	"ababa":               "ababa",
	"aabbcab":             "aabb",
	"caaabbb":             "aaabbb",
	"aaxababa":            "ababa",
	"abcbbbbcccbdddadacb": "bcbbbbcccb",
}

func TestSolution2(t *testing.T) {
	for in, out := range tableTwo {
		res := solution(in, 2)
		if res == out {
			continue
		}

		t.Errorf("for '%v', exp '%v', got '%v'",
			in, out, res)
	}
}
