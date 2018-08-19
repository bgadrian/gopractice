package dungeon_princess

/* The demons had captured the princess (P) and imprisoned her in the bottom-right corner of a dungeon. The dungeon consists of M x N rooms laid out in a 2D grid. Our valiant knight (K) was initially positioned in the top-left room and must fight his way through the dungeon to rescue the princess.

The knight has an initial health point represented by a positive integer. If at any point his health point drops to 0 or below, he dies immediately.

Some of the rooms are guarded by demons, so the knight loses health (negative integers) upon entering these rooms; other rooms are either empty (0's) or contain magic orbs that increase the knight's health (positive integers).

In order to reach the princess as quickly as possible, the knight decides to move only rightward or downward in each step.



Write a function to determine the knight's minimum initial health so that he is able to rescue the princess.

For example, given the dungeon below, the initial health of the knight must be at least 7 if he follows the optimal path RIGHT-> RIGHT -> DOWN -> DOWN.

-2 (K)	-3	3
-5	-10	1
10	30	-5 (P)


Note:

The knight's health has no upper bound.
Any room can contain threats or power-ups, even the first room the knight enters and the bottom-right room where the princess is imprisoned.
https://leetcode.com/problems/dungeon-game/description/
*/

type pos struct {
	row, col int
}

func solution(grid [][]int) int {
	//TODO add defensive validation rules

	//zero based matrix/grid
	//startPos := pos{0, 0}
	rows := len(grid)
	cols := len(grid[0])
	//target := pos{rows - 1, cols - 1}
	hp := make([][]int, rows)

	for row := range hp {
		hp[row] = make([]int, cols)
	}
	/*
		Because there are constraints, we can apply a
		greedy aproach and avoid a DP aproach/brute force.

		We start from the end position, and go backwards
		only top and left. We calculate the minim amount of
		HP hp to reached that position, and choose the
		best solution.
		The best solution is:
		* prevValue (min out of right/bottom cells)
		* minimum for this pos Max(1, -value+1 - prevValue)

				2
				5
		1	1	6
	*/

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	//bottom right cell, first one to calculate
	hp[rows-1][cols-1] = max(1, 1-grid[rows-1][cols-1])

	//the right column and bottom row
	for row := rows - 2; row >= 0; row-- {
		prev := hp[row+1][cols-1]
		curr := grid[row][cols-1]
		hp[row][cols-1] = max(1, prev-curr)
	}

	for col := cols - 2; col >= 0; col-- {
		prev := hp[rows-1][col+1]
		curr := grid[rows-1][col]
		hp[rows-1][col] = max(1, prev-curr)
	}

	//now we can calc the rest of the cells
	for row := rows - 2; row >= 0; row-- {
		for col := cols - 2; col >= 0; col-- {
			curr := grid[row][col]
			right := max(hp[row][col+1]-curr, 1)
			bottom := max(hp[row+1][col]-curr, 1)

			hp[row][col] = min(right, bottom)
		}
	}
	return hp[0][0]
}
