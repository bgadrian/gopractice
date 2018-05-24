package simulate_die_with_smaller_one

import (
	"testing"

	"github.com/engoengine/math"
)

func TestAll100(t *testing.T) {
	for bigDie := 3; bigDie <= 100; bigDie++ {
		for smallDie := 2; smallDie < bigDie; smallDie++ {
			buckets := make([]int, bigDie)
			for tests := 0; tests < 10000; tests++ {
				value := get(bigDie, smallDie)
				if value < 1 || value > bigDie {
					t.Errorf("value out of bounds, [%v,%v] got: %v",
						1, bigDie, value)
				}
				buckets[value-1]++
			}

			if isEvenlyDistributed(buckets, 30) == false {
				t.Errorf("the random is not so random '%v'",
					buckets)
			}
		}
	}
}

//Testing a random generator is hard work
//we just do a simple check if the values
//are uniformly distributed in buckets
//with a maxPercDiff margin of error.
//because the domain of the values is very small (<100)
//the differences can be pretty big.
func isEvenlyDistributed(buckets []int, maxPercDiff int) bool {
	min, max := math.MaxInt64, math.MinInt64
	for _, v := range buckets {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return (min/max)*100 <= maxPercDiff
}
