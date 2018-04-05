package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler017

var groups = []string{"", "Thousand", "Million", "Billion", "Trillion"}

//var groupsPlural = []string{"", "Thousands", "Millions", "Billions", "Trillions"}
var digits = []string{"Zero", "One", "Two", "Three", "Four",
	"Five", "Six", "Seven", "Eight", "Nine"}
var tens = []string{"Ten", "Eleven", "Twelve", "Thirteen", "Fourteen",
	"Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
var tw = []string{"", "", "Twenty", "Thirty", "Forty", "Fifty",
	"Sixty", "Seventy", "Eighty", "Ninety"}

func main() {
	for _, n := range input() {
		fmt.Println(getResult(n))
	}
}

func getResult(n int) string {
	//quick exit
	if n < 10 {
		return digits[n]
	}

	//TODO group these to a data structure
	var res []string
	asStr := strconv.Itoa(n)
	var curr, groupIndex int
	groupWasZero := true
	digit := len(asStr)

	addGroup := func() {
		groupWasZero = true
		if digit < 3 {
			return
		}
		groupIndex = int(math.Floor(float64(digit) / 3))
		res = append(res, groups[groupIndex])
	}

	next := func() {
		curr, _ = strconv.Atoi(string(asStr[0]))
		asStr = asStr[1:]
	}

	for ; digit >= 1; digit-- {
		next()

		if digit%3 == 0 {
			if curr == 0 {
				continue
			}
			groupWasZero = false

			//hundreds
			res = append(res, digits[curr])
			res = append(res, "Hundred")
			continue
		}

		//tens
		if digit%3 == 2 {
			if curr == 0 {
				continue
			}
			groupWasZero = false

			if curr == 1 {
				//special case, 10-19
				next()
				digit--
				res = append(res, tens[curr])
				addGroup()
				continue
			}
			//tens digit is >= 2
			res = append(res, tw[curr])
			continue
		}

		//is the last digit
		if curr > 0 {
			res = append(res, digits[curr])
			groupWasZero = false
		}

		if groupWasZero == false {
			addGroup()
		}
	}

	return strings.Join(res, " ")
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
