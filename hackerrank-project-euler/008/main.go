package main

import (
	"fmt"
	"log"
	"strconv"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler008

type inputData struct {
	k int    //K
	n string //the number
	r int    //result
}

func main() {
	for _, n := range input() {
		fmt.Println(getResult(n))
	}
}

//
//func getResultIfDigitsAreNot0(in *inputData) int {
//	biggestProduct, _ := strconv.Atoi(string(in.n[0]))
//	var zeroASCIICode byte = 48
//
//	for i := 1; i < in.k; i++ {
//		biggestProduct *= int(in.n[i] - zeroASCIICode)
//	}
//
//	product := biggestProduct
//	for i := in.k; i < len(in.n); i++ {
//		if in.n[i-in.k] == zeroASCIICode || in.n[i] == in.n[i-in.k] {
//			continue
//		}
//		product /= int(in.n[i-in.k] - zeroASCIICode)
//		product *= int(in.n[i] - zeroASCIICode)
//
//		if product > biggestProduct {
//			biggestProduct = product
//		}
//	}
//
//	return biggestProduct
//}

func getResult(in *inputData) int {
	var tmp int
	var digits []int
	for _, v := range in.n {
		tmp, _ = strconv.Atoi(string(v))
		digits = append(digits, tmp)
	}
	l := len(digits)

	biggestProduct, tmp := 0, 0

	//O(n) time and O(n + k) space complexity
	for i := 0; i+in.k < l; i++ {
		tmp = getProdFrom(digits[i : i+in.k])
		if tmp > biggestProduct {
			biggestProduct = tmp
		}
	}

	return biggestProduct
}

func getProdFrom(digits []int) int {
	res := digits[0]

	for i, v := 1, 0; i < len(digits); i++ {
		v = digits[i]
		if v == 0 {
			return 0 //0 is a no go in a product
		}

		res *= v
	}

	return res
}

func input() (ints []*inputData) {
	var T int

	_, err := fmt.Scanln(&T)
	if err != nil {
		log.Panic(err)
	}

	for l := T; l > 0; l-- {
		dummy := 0
		newInput := &inputData{}
		fmt.Scan(&dummy, &newInput.k)
		fmt.Scanln(&newInput.n)
		ints = append(ints, newInput)
	}

	return
}
