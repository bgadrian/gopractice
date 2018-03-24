package main

import (
	"fmt"
	"log"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler009

func main() {
	for _, n := range input() {
		fmt.Println(getResult(n))
	}
}

func getResult(sum int) int {
	semiPerim := sum / 2 //triangle inequality https://en.wikipedia.org/wiki/List_of_triangle_inequalities
	var cc, a, product int
	maxProduct := -1
	for c := semiPerim - 1; c >= 3; c-- { //c > b > a
		cc = c * c
		for b := c - 1; b >= c/2; b-- {
			a = sum - c - b
			if a*a+b*b != cc {
				continue
			}
			product = a * b * c
			if product > maxProduct {
				maxProduct = product
			}
		}
	}
	return maxProduct
	//TODO linear time formula (not tested) https://stackoverflow.com/a/44307023/1600231
	//TODO for more performance try generating? https://en.wikipedia.org/wiki/Pythagorean_triple#Generating_a_triple
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
