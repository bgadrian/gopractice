package stepping_numbers

/* Given N and M find all stepping numbers in range N to M

The stepping number:

A number is called as a stepping number if the adjacent digits have a difference of 1.
e.g 123 is stepping number, but 358 is not a stepping number

Example:

N = 10, M = 20
all stepping numbers are 10 , 12
Return the numbers in sorted order.

https://www.interviewbit.com/problems/stepping-numbers/
*/

import "math"
import "strconv"

func solution(n, m int) []int {

	if m < n {
		return []int{}
	}

	var result []int

	if n < 10 {

		for newn := n; newn < 10; newn++ {
			result = append(result, newn)
		}
		n = 10
	}

	digits := func(a int) int {
		return len(strconv.Itoa(a))
	}
	arrToInt := func(a []int) int {
		var res int
		max := len(a) - 1
		for i := 0; i <= max; i++ {
			res += a[max-i] * int(math.Pow(10, float64(i)))
		}
		return res
	}

	type generate func(digits int, sofar []int, f generate)

	g := func(digits int, sofar []int, f generate) {
		if len(sofar) == digits {
			r := arrToInt(sofar)
			if r >= n && r <= m {
				result = append(result, r)
			}
			return
		}

		last := sofar[len(sofar)-1]
		if last > 0 {
			f(digits, append(sofar, last-1), f)
		}
		if last < 9 {
			f(digits, append(sofar, last+1), f)
		}
	}

	ds := digits(n)
	de := digits(m)
	for digits := ds; digits <= de; digits++ {
		for first := 1; first <= 9; first++ {
			g(digits, []int{first}, g)
		}
	}

	return result
}
