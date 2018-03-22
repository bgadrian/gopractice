package main

import (
	"fmt"
	"log"
	"math"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler005

func main() {
	for _, n := range input() {
		fmt.Println(getResult(n))
	}
}

func getResult(n int) int {
	//for better algorithms see https://en.wikipedia.org/wiki/Least_common_multiple
	//bruteforce:
	for i := n; i < math.MaxInt64; i += n {
		ok := true
		for d := n - 1; d > 1; d-- {
			if i%d > 0 {
				ok = false
				break
			}
		}
		if ok {
			return i
		}
	}
	return 0
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
