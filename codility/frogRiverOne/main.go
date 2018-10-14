package solution

//frog other side of the river. A[K] = slot, K = the second when a leaf falled on slot
//-1 if frog cannot reach other side
func Solution(X int, A []int) int {
	// write your code in Go 1.4

	if len(A) < X {
		return -1
	}
	var uniqueCount int
	fallen := make(map[int]struct{}, X)

	for time, slot := range A {
		if _, duplicate := fallen[slot]; duplicate {
			continue
		}
		fallen[slot] = struct{}{}
		uniqueCount++
		if uniqueCount >= X {
			return time
		}
	}
	return -1
}
