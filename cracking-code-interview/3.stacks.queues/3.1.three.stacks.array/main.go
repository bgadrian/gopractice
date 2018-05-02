package __1_three_stacks_array

import (
	"errors"
	"fmt"
	"log"

	"engo.io/engo/math"
)

/*
3.1
Three in One: Describe how you could use a single array to implement three MultiStacks.
*/

type StackInfo struct {
	start, size, capacity int
	parent                *MultiStacks //ugly solution
}

func (si *StackInfo) lastElementIndex() int {
	return si.parent.adjustIndex(si.start + si.size - 1) //start 0, 3, => 2, 0+3-1
}

func (si *StackInfo) isFull() bool {
	return si.size >= si.capacity
}

func (si *StackInfo) isEmpty() bool {
	return si.size <= 0
}

/* check if the index of parent.a belongs to this stack*/
func (si *StackInfo) isWithinStackCapacity(parentIndex int) bool {
	if parentIndex < 0 || parentIndex >= len(si.parent.a) {
		//parentIndex = si.parent.adjustIndex(parentIndex)
		return false
	}

	iAmcircular := si.lastElementIndex() < si.start
	if iAmcircular {
		//is at the end
		if parentIndex >= si.start {
			return true
		}

		//is at the start
		sizeAtEnd := len(si.parent.a) - si.start - 1
		sizeAtBegining := si.capacity - sizeAtEnd
		if parentIndex < (sizeAtBegining) {
			return true
		}

		return false
	}

	return parentIndex >= si.start && parentIndex <= si.start+si.capacity-1
}

type MultiStacks struct {
	a     []int
	sInfo []*StackInfo
}

func newStacks(stackCount, totalCapacity int) *MultiStacks {
	s := &MultiStacks{}

	if stackCount > totalCapacity {
		log.Panic("cannot have more MultiStacks than cap")
	}

	defaultCap := int(totalCapacity / stackCount)
	s.a = make([]int, totalCapacity)
	for i := 0; i < stackCount; i++ {
		s.sInfo = append(s.sInfo, &StackInfo{
			start:    i * defaultCap,
			capacity: defaultCap,
			parent:   s,
		})
		//capacity = 3
		//stack 0 => [0,1,2], stack 1 => [3,4,5]
	}

	leftOvers := totalCapacity - defaultCap*stackCount
	if leftOvers > 0 {
		s.sInfo[len(s.sInfo)-1].capacity += leftOvers
	}

	return s
}

func (s *MultiStacks) push(stackIndex, value int) error {
	if s.allStacksAreFull() {
		return errors.New("all stacks are full, go away")
	}

	si := s.sInfo[stackIndex]

	if si.isFull() {
		s.expand(stackIndex)
	}
	si.size++
	lastIndex := si.lastElementIndex()
	s.a[lastIndex] = value

	return nil
}

func (s *MultiStacks) expand(stackIndex int) {
	si := s.sInfo[stackIndex]
	siRightIndex, siRight := s.getNextStackInfo(stackIndex)

	if siRight.isFull() {
		s.shift(siRightIndex)
	} else {
		s.shrink(siRightIndex)
	}

	si.capacity++
}

func (s *MultiStacks) getNextStackInfo(stackIndex int) (int, *StackInfo) {
	stackCount := len(s.sInfo)
	siRightIndex := int((stackIndex + 1) % stackCount)
	return siRightIndex, s.sInfo[siRightIndex]
}

func (s *MultiStacks) shift(stackIndex int) {
	si := s.sInfo[stackIndex]
	siRightIndex, siRight := s.getNextStackInfo(stackIndex)

	if siRight.isFull() {
		s.shift(siRightIndex)
	} else {
		s.shrink(siRightIndex)
	}

	//MOVE all to the right
	s.moveAllItemsToTheRight(si)
	si.start = s.adjustIndex(si.start + 1)
}

func (s *MultiStacks) moveAllItemsToTheRight(si *StackInfo) {
	for i := si.start + si.size; i > si.start; i-- {
		i = s.adjustIndex(i) //for overflow
		leftIndex := s.adjustIndex(i - 1)

		s.a[i] = s.a[leftIndex]
	}
}

func (s *MultiStacks) shrink(stackIndex int) {
	si := s.sInfo[stackIndex]
	if si.isFull() {
		log.Panicf("stack '%v' is full, cannot be shrinked", stackIndex)
	}

	if si.isEmpty() == false {
		s.moveAllItemsToTheRight(si)
	}

	s.a[si.start] = math.MinInt32 //for debugging
	si.start = s.adjustIndex(si.start + 1)
	si.capacity--
}

func (s *MultiStacks) allStacksAreFull() bool {
	for _, si := range s.sInfo {
		if si.isFull() == false {
			return false
		}
	}
	return true
}

func (s *MultiStacks) pop(stackIndex int) (int, error) {
	si := s.sInfo[stackIndex]
	if si.isEmpty() {
		return 0, fmt.Errorf("stack '%v' is empty", stackIndex)
	}

	popIndex := si.lastElementIndex()
	r := s.a[popIndex]
	s.a[popIndex] = math.MinInt32
	si.size--

	return r, nil
}

func (s *MultiStacks) isFull(stackIndex int) bool {
	return s.allStacksAreFull()
}

func (s *MultiStacks) adjustIndex(i int) int {
	if i < 0 {
		log.Panic("TODO")
	}
	max := len(s.a)
	if i >= max {
		i = ((i % max) + max) % max
	}
	return i
}
