package __7_rotate_matrix

import (
	"fmt"
	"testing"
)

type sts struct {
	in, out [][]int
}

var table = []*sts{
	{
		in: [][]int{
			{10, 11, 12, 13, 14},
			{25, 30, 31, 32, 15},
			{24, 37, 40, 33, 16},
			{23, 36, 35, 34, 17},
			{22, 21, 20, 19, 18},
		},
		out: [][]int{ //rotated 90degrees
			{22, 23, 24, 25, 10},
			{21, 36, 37, 30, 11},
			{20, 35, 40, 31, 12},
			{19, 34, 33, 32, 13},
			{18, 17, 16, 15, 14},
		},
	}, {
		in: [][]int{
			{10, 11},
			{13, 12},
		},
		out: [][]int{ //rotated 90degrees
			{13, 10},
			{12, 11},
		},
	}, {
		in: [][]int{
			{10},
		},
		out: [][]int{ //rotated 90degrees
			{10},
		},
	},
}

func TestSolution(t *testing.T) {
	for _, ts := range table {
		solution(ts.in)
		if comp(ts.in, ts.out) == false {
			t.Errorf("failed for '%v' vs '%v", ts.in, ts.out)
			fmt.Println("Failed for: result vs expected")
			for i := range ts.in {
				fmt.Printf("%v %v\n", ts.in[i], ts.out[i])
			}
		}
	}
}

func comp(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for row := 0; row < len(a); row++ {
		if len(a[row]) != len(b[row]) {
			return false
		}

		for col := 0; col < len(a[row]); col++ {
			if a[row][col] != b[row][col] {
				return false
			}
		}
	}
	return true
}
