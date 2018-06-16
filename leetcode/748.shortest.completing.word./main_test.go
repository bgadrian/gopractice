package _48_shortest_completing_word_

import (
	"fmt"
	"testing"
)

type st struct {
	license, result string
	words           []string
}

func (s *st) String() string {
	return fmt.Sprintf("{license: '%v', words: '%v'}",
		s.license, s.words)
}

var table = []*st{
	{
		"1s3 PSt", "steps",
		[]string{"step", "steps", "stripe", "stepple"},
	},
	{
		"1s3 456", "pest",
		[]string{"looks", "pest", "stew", "show"},
	},
}

func TestBasic(t *testing.T) {
	for _, test := range table {
		got := shortestCompletingWord(test.license, test.words)
		if got == test.result {
			continue
		}
		t.Errorf("for '%v' got '%v', exp '%v'",
			test, got, test.result)
	}
}
