package __12_count_paths_with_sum

/*
Paths with Sum: You are given a binary tree in which each node contains an integer value (which
4.12
might be positive or negative) . Design an algorithm to count the number of paths that sum to a
given value. The path does not need to start or end at the root or a leaf, but it must go downwards
(traveling only from parent nodes to child nodes).
*/

type node struct {
	left, right *node
	val         int
}

func solution(root *node, sum int) int {

	pathC := make(map[int]int)

	type recf func(n *node, t, r int, p map[int]int, f recf) int

	rec := func(node *node, targetSum, runningSum int, pathCount map[int]int, f recf) int {
		if node == nil {
			return 0
		}

		runningSum += node.val
		sum := runningSum - targetSum
		pathsAlready, _ := pathCount[sum] //0 is default

		if runningSum == targetSum {
			pathsAlready++
		}

		pathCount[runningSum]++
		pathsAlready += f(node.left, targetSum, runningSum, pathCount, f)
		pathsAlready += f(node.right, targetSum, runningSum, pathCount, f)
		pathCount[runningSum]--

		return pathsAlready
	}
	return rec(root, sum, 0, pathC, rec)
}
