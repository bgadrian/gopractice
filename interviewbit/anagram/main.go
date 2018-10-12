package anagram

import "bytes"
import "strconv"

//https://www.interviewbit.com/problems/anagrams/
func wordToHash(word string) string {
	//no time for a proper hash function
	//like a prime no for each letter in alphabet and multiply by count and so on..

	cache := make([]int, 27)
	for _, letter := range word {
		index := letter - 'a'
		cache[index]++
	}

	//var result strings.Builder
	var result bytes.Buffer
	for letter, count := range cache {
		if count == 0 {
			continue
		}
		result.WriteString(string(letter)) //WriteRune
		result.WriteString(":")
		result.WriteString(strconv.Itoa(count))
		result.WriteString(";")
	}
	return result.String()
}

func anagrams(words []string) [][]int {
	all := make(map[string][]int)

	for i, word := range words {
		hash := wordToHash(word)
		old := all[hash]
		old = append(old, i+1)
		all[hash] = old
	}

	var result [][]int
	for _, indexs := range all {
		result = append(result, indexs)
	}
	return result
}
