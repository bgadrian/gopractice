package _6_26_calculator_arithmetic_no_parant

import "testing"

var table = map[string]float64{
	"1+1":   2,
	"1-1":   0,
	"1*20":  20,
	"30/3":  10,
	"5+2*3": 11,
	"5*2+3": 13,

	"5*2+3*3":    19,
	"5*2*3/3":    10,
	"2+3+5-2*10": -10,
}

func TestSolutionBasic(t *testing.T) {
	for input, output := range table {
		r := solution(input)
		if r != output {
			t.Errorf("failed for '%v', exp '%v' got '%v'",
				input, output, r)
		}
	}
}
