package main

import "testing"

type edgeInput struct {
	edges  []*edge
	result int
}

var table = []*edgeInput{
	{edges: []*edge{
		{1, 2, 16},
		{1, 3, 12},
		{1, 4, 21},
		{2, 4, 17},
		{2, 5, 20},
		{3, 4, 28},
		{3, 6, 31},
		{4, 5, 18},
		{4, 6, 19},
		{4, 7, 23},
		{5, 7, 11},
		{6, 7, 27},
		{7, 6, 87},
	}, result: 93},
	{edges: []*edge{
		{1, 2, 3},
		{1, 2, 0},
		{1, 3, 2},
	}, result: 2},
}

func TestResult(t *testing.T) {
	for _, in := range table {
		r := getResult(in.edges)
		if in.result != r {
			t.Errorf("for %v exp %v got %v", in, in.result, r)
		}
	}
}
