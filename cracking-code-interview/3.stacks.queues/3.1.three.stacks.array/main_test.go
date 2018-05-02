package __1_three_stacks_array

import (
	"math/rand"
	"testing"
)

func TestStacksOneInTen(t *testing.T) {
	s := newStacks(1, 10)
	values := getRandoms(10)
	err := testPush(s, 0, values)
	if err != nil {
		t.Error(err)
		return
	}

	outValues, err := testPop(s, 0, 10)
	if err != nil {
		t.Error(err)
		return
	}

	values = reverseSliceInt(values)

	if compareSliceInts(values, outValues) == false {
		t.Errorf("failed , exp '%v' got '%v' ", values, outValues)
	}
}

func TestMultiple(t *testing.T) {
	count, cap := 0, 1
	table := [][]int{
		//stack counts, total capacity
		{1, 10},
		{2, 10},
		{3, 12},
		{5, 10}, //values:0, 2, 4 , 4, 0
	}

	for _, tt := range table {
		countPerStack := make([]int, tt[count])
		valuesCache := make([][]int, tt[count])

		//perStackCount := int(math.Floor(float64(tt[cap]) / float64(tt[count])))
		//distribute uneven count of values
		capLeft := tt[cap]
		for i := 0; i < tt[count]; i++ {
			tmp := i * 2
			if tmp > capLeft {
				tmp = capLeft
			}
			countPerStack[i] = tmp
			capLeft -= tmp
		}

		s := newStacks(tt[count], tt[cap])
		for i := 0; i < tt[count]; i++ {
			valuesCache[i] = getRandoms(countPerStack[i])
			testPush(s, i, valuesCache[i])
			valuesCache[i] = reverseSliceInt(valuesCache[i])
		}

		//pop and compare
		for i := 0; i < tt[count]; i++ {
			res, err := testPop(s, i, countPerStack[i])
			if err != nil {
				t.Error(err)
				return
			}
			if compareSliceInts(res, valuesCache[i]) == false {
				t.Errorf("for test '%v'\n exp '%v'\n got '%v'",
					tt, valuesCache[i], res)
				return
			}
		}
	}
}

func testPop(s *MultiStacks, stackIndex, count int) ([]int, error) {
	var r []int

	for i := 0; i < count; i++ {
		v, err := s.pop(stackIndex)
		if err != nil {
			return nil, err
		}
		r = append(r, v)
	}

	return r, nil
}

func testPush(s *MultiStacks, stackIndex int, values []int) error {
	for _, v := range values {
		err := s.push(stackIndex, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func getRandoms(count int) []int {
	r := make([]int, count)

	for i := 0; i < count; i++ {
		r[i] = rand.Int()
	}

	return r
}

func compareSliceInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func reverseSliceInt(a []int) []int {
	r := make([]int, len(a))

	for i, v := range a {
		r[len(a)-1-i] = v
	}

	return r
}
