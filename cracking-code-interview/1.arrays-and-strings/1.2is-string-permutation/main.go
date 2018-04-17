package __2is_string_permutation

import "sort"

/* 1.2 Check Permutation: Given two strings, write a method to decide if one is a permutation of the
other.
*/

func solution(a, b []rune) bool {
	//1. brute force, generate all permutations and search if is the same as B
	//2. check if all the letters and their occurrences count are the same
	//3. sorting the characters and compare directly

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return string(a) == string(b)
}
