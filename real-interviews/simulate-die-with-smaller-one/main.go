package simulate_die_with_smaller_one

import (
	"math"
	"math/rand"
)

/*
 Simulate a die with a smaller one, example:
you have a 6 sided die, you have to simulate a 12 sided
die (with all faces evenly chances to hit).

There are many ways to solve it, I implemented the simplest one.
https://www.quora.com/How-can-one-generate-a-random-number-uniformly-distributed-on-1-2-dots-7-with-only-a-6-faced-die
https://www.quora.com/How-do-I-use-a-fair-6-sided-die-to-generate-a-random-number-from-1-to-20-with-each-outcome-occurring-with-the-same-probability?share=1

*/
func get(big, small int) int {
	rolls := int(math.Ceil(float64(big) / (float64(small))))
	var sum int
	for sum <= big {
		for r := 0; r < rolls; r++ {
			sum += getFace(small)
		}
	}

	return sum%big + 1
}

//rand between 1 and die (inclusive)
func getFace(die int) int {
	return rand.Int()%die + 1
}
