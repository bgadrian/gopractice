package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler096
// SUDOKU solved with Constraint Propagation

func main() {
	for _, strRow := range getResult(input()) {
		fmt.Println(strRow)
	}
}

func newSudoku(lines []string) *sudoku {
	s := &sudoku{}
	s.rows = make(map[int][]*cell)
	s.cols = make(map[int][]*cell)
	s.squares = make(map[int][]*cell)

	for line, lineText := range lines {
		row := line + 1
		for head, asStr := range strings.Split(lineText, "") {
			col := head + 1
			square := int((math.Floor((float64(col)-1)/3) + 1) + //1,2,3
				math.Floor((float64(row)-1)/3)*3) //0,3,6
			asInt, _ := strconv.Atoi(asStr)
			c := newCell(asInt)

			c.col = col
			c.row = row
			c.square = square

			s.rows[row] = append(s.rows[row], c)
			s.cols[col] = append(s.cols[col], c)
			s.squares[square] = append(s.squares[square], c)
			s.list = append(s.list, c)
		}
	}

	return s
}

func (s *sudoku) isSolved() bool {
	for _, cell := range s.list {
		if cell.isFixed() == false {
			return false
		}
	}
	return true
}

type sudoku struct {
	//1-9 keys, used for faster lookups and searches
	rows, cols, squares map[int][]*cell
	list                []*cell
}

func newCell(n int) *cell {
	c := &cell{}
	c.value = n
	c.possibilities = make(map[int]struct{})

	if c.value == 0 {
		for i := 1; i <= 9; i++ {
			c.addPossibility(i)
		}
	} else {
		c.addPossibility(c.value)
	}

	return c
}

type cell struct {
	possibilities    map[int]struct{} //hashset
	value            int
	square, col, row int
}

func (c *cell) isFixed() bool {
	return c.value > 0
}

func (c *cell) addPossibility(n int) {
	if c.isFixed() {
		return //nothing to do now, is too late
	}

	c.possibilities[n] = struct{}{}
}

func (c *cell) removePossibility(n int) {
	if c.isFixed() {
		return //nothing to do now, is too late
	}

	delete(c.possibilities, n)
	if len(c.possibilities) == 1 {
		//only one value left
		for val := range c.possibilities {
			c.value = val
		}
	}
}

func getResult(n []string) []string {
	s := newSudoku(n)
	var newlyFound []*cell //a queue

	//we first populate with the given values
	for _, c := range s.list {
		if c.isFixed() == false {
			continue
		}
		newlyFound = append(newlyFound, c)
	}

	removeValue := func(value int, l []*cell) (fixed []*cell) {
		for _, p := range l {
			if p.isFixed() {
				continue
			}

			p.removePossibility(value)
			if p.isFixed() {
				fixed = append(fixed, p)
			}
		}
		return
	}

	//while we found at least one new cell fixed value ..continue
	for len(newlyFound) > 0 {
		c := newlyFound[0]
		newlyFound = newlyFound[1:]

		//we remove the c's value from all its peers
		newlyFound = append(newlyFound, removeValue(c.value, s.cols[c.col])...)
		newlyFound = append(newlyFound, removeValue(c.value, s.rows[c.row])...)
		newlyFound = append(newlyFound, removeValue(c.value, s.squares[c.square])...)
	}

	if s.isSolved() == false {
		log.Panic("cannot solve this sudoku this way, make a brute force?")
	}

	var res []string
	for row := 1; row <= 9; row++ {
		var rowAsStrs []string

		for _, c := range s.rows[row] {
			rowAsStrs = append(rowAsStrs, strconv.Itoa(c.value))
		}
		res = append(res, strings.Join(rowAsStrs, ""))
	}
	return res
}

func input() (rows []string) {
	var n string
	var err error

	for l := 9; l > 0; l-- {
		_, err = fmt.Scanln(&n)
		if err != nil {
			log.Panic(err)
		}

		rows = append(rows, n)
	}

	return
}
