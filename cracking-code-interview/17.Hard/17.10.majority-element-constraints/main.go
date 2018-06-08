package _7_10_majority_element_constraints

/* 17.10 Majority Element: A majority element is an element that makes up more than half of the items in
an array. Given a positive integers array, find the majority element. If there is no majority element,
return -1. Do this in O( N) time and O( 1) space.
Input: 1 2 5 9 5 9 5 5 5
Output: 5 */

//Majority element or -1
func solution(a []int) int {
	mostFreq, count := 0, 0

	for _, v := range a {
		if count <= 0 {
			//start a new subset
			mostFreq = v
		}

		if mostFreq == v {
			count++
		} else {
			//reaches 0 when is not the maj in the subset
			count--
		}
	}

	count = 0
	for _, v := range a {
		if v == mostFreq {
			count++
		}
	}

	if count > len(a)/2 {
		return mostFreq
	}

	return -1
}
