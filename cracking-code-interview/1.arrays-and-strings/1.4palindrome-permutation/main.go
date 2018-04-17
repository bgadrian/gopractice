package __4palindrome_permutation

import "unicode"

/*
1.4 Palindrome Permutation: Given a string, write a function to check if it is a permutation of
a palindrome. A palindrome is a word or phrase that is the same forwards and backwards. A
permutation is a rearrangement of letters. The palindrome does not need to be limited to just
dictionary words.
EXAMPLE
Input: Tact Coa
Output: True (permutations:"taco cat'; "atco cta'; etc.)
*/

func solution(s []rune) bool {

	//to have a palindrome we must have at most 1 odd occurrence count of distinct character

	//if the characters are only ASCII, we can make a 255 bits array
	//0 value means the [index] character has even occurrences or 0
	//1 means it has odd
	//at the end we must have at most 1 value of 1
	//int64 has only 64bits, so we can use it only if we have max 64 distinct characters
	//we can presume we only have letters (26) so is safe to use for now even a int32, otherwise we will make an array of booleans

	allowed := uint(1)
	var count uint
	for _, r := range s {
		//we ignore punctuation
		if unicode.IsLetter(r) == false {
			continue
		}
		rc := unicode.ToLower(r)
		bit := uint(rc - 'a')
		count ^= 1 << bit
	}

	odds := uint(0)
	for bit := uint(0); bit < 32; bit++ {
		if count&(1<<bit) == 0 { //has bits[bit] == 0, even, ignore
			continue
		}
		odds++
		if odds > allowed {
			return false
		}
	}

	return true
}
