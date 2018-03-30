package main

import (
	"fmt"
	"log"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler010

func main() {
	for _, n := range input() {
		fmt.Println(getResult(n))
	}
}

//to pass all tests, the trick is to store/memoize the already calculated values
var arr = []int{2, 3}
var counter = 4

func getResult(n int) int {
	p, sum := 0, 0
	for i := 1; p <= n; i++ {
		sum += p
		p = getNthPrime(i)
	}

	return sum
}

func getNthPrime(n int) int {
	//https://en.wikipedia.org/wiki/Formula_for_primes
	tmp := 0

	for len(arr) < n {
		if counter%2 != 0 && counter%3 != 0 {
			tmp = 4
			for tmp*tmp <= counter {
				if counter%tmp == 0 {
					break
				}
				tmp++
			}
			if tmp*tmp > counter {
				arr = append(arr, counter)
			}
		}
		counter++
	}
	return arr[n-1]
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
