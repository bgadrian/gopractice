package _7_10_majority_element_constraints

import "testing"

var table = map[int][][]int{
	3:  {{3}, {3, 3}, {3, 3, 1}, {1, 3, 3}, {3, 1, 3}},
	2:  {{2, 1, 2, 1, 2}, {1, 2, 1, 2, 2}},
	-1: {{}, {1, 2, 1, 2}, {1, 2}, {1, 1, 1, 2, 2, 2}},
}

func TestSolution(t *testing.T) {
	for wanted, tt := range table {
		for _, test := range tt {
			res := solution(test)
			if res != wanted {
				t.Errorf("for '%v', got '%v', exp '%v'",
					test, res, wanted)
			}
		}
	}
}
