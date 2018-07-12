package sum_4

import (
	"sort"
	"testing"
)

type st struct {
	S      []int
	target int
	result [][]int
}

var table = []*st{
	{
		S:      []int{1, 0, -1, 0, -2, 2},
		target: 0,
		result: [][]int{
			{-2, -1, 1, 2},
			{-2, 0, 0, 2},
			{-1, 0, 0, 1},
		},
	},
}

func TestSolution(t *testing.T) {
	for _, test := range table {
		res := solution(test.S, test.target)

		sort.Sort(matrix(test.result))
		sort.Sort(matrix(res))
		if equal(res, test.result) {
			continue
		}

		t.Errorf("failed for '%v', exp '%v', got '%v'",
			test.S, test.result, res)
	}
}

type matrix [][]int

func (s matrix) Len() int      { return len(s) }
func (s matrix) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s matrix) Less(i, j int) bool {
	for indexI := range s[i] {
		if s[i][indexI] == s[j][indexI] {
			continue
		}
		return s[i][indexI] < s[j][indexI]
	}
	return false
}

func equal(a, b [][]int) bool {
	la, lb := len(a), len(b)
	if la != lb {
		return false
	}

	for i, va := range a {
		vb := b[i]

		if len(va) != len(vb) {
			return false
		}

		for in := range va {
			if va[in] == vb[in] {
				continue
			}
			return false
		}
	}

	return true
}
