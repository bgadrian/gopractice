package permutations1ton

import "log"

//check if the input is a permutation of 1..N ints where N is the size of A
//N is max 100.000 so it fits in memory with a map
//struct{} occupies 0 memory so its ok
//Linear time and N space
func Solution(A []int) int {
	size := len(A)
	if size == 0 {
		log.Panic("you lied")
	}

	if size == 1 {
		if A[0] == 1 {
			return 1
		}
		return 0
	}

	duplicates := make(map[int]struct{}, size)
	for _, no := range A {
		if no > size {
			return 0 //if is too big it means is missing smaller one
		}
		if _, exists := duplicates[no]; exists {
			return 0 //duplicate
		}
		duplicates[no] = struct{}{}
	}
	return 1
}
