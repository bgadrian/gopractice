package dungeon_princess

import "testing"

type test struct {
	should int
	grid   [][]int
}

func TestSolution(t *testing.T) {
	table := []test{
		{
			7, [][]int{
				{-2, -3, 3},
				{-5, -10, 1},
				{10, 30, -5},
			},
		},
	}

	for _, test := range table {
		got := solution(test.grid)
		if got == test.should {
			continue
		}
		t.Errorf("got %d should %d for %v",
			got, test.should, test.grid)
	}

}
