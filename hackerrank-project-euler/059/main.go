package main

import (
	"fmt"
	"log"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler059/problem

func main() {
	fmt.Println(string(getResult(input())))
}

func getResult(in []rune) (key []rune) {
	//the key has 3 characters, so 3 slots, being an array is 0,1,2
	const k1, k2, k3 = 0, 1, 2
	remainders := map[int]int{
		1: k1,
		2: k2,
		0: k3,
	}
	//a list of hashsets, for each key slot
	ks := []map[rune]struct{}{{}, {}, {}}
	//populate all possible key values
	for i := 'a'; i <= 'z'; i++ {
		ks[k1][i] = struct{}{} //hahsets have empty values (null in java)
		ks[k2][i] = struct{}{}
		ks[k3][i] = struct{}{}
	}

	//runtime of O(n * log(k)) where k is (122-97) possible values
	for i, r := range in {
		ki := remainders[(i+1)%3]
		for kv := range ks[ki] {
			if isValid(r^kv) == false {
				//with each failed try the possible key values shrink
				delete(ks[ki], kv)
			}
		}
	}

	//we make sure that is only 1 character left for each key slot, otherwise is not unique or the algorithm failed
	for _, ki := range []rune{k1, k2, k3} {
		if len(ks[ki]) != 1 {
			log.Panicf("k%v was not found, something went wrong '%v'", ki+1, ks[ki])
		}
		for kv := range ks[ki] {
			key = append(key, kv)
		}
	}

	return
}

// (a-z, A-Z, 0 - 9), brackets (), common symbols (;:,.'?-!) and spaces
var lookupIsValid = map[rune]struct{}{
	'(':  {},
	')':  {},
	';':  {},
	':':  {},
	',':  {},
	'.':  {},
	'\'': {},
	'?':  {},
	'-':  {},
	'!':  {},
	' ':  {},
}

//constant time lookup
func isValid(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	if r >= 'A' && r <= 'Z' {
		return true
	}
	if r >= '0' && r <= '9' {
		return true
	}

	if _, ok := lookupIsValid[r]; ok {
		return true
	}
	return false
}

func input() (text []rune) {
	var T int64

	_, err := fmt.Scanln(&T)
	if err != nil {
		log.Panic(err)
	}

	var r rune

	for l := T; l > 0; l-- {
		_, err = fmt.Scan(&r)
		if err != nil {
			log.Panic(err)
		}

		text = append(text, r)
	}

	return
}
