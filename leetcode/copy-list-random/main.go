package copy_list_random

type Node struct {
	next, random *Node
}

/* A linked list is given such that each node contains an additional random pointer which could point to any node in the list or NULL.

Return a deep copy of the list.

Example

Given list

   1 -> 2 -> 3
with random pointers going from

  1 -> 3
  2 -> 1
  3 -> 1
You should return a deep copy of the list. The returned answer should not contain the same node as the original list, but a copy of them. The pointers in the returned list should not link to any node in the original input list.

O(n) time and O(n) space
*/
func solution(head *Node) *Node {
	cache := make(map[*Node]*Node)
	result := &Node{}
	cache[head] = result

	get := func(input *Node) *Node {
		out, ok := cache[input]
		if ok == false {
			out = &Node{}
			cache[input] = out
		}
		return out
	}

	for currInput := head; currInput != nil; {
		nextInput := currInput.next
		randInput := currInput.random

		currOutput := get(currInput)
		if currInput.next != nil {
			nextOutput := get(nextInput)
			currOutput.next = nextOutput
		}

		if currInput.random != nil {
			randOutput := get(randInput)
			currOutput.random = randOutput
		}
		currInput = currInput.next
	}

	return result
}
