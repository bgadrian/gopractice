package main

import (
	"fmt"
	"log"
	"math/big"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler116

func main() {
	for _, n := range input() {
		fmt.Println(getResult(n))
	}
}

func getResult(n int) *big.Int {
	sum := getCombinationsOf(n, 2)        //red
	sum.Add(sum, getCombinationsOf(n, 3)) //green
	sum.Add(sum, getCombinationsOf(n, 4)) //blues
	return sum
}

func getCombinationsOf(black, colored int) *big.Int {
	//got it from https://github.com/nayuki/Project-Euler-solutions/blob/master/java/p116.java
	//I'm pretty sure a O(1) combinatorics formula exists but I can't find it yet
	cache := make([]*big.Int, black+1)
	cache[0] = big.NewInt(1)

	for i := 1; i <= black; i++ {
		cache[i] = big.NewInt(0).Set(cache[i-1])
		if i < colored {
			continue
		}
		cache[i].Add(cache[i], cache[i-colored])
	}

	return cache[black].Add(cache[black], big.NewInt(-1)) //exclude all blacks
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
