package main

import (
	"fmt"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler011

func main() {
	fmt.Println(getResult(input(20), 4))
}

func getResult(n [][]int, k int) (res int) {
	var tmp, j int
	l := len(n) //presume is a matrix/square
	maxIndex := l - k

	for row, r := range n {
		for col, v := range r {
			//row forward (top down === down top)
			if col <= maxIndex {
				tmp = v
				for j = 1; j < k; j++ {
					tmp *= n[row][col+j]
					if tmp == 0 {
						break
					}
				}
				if tmp > res {
					res = tmp
				}
			}

			//column bottom (left right === right left)
			if row <= maxIndex {
				tmp = v
				for j = 1; j < k; j++ {
					tmp *= n[row+j][col]
					if tmp == 0 {
						break
					}
				}
				if tmp > res {
					res = tmp
				}
			}

			//diagonal top left - bottom right
			if row <= maxIndex && col <= maxIndex {
				tmp = v
				for j = 1; j < k; j++ {
					tmp *= n[row+j][col+j]
					if tmp == 0 {
						break
					}
				}
				if tmp > res {
					res = tmp
				}
			}

			//diagonal top right - bottom left
			if row <= maxIndex && col >= k {
				tmp = v
				for j = 1; j < k; j++ {
					tmp *= n[row+j][col-j]
					if tmp == 0 {
						break
					}
				}
				if tmp > res {
					res = tmp
				}
			}
		}
	}
	return
}
func input(n int) (ints [][]int) {

	for l := n; l > 0; l-- {
		row := make([]int, n)
		for i := range row {
			fmt.Scan(&row[i])
		}
		ints = append(ints, row)
	}

	return
}
