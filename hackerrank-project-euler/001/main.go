package main

import (
	"fmt"
	"log"
	"math"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler001/problem

func main() {
	for _, n := range input() {
		//For each test case, print an integer that denotes the sum of all the multiples of  or  below .
		fmt.Println(getResult(n))
	}
}

func getResult(n int64) int64 {
	return getSumsMultipliers(n, 3) +
		getSumsMultipliers(n, 5) -
		getSumsMultipliers(n, 15)
}

func getSumsMultipliers(n, div int64) int64 {
	maxPower := int64(math.Floor(float64(n) / float64(div)))
	maxNumber := maxPower * div

	//<= n
	if maxNumber == n {
		maxNumber -= div
		maxPower--
	}
	if maxPower <= 0 {
		return 0
	}
	sum := (maxPower * (div + maxNumber)) / 2
	return sum
}

func input() (ints []int64) {
	var T int64

	_, err := fmt.Scanln(&T)
	if err != nil {
		log.Panic(err)
	}

	var n int64

	for l := T; l > 0; l-- {
		_, err = fmt.Scanln(&n)
		if err != nil {
			log.Panic(err)
		}

		ints = append(ints, n)
	}

	return
}
