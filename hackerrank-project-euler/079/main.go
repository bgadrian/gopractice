package main

import (
	"fmt"
	"log"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler079/problem

func main() {
	fmt.Println(string(getResult(input())))
}

func getResult(in [][]rune) (res []rune) {
	//key is the character, values are the dependencies that must be solved first
	//aka appear before in the result
	edges := make(map[rune][]rune)

	for _, rs := range in {
		for i, r := range rs {
			if _, exists := edges[r]; exists == false {
				edges[r] = make([]rune, 0)
			}

			if i > 0 {
				edges[r] = append(edges[r], rs[i-1])
			}
		}
	}

	var nextLetter rune

	for len(edges) > 0 {
		//search for the next available letter (no depe, smallest one because lex order
		for char, dependencies := range edges {
			if len(dependencies) > 0 {
				continue
			}
			if nextLetter == 0 || char < nextLetter {
				nextLetter = char
			}
		}

		if nextLetter == 0 {
			return []rune("SMTH WRONG")
		}

		//remove the letter from dependencies (remove the onion level)
		delete(edges, nextLetter)
		for char := range edges {
			for i, v := range edges[char] {
				if v != nextLetter {
					continue
				}

				edges[char] = append(edges[char][:i], edges[char][i+1:]...)
				break //values are unique
			}
		}
		res = append(res, nextLetter)
		nextLetter = 0
	}

	return
}

func input() (text [][]rune) {
	var T int64

	_, err := fmt.Scanln(&T)
	if err != nil {
		log.Panic(err)
	}

	var r string

	for l := T; l > 0; l-- {
		_, err = fmt.Scan(&r)
		if err != nil {
			log.Panic(err)
		}

		text = append(text, []rune(r))
	}

	return
}
