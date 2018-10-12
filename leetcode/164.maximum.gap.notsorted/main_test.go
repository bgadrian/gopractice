package _64_maximum_gap_notsorted

import "testing"

type tcase struct {
	should int
	in     []int
}

var table = []tcase{
	{3, []int{3, 6, 9, 1}},
	{0, []int{10}},
}

func TestBasic(t *testing.T) {
	for _, test := range table {
		got := solution(test.in)
		if got == test.should {
			continue
		}
		t.Errorf("got %d should %d, for %v",
			got, test.should, test.in)
	}
}
