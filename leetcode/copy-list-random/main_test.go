package copy_list_random

import "testing"

type shortList struct {
	randoms []int
}

var table = []*shortList{
	//A,B,C,D randoms : A->C, B->A, D->B, 0 based
	{[]int{2, 0, -1, 1}},
	{[]int{1, 2, 3, 0}},
}

func fromTTToLL(test *shortList) *Node {
	list := make([]*Node, len(test.randoms))
	for i := range list {
		list[i] = &Node{}
		if i == 0 {
			continue
		}
		list[i-1].next = list[i]
	}

	for from, to := range test.randoms {
		if to == -1 {
			continue
		}
		list[from].random = list[to]
	}
	return list[0] //head
}

func fromLLToTT(head *Node) *shortList {
	result := &shortList{}
	cache := make(map[*Node]int) //place in list

	i := 0
	for curr := head; curr != nil; i++ {
		cache[curr] = i
		result.randoms = append(result.randoms, -1)
		curr = curr.next
	}

	i = 0
	for curr := head; curr != nil; i++ {
		if curr.random != nil {
			randI := cache[curr.random]
			result.randoms[i] = randI
		}
		curr = curr.next
	}
	return result
}

func equal(a, b *shortList) bool {
	if len(a.randoms) != len(b.randoms) {
		return false
	}

	for i, va := range a.randoms {
		vb := b.randoms[i]
		if va == vb {
			continue
		}
		return false
	}
	return true
}

func TestTest(t *testing.T) {
	for _, vt := range []*shortList{
		{[]int{0, 2, -1}},
		{[]int{1, 1, 1}},
	} {
		vx := fromLLToTT(fromTTToLL(vt))

		if equal(vt, vx) == false {
			t.Errorf("failed for '%v', got '%v'", vt, vx)
		}
	}
}

func TestSolution(t *testing.T) {
	for _, test := range table {
		inputAsLL := fromTTToLL(test)
		got := solution(inputAsLL)
		gotAsTT := fromLLToTT(got)

		if equal(test, gotAsTT) {
			continue
		}
		t.Errorf("failed '%v' got '%v'",
			test, gotAsTT)
	}
}
