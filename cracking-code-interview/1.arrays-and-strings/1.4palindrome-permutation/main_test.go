package __4palindrome_permutation

import "testing"

var good = []string{"Tact Coa", "attcaCo", "", "a", "aab", "aaabbbbaaa"}
var bad = []string{"act coa", "ab", "aaaabbbccccddd"}

func TestSolution(t *testing.T) {
	for _, tc := range good {
		r := solution([]rune(tc))
		if r == false {
			t.Errorf("failed for '%v'", tc)
		}
	}
	for _, tc := range bad {
		r := solution([]rune(tc))
		if r {
			t.Errorf("positive failed for '%v'", tc)
		}
	}
}
