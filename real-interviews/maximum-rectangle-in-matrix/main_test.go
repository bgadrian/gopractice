package maximum_rectangle_in_matrix

import "testing"

var tableMatrix = map[int][][][]int{
	0: {
		{}, {{0, 0}, {0, 0}},
	},
	2: {
		{
			{0, 0, 1, 0, 1},
			{0, 0, 0, 1, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 1, 0, 0},
			{0, 0, 0, 1, 1},
		},
	},
}

func TestMatrix(t *testing.T) {
	for expected, list := range tableMatrix {
		for _, test := range list {
			result := maxOnesAreaMatrix(test)
			if result == expected {
				continue
			}
			t.Errorf("for '%v', exp '%v' got '%v'",
				test, expected, result)
		}
	}
}

var tableHisto = map[int][][]int{
	0:  {{}, {0, 0, 0}},
	1:  {{1}, {0, 1, 0, 1}},
	3:  {{1, 2, 1}},
	4:  {{1, 2, 4}},
	5:  {{1, 5, 1}, {2, 1, 2, 3, 1}},
	12: {{6, 2, 5, 4, 5, 1, 6}},
}

func TestHistoBasic(t *testing.T) {

	for expected, list := range tableHisto {
		for _, test := range list {
			result := maxHistogramArea(test)
			if result == expected {
				continue
			}
			t.Errorf("for '%v', exp '%v' got '%v'",
				test, expected, result)
		}
	}
}
