package n_queens

import "testing"

type tt struct {
	n      int
	result []*pos //queen positions
}

var table = []*tt{
	{0, []*pos{}},
	{1, []*pos{{0, 0}}},
	{2, []*pos{}},
	{3, []*pos{}},
	{4, []*pos{{1, 0}, {3, 1},
		{0, 2}, {2, 3}}},
}

func TestProcedural(t *testing.T) {
	for _, tcase := range table {
		res := procedural(tcase.n)
		if comp(tcase.result, res) == false {
			t.Errorf("failed for n:'%v', exp '%v', got '%v'\n",
				tcase.n, tcase.result, res)
		}
	}
}

func comp(A, B []*pos) bool {
	if len(A) != len(B) {
		return false
	}
	for i, a := range A {
		b := B[i]
		if a.row != b.row ||
			a.col != b.col {
			return false
		}
	}

	return true
}
