package clone_graph

import (
	"container/list"
	"fmt"
	"testing"
)

type GraphStr struct {
	nodes []rune
	edges [][2]rune
}

var table = []*GraphStr{
	{[]rune{'A', 'B', 'C', 'D', 'E'},
		[][2]rune{
			{'A', 'B'},
			{'A', 'D'},
			{'B', 'C'},
			{'C', 'E'},
		}},
	{[]rune{'A', 'B', 'C'},
		[][2]rune{
			{'A', 'C'},
			{'B', 'C'},
		}},
}

var genG = func(gs *GraphStr) *Node {
	cache := make(map[rune]*Node)
	var res *Node

	for _, l := range gs.nodes {
		cache[l] = &Node{label: string(l)}

		if res == nil {
			res = cache[l]
		}
	}

	for _, connection := range gs.edges {
		a := cache[connection[0]]
		b := cache[connection[1]]

		a.friends = append(a.friends, b)
		b.friends = append(b.friends, a)
	}

	return res
}

func TestSolution(t *testing.T) {
	for _, test := range table {
		g := genG(test)
		clone := Solution(g)
		ok, err := equal(g, clone)
		if ok {
			continue
		}
		t.Error(err)
	}
}

func equal(original, copy *Node) (bool, error) {
	visited := make(map[string]struct{})

	toVisitA := list.New()
	toVisitB := list.New()

	toVisitA.PushFront(original)
	toVisitB.PushFront(copy)

	for toVisitA.Len() > 0 {
		aEl := toVisitA.Back()
		toVisitA.Remove(aEl)
		a := aEl.Value.(*Node)

		bEl := toVisitB.Back()
		toVisitB.Remove(bEl)
		b := bEl.Value.(*Node)

		//same pointer to same instance
		if a == b {
			return false, fmt.Errorf("same node found in clone '%s'", a.label)
		}

		if a.label != b.label {
			return false, fmt.Errorf("'%s' <> '%s'",
				a.label, b.label)
		}

		if len(a.friends) != len(b.friends) {
			return false, fmt.Errorf("friends for '%s': should '%d', got '%d'",
				a.label, len(a.friends), len(b.friends))
		}

		visited[a.label] = struct{}{}

		for i, af := range a.friends {
			bf := b.friends[i]

			if af.label != bf.label {
				return false, fmt.Errorf("friend '%d' of '%s' is different: should '%s', got '%s'",
					i, a.label, af.label, bf.label)
			}

			if _, visited := visited[af.label]; visited {
				continue
			}

			toVisitA.PushFront(af)
			toVisitB.PushFront(bf)
		}
	}

	return true, nil
}
