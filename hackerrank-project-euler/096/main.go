package main

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

//https://www.hackerrank.com/contests/projecteuler/challenges/euler096
// highly UNOptimized SUDOKU solved with Constraint Propagation and backtracking,
// solve only easy/medium
//to solve this Problem a better aproach is needed like Exact Cover algorithm,
// the algorithm https://github.com/ShivanKaul/Sudoku-DLX and https://gist.github.com/Eibwen/e8fec52e0f5927108d86

func main() {
	for _, strRow := range getResult(input()) {
		fmt.Println(strRow)
	}
}

func newSudokuFromStdin(lines []string) *sudoku {
	var list []*cell

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

			list = append(list, c)
		}
	}

	return newSudokuFromCells(list)
}

func newSudokuFromCells(cells []*cell) *sudoku {
	s := &sudoku{}
	s.emptyCount = 9 * 9
	s.emptyCount = 9 * 9

	for _, orig := range cells {
		c := newCell(orig.value) //clone cell
		c.col = orig.col
		c.row = orig.row
		c.square = orig.square

		if c.isSolved() {
			s.emptyCount--
		}

		s.list = append(s.list, c)
	}

	return s
}

func (s *sudoku) isSolved() bool {
	return s.emptyCount <= 0
}

func (s *sudoku) getCell(row, col int) *cell {
	//TODO optimize
	for _, c := range s.list {
		if c.row == row && c.col == col {
			return c
		}
	}

	return nil
}

func (s *sudoku) getAsStrings() []string {
	res := make([]string, 9)

	for row := 1; row <= 9; row++ {
		var rowAsStrs []string

		for _, c := range s.list {
			if c.row != row {
				continue
			}

			rowAsStrs = append(rowAsStrs, strconv.Itoa(c.value))
		}
		res[row-1] = strings.Join(rowAsStrs, "")
	}
	return res
}

func (s *sudoku) getUnsolvedCells() (result []*cell) {
	for _, c := range s.list {
		if c.isSolved() {
			continue
		}
		result = append(result, c)
	}
	return
}

func (s *sudoku) getSolvedCells() (result []*cell) {
	for _, c := range s.list {
		if c.isSolved() == false {
			continue
		}
		result = append(result, c)
	}
	return
}

func (s *sudoku) applyConstraints() {
	//while we found at least one new cell fixed value ..continue
	newlyFound := s.getSolvedCells()
	for len(newlyFound) > 0 {
		c := newlyFound[0]
		newlyFound = newlyFound[1:]

		for _, unsol := range s.getUnsolvedCells() {
			if unsol.row != c.row || unsol.col != c.col || unsol.square != c.square {
				continue
			}

			unsol.removePossibility(c.value)
			if unsol.isSolved() {
				s.emptyCount--
				newlyFound = append(newlyFound, unsol)
			}
		}
	}
}

type sudoku struct {
	list       []*cell
	emptyCount int
}

func newCell(n int) *cell {
	c := &cell{}
	c.possibilities = make(map[int]struct{})

	if n == 0 {
		for i := 1; i <= 9; i++ {
			c.addPossibility(i)
		}
	} else {
		c.addPossibility(c.value)
	}
	c.value = n

	return c
}

type cell struct {
	possibilities    map[int]struct{} //hashset
	value            int
	square, col, row int
}

func (c *cell) isSolved() bool {
	return c.value > 0
}

func (c *cell) addPossibility(n int) {
	if c.isSolved() {
		log.Panicf("someone tried to ressurect a solved cell %v", c)
		return //nothing to do now, is too late
	}

	c.possibilities[n] = struct{}{}
}

func (c *cell) removePossibility(n int) {
	if c.isSolved() {
		log.Panicf("someone tried to remove a poss from  a solved cell %v", c)
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
	var s *sudoku

	//quick exit for easy difficulty
	s = newSudokuFromStdin(n)
	s.applyConstraints()
	//
	if s.isSolved() {
		return s.getAsStrings()
	}

	antiInfinite := 100000

	//try all the possible left solutions
	type tryAll func(s *sudoku, this tryAll) *sudoku

	//returns Nil if failed or a Solved instance
	rec := func(s *sudoku, this tryAll) *sudoku {
		antiInfinite--

		if antiInfinite <= 0 {
			log.Panic("anti infinite threshold hit")
			return nil
		}

		//we take all the unsolved cells, and all their posibilities
		//we select one and recursive try again
		unsolvedCells := s.getUnsolvedCells()
		sort.Slice(unsolvedCells, func(i, j int) bool {
			return len(unsolvedCells[i].possibilities) < len(unsolvedCells[j].possibilities)
		})

		for _, pivotCell := range unsolvedCells {
			var pivotPossibilites []int
			for poss := range pivotCell.possibilities {
				pivotPossibilites = append(pivotPossibilites, poss)
			}
			//sort.Ints(pivotPossibilites) //remove the randomness, for debugging, because map-range

			for _, poss := range pivotPossibilites {
				sCloned := newSudokuFromCells(s.list)
				clonedPivot := sCloned.getCell(pivotCell.row, pivotCell.col)

				//double check
				if clonedPivot.row != pivotCell.row || clonedPivot.col != pivotCell.col {
					log.Panicf("your logic is flawed, cloned pivot, %v vs %v", clonedPivot, pivotCell)
				}
				clonedPivot.value = poss

				sCloned.applyConstraints()

				if sCloned.isSolved() {
					return sCloned
				}

				//fmt.Printf("set value of [%v][%v]=%v from %v\n",
				//	pivotCell.row, pivotCell.col, poss, pivotPossibilites)

				//try recursive, this means is another pivot down the road
				sCloned = this(sCloned, this)
				if sCloned != nil { //means is solved
					return sCloned
				}
			}
		}
		return nil
	}

	s = rec(s, rec)

	if s == nil || s.isSolved() == false {
		log.Panic("cannot solve this sudoku, there is a flaw in this algorithm")
	}

	return s.getAsStrings()
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
