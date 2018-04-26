package __7_intersection

import (
	"container/list"
)

/* 2.7
I Linked Lists
Intersection: Given two (singly) linked lists, determine if the two lists intersect. Return the
intersecting node. Note that the intersection is defined based on reference, not value. That is, if the
kth node of the first linked list is the exact same node (by reference) as the jth node of the second
linked list, then they are intersecting.
*/

func solution(a, b *list.List) *list.Element {
	//we will solve it by value not by reference
	//just because is less code to write
	//the only difference is that
	//the equality check would be node==node instead of values

	getLen := func(c *list.List) (count int) {
		for n := c.Front(); n != nil; n = n.Next() {
			count++
		}
		return
	}
	lenA, lenB := getLen(a), getLen(b) //or a.Len(), b.Len()
	short, long := a.Front(), b.Front()
	diff := lenB - lenA
	if lenA > lenB {
		short, long = long, short
		diff *= -1
	}

	for i := 0; i < diff; i++ {
		long = long.Next()
	}

	for short != nil { //short or long doesn't matter
		if short.Value.(int) == long.Value.(int) {
			return short
		}
		short, long = short.Next(), long.Next()
	}

	return nil
}
