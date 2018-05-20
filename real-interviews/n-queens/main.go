package n_queens

import "fmt"

/*
Given an N sized chess board, place N queens
so it doesn't threat one another.
*/

type pos struct {
	row, col int
}

func (p *pos) String() string {
	return fmt.Sprintf("[row:%v,col:%v]", p.row, p.col)
}

type cell struct {
	p     *pos
	queen bool //has a queen on it or not
}

/*
We use backtracking, procedural and recursion code,
*/
func procedural(n int) []*pos {
	//for O(1) lookups on rows
	//rows := make(map[int]struct{})
	board := make([][]*cell, n)
	var queens []*pos

	for row := 0; row < n; row++ {
		board[row] = make([]*cell, n)
		for col := 0; col < n; col++ {
			board[row][col] = &cell{
				p: &pos{row, col},
			}
		}
	}

	//hasQueen := func(p *pos) bool {
	//	return board[p.row][p.col].queen
	//}
	//
	//validPos := func(p *pos) bool {
	//	if p.row < 0 || p.row >= n {
	//		return false
	//	}
	//	if p.col < 0 || p.col >= n {
	//		return false
	//	}
	//	return true
	//}

	//isPosThreaten := func(origin *pos) bool {
	//	_, ok := rows[origin.row]
	//	if ok {
	//		return true
	//	}
	//
	//	topLeft, bottomLeft := &pos{}, &pos{}
	//
	//	//only search on diagonal positions
	//	for offset := 1; offset < origin.col; offset++ {
	//		topLeft.row = origin.row - offset
	//		topLeft.col = origin.col - offset
	//
	//		if validPos(topLeft) && hasQueen(topLeft) {
	//			return true
	//		}
	//
	//		bottomLeft.row = origin.row + offset
	//		bottomLeft.col = topLeft.col
	//
	//		if validPos(bottomLeft) && hasQueen(bottomLeft) {
	//			return true
	//		}
	//
	//		return false
	//	}
	//}

	isPosThreaten := func(origin *pos) bool {
		for _, q := range queens {
			if q.row == origin.row {
				return true
			}

			diffRow := origin.row - q.row
			diffCol := origin.col - q.col

			if diffRow == diffCol || diffRow == -diffCol {
				return true
			}

		}
		return false
	}

	type rec func(col int, r rec) bool

	search := func(col int, r rec) bool {
		if col >= n {
			return false //reached the end
		}

		for row := 0; row < n; row++ {
			if isPosThreaten(&pos{row, col}) {
				continue
			}
			//we try to put it here
			board[row][col].queen = true
			//rows[row] = struct{}{}
			queens = append(queens, &pos{row, col})

			if col == n-1 {
				return true // we are done
			}

			nextWasOk := r(col+1, r)
			if nextWasOk {
				return true
			}

			//rollback
			board[row][col].queen = false
			//delete(rows, row)
			queens = queens[:len(queens)-1]
		}

		return false
	}

	if search(0, search) {
		//return the positions
		return queens
	}

	return nil
}
