package main

import (
	"fmt"
	"log"
	"math"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler003

func main() {
	for _, n := range input() {
		//For each test case, print an integer that denotes the sum of all the multiples of  or  below .
		fmt.Println(getResult(n))
	}
}

func getResult(n int) int {
	maxPrime := -1

	for n%2 == 0 {
		maxPrime = 2
		n >>= 1
	}

	// n must be odd at this point, thus skip
	// the even numbers and iterate only for
	// odd integers
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			maxPrime = i
			n = n / i
		}
	}

	// This condition is to handle the case
	// when n is a prime number greater than 2
	if n > 2 {
		return n
	}

	return maxPrime
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
