package longest_substr_distinct_chars

import "log"

// solution Given a string, find the longest substring that contains only K  unique characters.
// For example, given "abcbbbbcccbdddadacb", the longest substring that contains 2 unique character is "bcbbbbcccb".
// we presume "s" is ASCII, we don't use utf8 functions

//window represents a substring
type window struct {
	start, end int
}

func (w window) size() int {
	return w.end - w.start + 1
}

func solution(s string, k int) string {
	if k < 1 {
		log.Panicln("k must be >= 1")
	}

	lastSeenAt := make(map[rune]int)
	curr := window{0, 1}
	max := curr

	for i, r := range s {
		if _, ok := lastSeenAt[r]; ok {
			//already exists
			lastSeenAt[r] = i
		} else {
			//else, is a new distinct char
			//we need to remove first, the leftmost current char
			if len(lastSeenAt) == k {
				var toRem rune
				min := len(s)
				for let, lastRight := range lastSeenAt {
					if lastRight >= min {
						continue
					}
					min = lastRight
					toRem = let
				}
				curr.start = min + 1
				delete(lastSeenAt, toRem)
			}
		}
		curr.end = i

		if curr.size() > max.size() {
			max = curr
		}
		lastSeenAt[r] = i
	}

	res := s[max.start : max.end+1]
	return res
}
