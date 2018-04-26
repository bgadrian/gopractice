package __2_kth_to_last

import (
	list2 "container/list"
	"log"
)

func solutionRec(a []int, k int) int {

	//recursive solution, we go to the end of the list, one stack for each element
	//and when we return we find the kth to last. Consumes a lot of memory, will crash for big lists
	//also Go doesn't support tail call optimization
	list := list2.New()
	for _, v := range a {
		list.PushBack(v)
	}

	type rec func(node *list2.Element, k int, r rec) (int, int)

	r := func(node *list2.Element, k int, r rec) (iReverse int, result int) {

		if node == nil {
			//start rolling back
			return 0, 0
		}
		next := node.Next()

		iReverse, result = r(next, k, r)
		iReverse++
		if k == iReverse {
			return iReverse, node.Value.(int)
		}

		return iReverse, result
	}

	_, result := r(list.Front(), k, r)
	return result
}

func solutionIt(a []int, k int) int {
	//iterative method, we go k nodes with an iterator
	//then we move 2 iterators until we reach the end
	list := list2.New()
	for _, v := range a {
		list.PushBack(v)
	}

	short, long := list.Front(), list.Front()
	for i := 0; i < k; i++ {
		if long == nil {
			log.Panic("k is longer then the length")
		}
		long = long.Next()
	}

	for long != nil {
		short = short.Next()
		long = long.Next()
	}

	return short.Value.(int)
}
