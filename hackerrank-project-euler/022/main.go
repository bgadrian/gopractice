package main

import (
	"fmt"
	"log"
	"sort"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler022

func main() {
	names, inputs := input()
	scores := getScoreMap(names)
	for _, n := range inputs {
		fmt.Println(scores[n])
	}
}

//Because ALL the names are given from the begining AND there are no duplicates
//and no names are added after we can precalculate all of them
//and store them to a map so the query Lookup is O(1)
//for a short period of time we occupy O(3n) memory
func getScoreMap(names []string) map[string]int {
	sort.Strings(names)
	res := make(map[string]int, len(names))
	base := int('A') - 1
	var place, score int
	var r rune

	for index, name := range names {
		place = index + 1 //slices are zero based
		score = 0
		for _, r = range name {
			score += int(r) - base //A is 1
		}
		res[name] = score * place
	}
	return res
}

func input() (names, inputs []string) {
	var T int

	_, err := fmt.Scanln(&T)
	if err != nil {
		log.Panic(err)
	}

	var n string

	for l := T; l > 0; l-- {
		_, err = fmt.Scanln(&n)
		if err != nil {
			log.Panic(err)
		}

		names = append(names, n)
	}

	_, err = fmt.Scanln(&T)
	if err != nil {
		log.Panic(err)
	}

	for l := T; l > 0; l-- {
		_, err = fmt.Scanln(&n)
		if err != nil {
			log.Panic(err)
		}

		inputs = append(inputs, n)
	}

	return
}
