package __2is_string_permutation

import "testing"

var good = [][]string{
	{"abba", "bbaa"},
	{"a  ", " a "},
}
var bad = [][]string{
	{"abbaa", "bbaa"},
	{"abba", "bbaax"},
	{"a ", " a "},
}

func TestSolution(t *testing.T) {
	for _, tc := range good {
		r := solution([]rune(tc[0]), []rune(tc[1]))
		if r == false {
			t.Errorf("failed for '%v','%v'", tc[0], tc[1])
		}
	}
	for _, tc := range bad {
		r := solution([]rune(tc[0]), []rune(tc[1]))
		if r {
			t.Errorf("failed for '%v','%v'", tc[0], tc[1])
		}
	}
}
