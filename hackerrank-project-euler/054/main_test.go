package main

import (
	"testing"
)

var tableP2 = []string{
	"5H 5C 6S 7S KD 2C 3S 8S 8D TD",
	"2D 9C AS AH AC 3D 6D 7D TD QD",
}

var tableP1 = []string{
	"5D 8C 9S JS AC 2C 5C 7D 8S QH",
	"4D 6S 9H QH QC 3D 6D 7H QD QS",
	"2H 2D 4C 4D 4S 3C 3D 3S 9S 9D",

	//low straight
	"AD 2C 3S 4H 5C 3D 6C 7D TS QD",
}

func TestResult(t *testing.T) {
	for _, in := range tableP1 {
		p1, p2 := strToHands(in)
		output := getResult(p1, p2)
		if output != p1 {
			t.Errorf("failed for %v exp p1", in)
		}
	}

	for _, in := range tableP2 {
		p1, p2 := strToHands(in)
		output := getResult(p1, p2)
		if output != p2 {
			t.Errorf("failed for %v exp p2", in)
		}
	}
}
