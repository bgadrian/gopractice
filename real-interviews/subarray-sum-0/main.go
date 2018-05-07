package subarray_sum_0

func solution(a []int) (int, int) {
	if len(a) == 0 {
		return 0, 0
	}
	if a[0] == 0 {
		return 0, 1
	}

	s := map[int]int{} //sum=indexEnd
	cs := 0

	for i := 0; i < len(a); i++ {
		cs += a[i]

		if endIndex, ok := s[cs]; ok {
			return endIndex + 1, i - endIndex
		}
		s[cs] = i
	}
	if endIndex, ok := s[0]; ok {
		return 0, endIndex + 1
	}
	return 0, 0
}
