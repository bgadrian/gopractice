package main

import (
	"fmt"
	"log"
	"math"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler024

func main() {
	for _, n := range input() {
		fmt.Println(getResult("abcdefghijklm", n))
	}
}

var fm = map[int]int{0: 1, 1: 1}

func fact(n int) int {
	if r, ok := fm[n]; ok {
		return r
	}

	fm[n] = n * fact(n-1)
	return fm[n]
}

func getResult(word string, k int) string {

	s := []rune(word)
	var p []rune
	k-- //zero based algorithm

	for len(s) > 0 {
		f := fact(len(s) - 1)
		i := int(math.Floor(float64(k) / float64(f)))
		x := s[i]
		k = k % f
		p = append(p, x)
		//remove x from s
		copy(s[i:], s[i+1:])
		s = s[:len(s)-1]
	}
	return string(p)
}

func input() (ints []int) {
	var T int

	_, err := fmt.Scanln(&T)
	if err != nil {
		log.Panic(err)
	}

	var n int

	for l := T; l > 0; l-- {
		_, err = fmt.Scanln(&n)
		if err != nil {
			log.Panic(err)
		}

		ints = append(ints, n)
	}

	return
}
