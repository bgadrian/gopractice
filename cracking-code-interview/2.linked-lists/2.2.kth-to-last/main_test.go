package __2_kth_to_last

import (
	"testing"
)

type tll struct {
	elements []int
	k, out   int
}

var table = []tll{
	{elements: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, k: 5, out: 11},
	{elements: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, k: 1, out: 15},
	{elements: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, k: 10, out: 6},
	{elements: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, k: 15, out: 1},
}

func TestSolution(t *testing.T) {
	for _, tt := range table {
		r := solutionRec(tt.elements, tt.k)
		if r != tt.out {
			t.Errorf("exp '%v', got '%v', for '%v' k=%v", tt.out, r, tt.elements, tt.k)
		}
	}
}

func TestSolutionIterative(t *testing.T) {
	for _, tt := range table {
		r := solutionIt(tt.elements, tt.k)
		if r != tt.out {
			t.Errorf("exp '%v', got '%v', for '%v' k=%v", tt.out, r, tt.elements, tt.k)
		}
	}
}
