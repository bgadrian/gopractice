package __3URLify_replace_spaces

import (
	"log"
	"unicode"
)

/*
1.3 URLify: Write a method to replace all spaces in a string with '%20': You may assume that the string
has sufficient space at the end to hold the additional characters, and that you are given the "true"
length of the string. (Note: if implementing in Java, please use a character array so that you can
perform this operation in place.)
EXAMPLE
Input: "Mr John Smith    ', 13
Output: "Mr%20John%20Smith"
*/
func solution(s []rune, trueLength int) string {
	spaceCount := int((len(s) - trueLength) / 2)
	//two pointers, one for the empty spaces at the end and one for the traversing in reverse
	emptyIndex := len(s) - 1
	i := trueLength - 1

	for spaceCount > 0 { //while there are left spaces to replace
		r := s[i]
		if unicode.IsSpace(r) {
			s[emptyIndex] = '0'
			s[emptyIndex-1] = '2'
			s[emptyIndex-2] = '%'
			emptyIndex -= 3
			spaceCount--
		} else {
			s[emptyIndex] = s[i]
			emptyIndex--
		}
		i--

		//sanity check
		if i < 0 {
			log.Panic("my head is on fire!!")
		}
	}

	return string(s)
}
