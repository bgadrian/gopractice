package __3URLify_replace_spaces

import "testing"

type urilifyTest struct {
	in, out string
	tl      int
}

var table = []*urilifyTest{
	{"Mr John Smith    ", "Mr%20John%20Smith", 13},
	{"abcd", "abcd", 4},
	{"a b c d      ", "a%20b%20c%20d", 7},
}

func TestSolution(t *testing.T) {
	for _, c := range table {
		r := solution([]rune(c.in), c.tl)
		if r != c.out {
			t.Errorf("exp '%v' got '%v'", c.out, r)
		}
	}
}
