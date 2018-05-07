package subarray_sum_0

import "testing"

var table = [][][]int{
	{
		{1, 3, -3, 4, 0}, //input
		{1, 2},           //output
	}, {
		{},     //input
		{0, 0}, //output
	}, {
		{0},    //input
		{0, 1}, //output
	}, {
		{3, 8, 10}, //input
		{0, 0},     //output
	}, { //entire array
		{1, 5, 4, -10}, //input
		{0, 4},         //output
	},
}

func TestSolution(t *testing.T) {
	for _, tt := range table {
		indexStart, size := solution(tt[0])
		if indexStart != tt[1][0] || size != tt[1][1] {
			t.Errorf("for '%v' failed, exp '%v', got '%v,'%v'",
				tt[0], tt[1], indexStart, size)
			continue
		}
	}
}
