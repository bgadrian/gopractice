package clone_graph

import (
	"container/list"
	"fmt"
	"strings"
)

/* Clone an undirected graph. Each node in the graph contains a label and a list of its neighbors.

https://www.interviewbit.com/problems/clone-graph/
https://leetcode.com/problems/clone-graph/description/
*/

type Node struct {
	label   string
	friends []*Node
}

func (n Node) String() string {
	res := fmt.Sprintf("'%s' -> ", n.label)
	var fs []string
	for _, friend := range n.friends {
		fs = append(fs, friend.label)
	}
	return res + strings.Join(fs, ",")
}

// the graph must be connected. Nstart is any node
func Solution(nstart *Node) *Node {
	//we need a way to differentiate the copies
	copies := make(map[*Node]*Node)
	toVisit := list.New()
	var result *Node

	toVisit.PushFront(nstart)
	for toVisit.Len() > 0 {
		orig := toVisit.Back()
		toVisit.Remove(orig)

		node := orig.Value.(*Node)
		clone := &Node{label: node.label}
		//add the clone + mark it as visited
		copies[node] = clone

		if result == nil {
			//it is the first node
			result = clone
		}

		for _, f := range node.friends {
			friendClone, wasCloned := copies[f]

			if wasCloned {
				//only add the link when we found
				//the other end
				a, b := clone, friendClone
				a.friends = append(a.friends, b)
				b.friends = append(b.friends, a)
				continue
			}

			toVisit.PushFront(f)
		}
	}
	return result
}
