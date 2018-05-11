package __12_count_paths_with_sum

import (
	"math"
	"testing"
)

const null = math.MinInt32

func TestSolution(t *testing.T) {
	tree := []int{
		10,
		5, -3,
		3, 2, null, 11,
		3, -2, null, 1, null, null,
	}
	sum := 8
	res := 3
	r := solution(makeTree(tree), sum)

	if res != r {
		t.Errorf("for '%v' sum:'%v', exp '%v', got '%v'",
			tree, sum, res, r)
	}
}

//needs a complete binary tree
//formula is parent at (floor(i-1)/2) with kids at (i*2+1 and i*2+2 for right)
func makeTree(a []int) *node {
	list := make([]*node, len(a))

	for i := 0; i < len(a); i++ {
		if a[i] == null {
			continue
		}
		me := &node{val: a[i]}
		list[i] = me
		if i == 0 {
			continue
		}
		iAmLeft := i%2 == 1
		parent := list[(i-1)/2]
		if iAmLeft {
			parent.left = me
		} else {
			parent.right = me
		}
	}
	return list[0]
}
