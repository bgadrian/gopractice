package main

import "testing"

var table = map[int]string{
	0:             "Zero",
	10:            "Ten",
	17:            "Seventeen",
	60:            "Sixty",
	101:           "One Hundred One",
	1003:          "One Thousand Three",
	1010:          "One Thousand Ten",
	1400:          "One Thousand Four Hundred",
	1770:          "One Thousand Seven Hundred Seventy",
	3080900007:    "Three Billion Eighty Million Nine Hundred Thousand Seven",
	30000010:      "Thirty Million Ten",
	7000000:       "Seven Million",
	1010000000:    "One Billion Ten Million",
	1000010000:    "One Billion Ten Thousand",
	9001001001001: "Nine Trillion One Billion One Million One Thousand One",
}

func TestResult(t *testing.T) {
	for in, out := range table {
		r := getResult(in)
		if out != r {
			t.Errorf("for %v exp '%v' got '%v'", in, out, r)
		}
	}
}
