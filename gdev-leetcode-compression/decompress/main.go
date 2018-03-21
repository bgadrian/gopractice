/*main
https://techdevguide.withgoogle.com/paths/advanced/compress-decompression/#code-challenge
In this exercise, you're going to decompress a compressed string.
Your input is a compressed string of the format number[string] and the decompressed output form should be the string written number times. For example:
The input
3[abc]4[ab]c
Would be output as
abcabcabcababababc
*/
package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(decompressRecursive([]rune(os.Args[1])))
		return
	}

	fmt.Println("missing argument")
}

func decompressIterative(w []rune) []rune {

	stacks := [][]rune{{}}
	counts := []int{0}
	index := 0

	for i, c := range w {
		if c == ']' {
			stacks[index-1] = append(stacks[index-1], repeatInt(stacks[index], counts[index])...)
			stacks = stacks[0:index]
			counts = counts[0:index]
			index--
			continue
		}

		if unicode.IsNumber(c) {
			isFirstDigit := i == 0 || unicode.IsNumber(w[i-1]) == false
			if isFirstDigit {
				//go deeper, new stack - child
				index++
				stacks = append(stacks, []rune{})
				counts = append(counts, 0)
			}

			cAsInt, _ := strconv.Atoi(string(c))
			counts[index] = counts[index]*10 + cAsInt
			continue
		}
		if unicode.IsLetter(c) {
			stacks[index] = append(stacks[index], c)
			continue
		}

		//if c == '[' {
		//}
	}

	if len(stacks) > 1 {
		panic("something went wrong, we missed something")
	}

	return stacks[0]
}

/*
	Slices share the same internal array, so memory is O(n)
	It doesn't go trough the same chars twice, so runtime O(n)
*/
func decompressRecursive(w []rune) ([]rune, int) {
	var times, result, inner []rune
	i := 0
	skip := 0
	innerCount := 0

	for ; i < len(w); i++ {
		c := w[i]
		if unicode.IsNumber(c) {
			times = append(times, c)
			continue
		}
		if unicode.IsLetter(c) {
			result = append(result, c)
			continue
		}

		if c == '[' {
			inner, skip = decompressRecursive(w[i+1:])
			i += skip
			innerCount++
			continue
		}

		//means c == ']'
		innerCount--
		if innerCount < 0 {
			break
		}

		result = append(result, repeat(inner, times)...)
		times = times[:0] //clear
		inner = inner[:0]
	}

	return result, i
}

//Like string.Repeat but with rune slices
func repeat(w []rune, times []rune) (r []rune) {
	c := string(times)
	l, _ := strconv.Atoi(c)
	for ; l > 0; l-- {
		r = append(r, w...)
	}
	return
}

func repeatInt(w []rune, times int) (r []rune) {
	for ; times > 0; times-- {
		r = append(r, w...)
	}
	return
}
