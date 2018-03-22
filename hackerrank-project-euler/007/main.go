package main

import (
	"fmt"
	"log"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler007

func main() {
	for _, n := range input() {
		fmt.Println(getResult(n))
	}
}

func getResult(n int) int {
	//https://en.wikipedia.org/wiki/Formula_for_primes
	arr := []int{2, 3}
	counter := 4
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
	return arr[len(arr)-1]
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
