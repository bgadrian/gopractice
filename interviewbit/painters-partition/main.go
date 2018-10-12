package painters_partition

import "math"

func sum(a []int) (sum int) {
	for _, v := range a {
		sum += v
	}

	return
}

func best(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

/* This is the brute force/memoization solution, O KN 2
for a better one see the Binary search https://www.youtube.com/watch?v=Ss9ta1zmiZo
https://articles.leetcode.com/the-painters-partition-problem-part-ii/
*/
func paint(K, T int, L []int) int {
	cache := make([]map[int]int, len(L))
	for i := range cache {
		cache[i] = make(map[int]int)
	}

	type rec func(i, k int, f rec) int
	f := func(i, k int, f rec) int {
		if v, ok := cache[i][k]; ok {
			return v
		}
		bestSubSum := math.MaxInt64
		if k == 1 {
			//one painter for all
			bestSubSum = sum(L[i:])
		} else if k >= len(L)-i {
			//best is max of remaining, more painters than boards
			max := L[i]
			for j := i + 1; j < len(L); j++ {
				if L[j] <= max {
					continue
				}
				max = L[j]
			}
			bestSubSum = max
		} else {
			for size := 1; size < len(L)-K-i; size++ {
				painterA := sum(L[i : i+size])
				restPainters := f(i+size, k-1, f)
				bestSubSum = best(bestSubSum, painterA+restPainters)
			}
		}
		//no idea why T exists ...
		cache[i][k] = bestSubSum * T % 10000003
		return cache[i][k]
	}

	return f(0, K, f)
}
