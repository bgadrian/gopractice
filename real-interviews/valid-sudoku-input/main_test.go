package valid_sudoku_input

import "testing"

func TestSolution(t *testing.T) {
	index := []string{"53..7....", "6..195...", ".98....6.", "8...6...3", "4..8.3..1", "7...2...6", ".6....28.", "...419..5", "....8..79"}
	exp := 1
	res := solution(index)

	if exp != res {
		t.Errorf("for '%v', exp '%v', got '%v'",
			index, exp, res)
	}
}
