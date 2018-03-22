package main

import (
	"fmt"
	"log"
	"strconv"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler004

func main() {
	for _, n := range input() {
		//For each test case, print an integer that denotes the sum of all the multiples of  or  below .
		fmt.Println(getResult(n))
	}
}

func getResult(n int) int {
	//TODO a more optimal solution is to hard code all the 6 digit palindromes in the code

	var tmp int
	res := 0
	for a := 999; a >= 100; a-- {
		for b := a; b >= 100; b-- { //comm a*b == b*a
			tmp = a * b
			if tmp >= n {
				continue
			}

			if isPalindrom(tmp) && tmp > res {
				//fmt.Printf("%v x %v = %v\n", a, b, tmp)
				res = tmp
			}
		}
	}
	return res
}

func isPalindrom(n int) bool {
	s := strconv.Itoa(n)
	//TODO more performant? division instead of string compare?
	l := len(s)

	//All palindromes with an even number of digits are divisible by 11.
	//http://jwilson.coe.uga.edu/emt669/Student.Folders/Bacchus.Mohamed/pal/pal.html
	if l%2 == 0 && n%11 > 0 {
		return false
	}

	for i := 0; i < l/2; i++ {
		if s[i] != s[l-i-1] {
			return false
		}
	}
	return true
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
