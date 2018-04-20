package __5one_away_edit_strings

import "testing"

var good = [][]string{
	{"pale", "ple"},
	{"pales", "pale"},
	{"pale", "bale"},
	{"", "a"},
}

var bad = [][]string{
	{"pale", "bae"},
	{"pale", "paleXX"},
	{"", "pa"},
	{"a", "cc"},
}

func TestSolution(t *testing.T) {
	for _, tc := range good {
		r := solution([]rune(tc[0]), []rune(tc[1]))
		if r == false {
			t.Errorf("failed for '%v', '%v'", tc[0], tc[1])
		}
	}
	for _, tc := range bad {
		r := solution([]rune(tc[0]), []rune(tc[1]))
		if r {
			t.Errorf("failed for '%v', '%v'", tc[0], tc[1])
		}
	}
}
