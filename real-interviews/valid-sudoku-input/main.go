package valid_sudoku_input

type set map[rune]struct{}

func has(s set, v rune) bool {
	_, ok := s[v]
	return ok
}

func add(s set, v rune) {
	s[v] = struct{}{}
}

func validValue(r rune) bool {
	return r >= '1' && r <= '9'
}

func solution(A []string) int {

	//no need to convert to numbers, we just need distinct values
	//on each row, col and square
	//O(n) time (81 cells) and O(n) space (2 x 81 + 9 Runes in memory)

	squareVals := make([]set, 9)
	colVals := make([]set, 9)

	//nil maps
	for i := range squareVals {
		squareVals[i] = make(set)
		colVals[i] = make(set)
	}

	if len(A) != 9 {
		return 0
	}

	for row, colStr := range A {
		rowVals := make(set)
		if len(colStr) != 9 {
			return 0
		}

		for col, v := range colStr {
			if v == '.' {
				continue
			}

			if validValue(v) == false {
				return 0
			}

			if has(rowVals, v) {
				return 0
			}
			add(rowVals, v)

			if has(colVals[col], v) {
				return 0
			}
			add(colVals[col], v)

			square := row/3*3 + col/3
			if has(squareVals[square], v) {
				return 0
			}
			add(squareVals[square], v)
		}
	}
	return 1
}
