package _48_shortest_completing_word_

import "strings"

/*
Find the minimum length word from a given dictionary words, which has all the letters from the string licensePlate. Such a word is said to complete the given string licensePlate
Here, for letters we ignore case. For example, "P" on the licensePlate still matches "p" on the word.
It is guaranteed an answer exists. If there are multiple answers, return the one that occurs first in the array.
The license plate might have the same letter occurring multiple times. For example, given a licensePlate of "PP", the word "pair" does not complete the licensePlate, but the word "supper" does.

Example 1:
Input: licensePlate = "1s3 PSt", words = ["step", "steps", "stripe", "stepple"]
Output: "steps"
Explanation: The smallest length word that contains the letters "S", "P", "S", and "T".
Note that the answer is not "step", because the letter "s" must occur in the word twice.
Also note that we ignored case for the purposes of comparing whether a letter exists in the word.
Example 2:
Input: licensePlate = "1s3 456", words = ["looks", "pest", "stew", "show"]
Output: "pest"
Explanation: There are 3 smallest length words that contains the letters "s".
We return the one that occurred first.
Note:
licensePlate will be a string with length in range [1, 7].
licensePlate will contain digits, spaces, or letters (uppercase or lowercase).
words will have a length in the range [10, 1000].
Every words[i] will consist of lowercase letters, and have length in range [1, 15].
*/

/*
The solution is procedural, O(n) time, the clone operations
are actually O(n * k) where k is the count of letters in license
The worst case scenario is when the solution is the last word,
and the words are in desc length order.
*/
func shortestCompletingWord(licensePlate string, words []string) string {
	uniq, clone := make(map[rune]byte), make(map[rune]byte)

	copyMap := func(source, destination map[rune]byte) {
		// for k := range destination {
		//     delete(destination, k)
		// }

		for k, v := range source {
			destination[k] = v
		}
		return
	}

	for _, r := range strings.ToLower(licensePlate) {
		if r < 'a' || r > 'z' { //replace with unicode.IsLetter(r) in a real project
			continue
		}
		uniq[r]++
	}

	result := ""
	var letter rune
	for _, w := range words { //find the first word
		if len(w) < len(uniq) {
			continue //no use to check here
		}
		if len(result) > 0 && len(result) <= len(w) {
			continue //no use to check here, cannot be better
		}

		copyMap(uniq, clone)
		for _, letter = range w { //for each letter in word w
			occ, has := clone[letter]
			if has == false { //is not in the license plate, or was already found in w, we ignore it
				continue
			}

			if occ > 1 {
				clone[letter]--
				continue
			}
			//occurenes is 1
			delete(clone, letter)

			if len(clone) == 0 { //word w has all the letters
				result = w
				break
			}
		}
	}

	return result
}
