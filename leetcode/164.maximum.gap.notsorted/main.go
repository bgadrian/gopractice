package _64_maximum_gap_notsorted

import "math"

/* Given an unsorted array, find the maximum difference between the successive elements in its sorted form.

Return 0 if the array contains less than 2 elements.

Example 1:

Input: [3,6,9,1]
Output: 3
Explanation: The sorted form of the array is [1,3,6,9], either
             (3,6) or (6,9) has the maximum difference 3.
Example 2:

Input: [10]
Output: 0
Explanation: The array contains less than 2 elements, therefore return 0.
Note:

You may assume all elements in the array are non-negative integers and fit in the 32-bit signed integer range.
Try to solve it in linear time/space.
*/

var getMin = func(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
var getMax = func(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
var getAbs = func(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
var getNextValidBucket = func(startI int, buckets []*bucket) (*bucket, int) {
	for startI < len(buckets)-1 {
		b := buckets[startI]
		if b.hasValue {
			return b, startI
		}
		startI++
	}
	return nil, -1
}

type bucket struct {
	min, max int
	hasValue bool
}

func (b *bucket) add(v int) {
	b.min = getMin(b.min, v)
	b.max = getMax(b.max, v)
	b.hasValue = true
}
func newBucket() *bucket {
	b := &bucket{}
	b.min = math.MaxInt64
	return b
}

func solution(a []int) int {

	/*
		the trick is to make it in linear time
		- all numbers are positive
		- array is shuffled

	*/

	//edge cases and sanitize
	le := len(a)
	if le < 2 {
		return -1
	}
	if le == 2 {
		return getAbs(a[0] - a[1])
	}
	//check of negative values?!

	var gap, diff int
	min := math.MaxInt64
	max := math.MinInt64

	for _, v := range a {
		min = getMin(min, v)
		max = getMax(max, v)
	}
	gap = int((max-min)/le - 1)
	diff = max - min

	countBucket := int(math.Ceil(float64(diff) / float64(gap)))
	buckets := make([]*bucket, countBucket)
	for i := range buckets {
		buckets[i] = newBucket()
	}

	getBucketIdFor := func(v int) int { return int(math.Floor(float64(v-min) / float64(gap))) }

	//add the values to the bucket
	for _, v := range a {
		idFor := getBucketIdFor(v)
		buckets[idFor].add(v)
	}

	//find the largest gap
	result := math.MinInt64

	prev, bi := getNextValidBucket(0, buckets)
	var next *bucket
	next, bi = getNextValidBucket(bi+1, buckets)

	for next != nil {
		result = getMax(result, next.min-prev.max)
		prev = next
		next, bi = getNextValidBucket(bi+1, buckets)
	}

	//return the result
	return result
}
