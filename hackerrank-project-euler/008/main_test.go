package main

import "testing"

var table = []inputData{
	{5, "3675356291", 3150},
	{5, "2709360626", 0},
}

func TestResult(t *testing.T) {
	for _, in := range table {
		output := getResult(&in)
		if output != in.r {
			t.Errorf("for %v exp %v got %v", in.n, in.r, output)
		}
	}
}
