package main

import (
	"fmt"
	"sort"
)

var k = 4

//https://www.hackerrank.com/contests/projecteuler/challenges/euler011

func main() {
	fmt.Println(getResult(input(20)))
}

func getResult(n [][]int) int {
	//TODO multiple goroutines
	results := []int{getResultCols(n), getResultRows(n)}
	results = append(results, getResultDiag(n)...)

	sort.Ints(results)
	return results[len(results)-1]
}

func getResultCols(n [][]int) (res int) {
	tmp := 0
	l := len(n)

	for col := 0; col < l; col++ {
		for row := 0; row <= l-k; row++ {
			tmp = n[row][col]
			for j := 1; j < k; j++ {
				tmp *= n[row+j][col]
				if tmp == 0 {
					break
				}
			}
			if tmp > res {
				res = tmp
			}
		}
	}
	return
}

func getResultRows(n [][]int) (res int) {
	tmp := 0
	l := len(n)

	for row := 0; row < l; row++ {
		for col := 0; col <= l-k; col++ {
			tmp = n[row][col]
			for j := 1; j < k; j++ {
				//TODO we can save n-k multiplication ops by using memoization
				tmp *= n[row][col+j]
				if tmp == 0 {
					break
				}
			}
			if tmp > res {
				res = tmp
			}
			//fmt.Printf("[%v][%v-%v]\n", row, col, (col + k - 1))
		}
	}
	return

}

func getResultDiag(n [][]int) (res []int) {
	l := len(n)
	tmp := 0

	leftToRight := func(row, col int) (tres int) {
		//for all left column values
		for c, r := col, row; c <= l-k && r <= l-k; c, r = c+1, r+1 {
			//starting with all values on diagonal
			tmp = n[r][c]
			//fmt.Printf("[%v][%v] - [%v][%v]\n", r, c, r+k-1, c+k-1)
			for j := 1; j < k; j++ {
				tmp *= n[r+j][c+j]
				if tmp == 0 {
					break
				}
			}
			if tmp > tres {
				tres = tmp
			}
		}
		return
	}

	rightToLeft := func(row, col int) (tres int) {
		//fmt.Printf("[%v][%v]\n", row, col)
		for c, r := col, row; c >= k-1 && r <= l-k; c, r = c-1, r+1 {
			//starting with all values on diagonal
			tmp = n[r][c]
			//fmt.Printf("[%v][%v] - [%v][%v]\n", r, c, r+k-1, c-k-1)
			for j := 1; j < k; j++ {
				tmp *= n[r+j][c-j]
				if tmp == 0 {
					break
				}
			}
			if tmp > tres {
				tres = tmp
			}
		}
		return
	}

	//top left to bottom right diagonals
	for col, row := 0, l-k; row >= 0; row-- { //left side
		res = append(res, leftToRight(row, col))
	}
	for col, row := 0, 1; col <= l-k; col++ { //top side
		res = append(res, leftToRight(row, col))
	}

	//top right to bottom left
	for col, row := l-1, 0; col >= k-1; col-- { //top
		res = append(res, rightToLeft(row, col))
	}
	for col, row := l-1, 1; row <= l-k; row++ { //right
		res = append(res, rightToLeft(row, col))
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
