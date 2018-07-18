package knightChess

/* Knight movement on a chess board

Given any source point and destination point on a chess board, we need to find whether Knight can move to the destination or not.

Knight's movements on a chess board

The above figure details the movements for a knight ( 8 possibilities ). Note that a knight cannot go out of the board.

If yes, then what would be the minimum number of steps for the knight to move to the said point.
If knight can not move from the source point to the destination point, then return -1

Input:

N, M, x1, y1, x2, y2
where N and M are size of chess board
x1, y1  coordinates of source point
x2, y2  coordinates of destination point
Output:

return Minimum moves or -1
Example

Input : 8 8 1 1 8 8
Output :  6
https://www.interviewbit.com/problems/knight-on-chess-board/
*/
type Node struct {
	row, col int
	jumps    []*Node //connections
}

type Graph struct {
	//we keep it in a matrix, so we can process it
	//0:0 is top left !!!!!
	nodes      [][]*Node
	start, end *Node
}

var newGraph = func(input [6]int) *Graph {
	rows, cols := input[0], input[1]
	//board input is in 1 based, slices are 0 based
	startCol, startRow := input[2]-1, input[3]-1
	endCol, endRow := input[4]-1, input[5]-1

	//generate all the cells
	g := &Graph{}
	g.nodes = make([][]*Node, rows)
	for row := 0; row < rows; row++ {
		g.nodes[row] = make([]*Node, cols)
		for col := 0; col < cols; col++ {
			g.nodes[row][col] = &Node{row: row, col: col}
		}
	}

	g.start = g.nodes[startRow][startCol]
	g.end = g.nodes[endRow][endCol]

	//check for out of bounds
	isValidPos := func(row, col int) bool {
		return row >= 0 && row < rows &&
			col >= 0 && col < cols
	}

	//all possible relative knight jumps
	knightPos := func(row, col int) [][2]int {
		//because the rules of chess never change
		//we can hard code them, more optimal and easy to
		//understand/debug
		return [][2]int{
			//go top and left/right
			{row - 2, col - 1},
			{row - 2, col + 11},
			//go right and top/bottom
			{row - 1, col + 2},
			{row + 1, col + 2},
			//go bottom and left/right
			{row + 2, col - 1},
			{row + 2, col + 1},
			//go left and top/bottom
			{row - 1, col - 2},
			{row + 1, col - 2},
		}
	}

	//add links
	for _, row := range g.nodes {
		for _, from := range row {
			//knight jumps logic + out of bounds/board
			for _, pos := range knightPos(from.row, from.col) {
				if isValidPos(pos[0], pos[1]) == false {
					continue
				}

				//undirected graph, connect 2 ways
				a := g.nodes[from.row][from.col]
				b := g.nodes[pos[0]][pos[1]]

				a.jumps = append(a.jumps, b)
				b.jumps = append(b.jumps, a)
			}
		}
	}

	return g
}

/**
	First we build a graph, then we can ran queries on it
From/to move the knight. The search is done
starting from the start and end position, in paralel,
2 BFSearches until we find a common cell.
node = cell
edge = one knight move
*/
func solution(input [6]int) int {

	g := newGraph(input)

	//we start to search from both points BFS
	//consider them epicenters, when the "circles"
	//intersect, it meas we found the shortest path

	if g.start == g.end {
		return 0
	}

	cp := func(src []*Node) []*Node {
		c := make([]*Node, len(src))
		copy(c, src)
		return c
	}

	//remember what nodes we found, and at what distance/steps
	aVisited := make(map[*Node]int)
	bVisited := make(map[*Node]int)

	var aCurr, bCurr []*Node
	aNext := cp(g.start.jumps)
	bNext := cp(g.end.jumps)

	var commonNode *Node

	step := 0

	for len(aNext) > 0 && len(bNext) > 0 {
		step++
		aCurr, bCurr = cp(aNext), cp(bNext)
		aNext, bNext = make([]*Node, 0), make([]*Node, 0)

		for _, n := range aCurr {
			aVisited[n] = step

			if _, isInB := bVisited[n]; isInB {
				commonNode = n
				break
			}

			for _, jump := range n.jumps {
				if _, alreadyVisited := aVisited[jump]; alreadyVisited {
					continue //avoid infinite loops
				}
				aNext = append(aNext, jump)
			}
		}

		for _, n := range bCurr {
			bVisited[n] = step

			if _, isInA := aVisited[n]; isInA {
				commonNode = n
				break
			}
			for _, jump := range n.jumps {
				if _, alreadyVisited := bVisited[jump]; alreadyVisited {
					continue //avoid infinite loops
				}
				bNext = append(bNext, jump)
			}
		}
	}

	if commonNode == nil {
		return -1
	}

	return aVisited[commonNode] + bVisited[commonNode]
}
