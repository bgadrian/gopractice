package main

import (
	"fmt"
	"log"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler006

func main() {
	for _, n := range input() {
		fmt.Println(getResult(n))
	}
}

func getResult(n int) int {
	//arithmetic series (1+2+3+...+n)
	sumN := int((n * (1 + n)) / 2)
	sumSquare := sumN * sumN

	//https://trans4mind.com/personal_development/mathematics/series/sumNaturalSquares.htm
	sumOfSquares := int((n * (n + 1) * (2*n + 1)) / 6)

	//absolute for ints
	res := sumSquare - sumOfSquares
	if res < 0 {
		res *= -1
	}
	return res
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
