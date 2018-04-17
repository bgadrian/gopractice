package test

/*
 Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you
cannot use additional data structures?
*/

func solution(s string) bool {
	//- all algorithms are linear time, we have to go trough all characters once (worst case)
	//1. using a hashset/map, would occupy O(n) extra memory (n*int)
	//2. if we know the domain of the characters and is limited (ex ASCII) we use only 255bits
	//using the ASCII code as the arrays index
	var m = make([]bool, 255) //default false
	for _, r := range s {
		if m[r] {
			return false
		}
		m[r] = true
	}

	return true
}

func solutionNoDataStructures(s string) bool {
	//do not use additional data structures
	//1. brute force - for each character search in the string for other occurence (>i)
	//2. we can apply an inplace sorting algorithm and we check for duplicates on consecutives indexs
	//but we still need a []rune array, because strings are immutable
	return false
}
