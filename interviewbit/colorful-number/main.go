package colorful_number

func colorful(no int) int {
	var digits []int
	for no > 0 {
		digits = append(digits, no%10)
		no /= 10
	}

	unique := make(map[int]struct{}, len(digits))
	run := func(products []int) bool {
		for _, product := range products {
			if _, exists := unique[product]; exists {
				return false
			}
			unique[product] = struct{}{}
		}
		return true
	}

	products := make([]int, len(digits))
	copy(products, digits)
	for size := 2; size < len(digits); size++ {
		if run(products) == false {
			return 0
		}
		for startIndex := 0; startIndex < size; startIndex++ {
			products[startIndex] *= digits[startIndex+size-1]
		}
		products = products[:len(products)-1]
	}
	return 1
}
