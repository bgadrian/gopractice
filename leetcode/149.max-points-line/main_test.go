package _49_max_points_line

import "testing"

var table = map[int][]Point{
	//2: {{-1, -1}, {1, 1}},
	3: {{1, 1}, {2, 2}, {3, 3}},
	4: {{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}},
}

func TestSolution(t *testing.T) {
	for should, input := range table {
		got := solution(input)
		if got == should {
			continue
		}
		t.Errorf("for '%v', exp '%v' got '%v'",
			input, should, got)
	}
}
