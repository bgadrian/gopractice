package sum_4

import (
	"fmt"
	"sort"
)

/*
Given an array S of n integers, are there elements a, b, c, and d in S such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

 Note:
Elements in a quadruplet (a,b,c,d) must be in non-descending order. (ie, a ≤ b ≤ c ≤ d)
The solution set must not contain duplicate quadruplets.
Example :
Given array S = {1 0 -1 0 -2 2}, and target = 0
A solution set is:

    (-2, -1, 1, 2)
    (-2,  0, 0, 2)
    (-1,  0, 0, 1)
Also make sure that the solution set is lexicographically sorted.
Solution[i] < Solution[j] iff Solution[i][0] < Solution[j][0] OR (Solution[i][0] == Solution[j][0] AND ... Solution[i][k] < Solution[j][k])
*/
func solution(A []int, target int) (result [][]int) {
	s := A
	//sort.Ints(s)
	ls := len(s)
	if ls < 4 {
		return //panic? cannot make quadruplet out of <4 items
	}

	resSet := make(map[string]struct{})

	cache := make(map[int][][2]int) //sum => list of indexs pair that make that sum
	for cIndex := 0; cIndex < ls-1; cIndex++ {
		c := s[cIndex]
		//if cIndex > 1 && s[cIndex-1] == c {
		//	continue
		//}
		for dIndex := cIndex + 1; dIndex < ls; dIndex++ {
			d := s[dIndex]
			//if s[dIndex-1] == d {
			//	continue
			//}
			sum := c + d
			ABSum := target - sum

			if ABPairs, ok := cache[ABSum]; ok {
				//at least 1 pair adds up to the difference
				for _, abPair := range ABPairs {
					//because s is sorted, a,b,c,d will be sorted, like they should
					a, b := abPair[0], abPair[1]
					if a == c || a == d || b == c || b == d {
						continue
					}
					tmp := []int{a, b, c, d}
					//tmp := []int{abPair[0], abPair[1], c, d}
					sort.Ints(tmp)

					//we count the unique results using a toString hack
					asStr := fmt.Sprint(tmp)
					if _, ok := resSet[asStr]; ok {
						continue
					}
					result = append(result, tmp)

					//add the slice to result
					//result = append(result, tmp)
					resSet[asStr] = struct{}{}
				}
			}

			cache[sum] = append(cache[sum], [2]int{c, d})
		}
	}

	return
}
