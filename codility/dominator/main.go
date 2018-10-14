package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4

	size := len(A)
	if size == 0 {
		return -1
	}
	if size == 1 {
		return 0
	}

	dom := A[0]
	count := 1

	for i := 1; i < size; i++ {
		val := A[i]

		if val == dom {
			count++
			continue
		}
		count--
		if count <= 0 {
			dom = val
			count = 1
		}
	}

	count = 0
	index := -1
	for i, val := range A {
		if val == dom {
			count++
			index = i
		}
	}

	if count > size/2 {
		return index
	}
	return -1
}
