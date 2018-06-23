package __add_two_numbers_ll

type ListNode struct {
	Val  int
	Next *ListNode
}

/* You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	next := func(curr *ListNode) *ListNode {
		if curr.Next == nil {
			curr.Next = &ListNode{}
		}
		return curr.Next
	}

	type overflow func(to *ListNode, f overflow)

	over := func(to *ListNode, f overflow) {
		if to == nil || to.Val < 10 {
			return
		}

		nextNode := next(to)
		nextNode.Val += 1
		to.Val -= 10
		f(nextNode, f)
	}

	add := func(what *ListNode, to *ListNode) {
		for what != nil {
			to.Val += what.Val
			over(to, over) //check for overflow
			what = what.Next

			if what == nil {
				break
			}
			to = next(to)
		}
	}

	add(l1, l2) //or create a new head and add(l1, head) add (l2, head)
	return l2
}
