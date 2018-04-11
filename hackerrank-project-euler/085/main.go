package main

import (
	"fmt"
	"log"
	"math"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler085

func main() {
	for _, n := range input() {
		fmt.Println(getResult(n))
	}
}

func getResult(n int) int {
	cols, rows := n, 1
	var rects, closestArea, diff int
	closestRects := math.MaxInt64

	for cols >= rows {
		rects = (cols * (cols + 1) * rows * (rows + 1)) / 4

		diff = int(math.Abs(float64(n - rects)))

		if diff < closestRects {
			closestRects = diff
			closestArea = cols * rows
		}

		if rects > n {
			cols--
			continue
		}
		rows++
	}
	return closestArea
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
