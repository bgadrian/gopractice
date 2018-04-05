package main

import (
	"testing"
)

var names = []string{
	"ALEX",
	"LUIS",
	"JAMES",
	"PAMELA", //240
	"BRIAN",
}
var table = map[string]int{
	"PAMELA": 240,
}

func TestResult(t *testing.T) {
	scores := getScoreMap(names)

	for in, out := range table {
		if scores[in] != out {
			t.Errorf("for %v exp '%v' got '%v'", in, out, scores[in])
		}
	}
}
