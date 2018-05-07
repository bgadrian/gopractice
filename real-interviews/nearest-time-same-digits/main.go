package nearest_time_same_digits

import (
	"strconv"
	"strings"
)

/*
you get a 24h time digits
return the nearest time that is formed from the same digits
*/

func intToArr(t string) (r []int) {
	for _, v := range t {
		ts, _ := strconv.Atoi(string(v))
		r = append(r, ts)
	}
	return
}

func arrToInt(a []int) string {
	rs := make([]string, len(a))
	for i, v := range a {
		rs[i] = strconv.Itoa(v)
	}
	return strings.Join(rs, "")
}

func isValid(a []int) bool {
	h := a[0]*10 + a[1]
	if h > 23 {
		return false
	}
	m := a[2]*10 + a[3]
	if m > 59 {
		return false
	}

	s := a[4]*10 + a[5]
	if s > 59 {
		return false
	}
	return true
}
func solution(t string) string {
	a := intToArr(t)
	//we start from right to left
	for iDest := len(a) - 2; iDest >= 0; iDest-- {
		for iSource := len(a) - 1; iSource >= iDest+1; iSource-- {
			if a[iDest] == a[iSource] {
				continue
			}
			a[iDest], a[iSource] = a[iSource], a[iDest]
			if isValid(a) {
				right := a[iDest:]
				if a[iDest] > a[iSource] {
					//try to sort in ASC order
					//move a[iSource] as right as possible
				} else {
					//try to srt in DESC order
					//move a[iSource] as left as poss, max a[iDest]
				}
				return arrToInt(a)
			}
			//rollback
			a[iDest], a[iSource] = a[iSource], a[iDest]
		}
	}

	return t
}
