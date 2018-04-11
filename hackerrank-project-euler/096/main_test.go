package main

import "testing"

type sudokuRes struct {
	in, out []string
}

var table = []*sudokuRes{
	{
		in: []string{
			"123456780",
			"456780123",
			"780123456",
			"234567801",
			"567801234",
			"801234567",
			"345678012",
			"678012345",
			"012345678"},
		out: []string{
			"123456789",
			"456789123",
			"789123456",
			"234567891",
			"567891234",
			"891234567",
			"345678912",
			"678912345",
			"912345678"},
	},
}

func compStrs(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, va := range a {
		if va != b[i] {
			return false
		}
	}
	return true
}

func TestResult(t *testing.T) {
	for _, data := range table {
		r := getResult(data.in)
		if compStrs(r, data.out) == false {
			t.Errorf("for %v exp %v got %v", data.in, data.out, r)
		}
	}
}
