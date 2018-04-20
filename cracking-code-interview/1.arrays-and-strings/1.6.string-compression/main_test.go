package __6_string_compression

import "testing"

var scs = [][]string{
	{"aabcccccaaa", "a2b1c5a3"},
	{"abcd", "abcd"},
	{"", ""},
	{"a", "a"},
	{"aaaaaaaaa", "a9"},
	{"aaab", "a3b1"},
}

func TestSolution(t *testing.T) {
	for _, sc := range scs {
		r := solution([]rune(sc[0]))
		if string(r) != sc[1] {
			t.Errorf("for test '%v', exp '%v', got '%v'", sc[0], sc[1], string(r))
		}
	}
}
