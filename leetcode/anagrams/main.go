package anagrams

//convert a word in a number
//no matter the position of the letters
//in the word, the result is the same
//does NOT work if the letters repeat
//for that you need to use the sorted string as key
var hashToNum = func(a string) int32 {
	var result int32
	//26 letters, we put a 1 bit in each location
	for _, r := range a {
		position := uint(r - 'a')
		mask := int32(1) << position
		result |= mask
	}
	return result
}

/**
Given an array of strings, return all groups of strings that are anagrams. Represent a group by a list of integers representing the index in the original list. Look at the sample case for clarification.

 Anagram : a word, phrase, or name formed by rearranging the letters of another, such as 'spar', formed from 'rasp'
 Note: All inputs will be in lower-case.
Example :

Input : cat dog god tca
Output : [[1, 4], [2, 3]]
cat and tca are anagrams which correspond to index 1 and 4.
dog and god are another set of anagrams which correspond to index 2 and 3.
The indices are 1 based ( the first element has index 1 instead of index 0).

 Ordering of the result : You should not change the relative ordering of the words / phrases

https://www.interviewbit.com/problems/anagrams/

*/
func anagrams(input []string) [][]int {
	cache := make(map[int32][]int)

	for i, w := range input {
		h := hashToNum(w)
		cache[h] = append(cache[h], i+1)
	}

	var result [][]int
	for _, v := range cache {
		result = append(result, v)
	}
	return result
}
