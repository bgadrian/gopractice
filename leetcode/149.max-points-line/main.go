package _49_max_points_line

/* Given n points on a 2D plane, find the maximum number of points that lie on the same straight line.

Sample Input :

(1, 1)
(2, 2)
Sample Output :

2
https://www.interviewbit.com/problems/points-on-the-straight-line/
https://leetcode.com/problems/max-points-on-a-line/description/
*/
type Point struct {
	x, y int
}

func (p1 Point) eq(p2 Point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

func (p1 Point) slope(p2 Point) float64 {
	return float64(p2.y-p1.y) / float64(p2.x-p1.x)
}

// O(n squared) time and O(n) space.
// isn't safe for 0 slopes. Can be optimied further
// in checking for duplicate points, and memoize slopes
func solution(points []Point) int {

	le := len(points)
	var result int

	for i := 0; i < le-2; i++ {
		p1 := points[i]
		slopes := make(map[float64]int) //slope =>count
		for j := i + 1; j < le; j++ {
			p2 := points[j]
			slopes[p1.slope(p2)]++
		}

		for _, count := range slopes {
			count++ //because the first point, pivot, p1
			//wasn't counted anywhere
			if count < result {
				continue
			}
			result = count
		}
	}
	return result
}
