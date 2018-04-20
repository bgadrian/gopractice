package __6_string_compression

import "strconv"

/* 1.6
String Compression: Implement a method to perform basic string compression using the counts
of repeated characters. For example, the string aabcccccaaa would become a2blc5a3 . If the
"compressed" string would not become smaller than the original string, your method should return
the original string. You can assume the string has only uppercase and lowercase letters (a - z).*/

func solution(s []rune) []rune {
	if len(s) <= 1 {
		return s
	}

	var comp = []rune{s[0]}
	occurences := 1
	for i := 1; i < len(s); i++ {
		r := s[i]
		if comp[len(comp)-1] == r {
			occurences++
			continue
		}
		//are different
		comp = append(comp, []rune(strconv.Itoa(occurences))...)
		comp = append(comp, r)
		occurences = 1
	}
	comp = append(comp, []rune(strconv.Itoa(occurences))...)

	if len(s) < len(comp) {
		return s
	}
	return comp
}
