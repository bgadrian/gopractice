package knightChess

import "testing"

var table = map[int][6]int{
	6: {8, 8, 1, 1, 8, 8},
}

func TestSolution(t *testing.T) {
	for should, input := range table {
		got := solution(input)
		if got == should {
			continue
		}
		t.Errorf("failed for '%v', exp '%v' got '%v'",
			input, should, got)
	}
}
