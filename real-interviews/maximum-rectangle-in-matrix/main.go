package maximum_rectangle_in_matrix

import "container/list"

func maxOnesAreaMatrix(a [][]int) int {
	if len(a) == 0 {
		return 0
	}

	tmp := make([]int, len(a[0]))
	maxArea := 0

	for _, vl := range a {
		for i, v := range vl {
			if v == 0 {
				tmp[i] = 0
			}
			//we presume that all values are 0 or 1
			tmp[i] += v
		}
		area := maxHistogramArea(tmp)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

/*Calculate the greatest rectangle/area from a histogram.
We only add/remove all elements once so is O(n).
Code is written in procedural style.
*/
func maxHistogramArea(histo []int) int {
	s := list.New() //the "stack", using a simple linked list
	maxArea := 0

	//helper, for duplicate code.
	pop := func(width int) {
		top := s.Front()
		s.Remove(top)

		if s.Len() > 0 {
			width = (width - 1) - s.Front().Value.(int)
		}
		height := histo[top.Value.(int)]

		area := height * width
		if area > maxArea {
			maxArea = area
		}
	}

	for i, v := range histo {
		for s.Len() > 0 && histo[s.Front().Value.(int)] > v {
			pop(i)
		}

		s.PushFront(i)
	}

	for s.Len() > 0 {
		pop(len(histo))
	}

	return maxArea
}
