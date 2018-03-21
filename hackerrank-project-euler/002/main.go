package main

import (
	"fmt"
	"log"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler002

func main() {
	for _, n := range input() {
		//For each test case, print an integer that denotes the sum of all the multiples of  or  below .
		fmt.Println(getResult(n))
	}
}

func getResult(n int) int {
	//Recurrence for Even Fibonacci sequence is:
	//EFn = 4EFn-1 + EFn-2
	//with seed values
	//EF0 = 0 and EF1 = 2.
	if n < 2 {
		return 0
	}

	ef0, ef1 := 0, 2
	sum := ef0 + ef1
	var ef2 int

	for ef1 <= n {
		ef2 = 4*ef1 + ef0
		if ef2 > n {
			break
		}
		ef1, ef0 = ef2, ef1
		sum += ef1
	}
	return sum
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
