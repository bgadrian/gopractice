package robot_circular_path_2D

import "testing"

var positive = [][]rune{
	{},
	{'L'},
	{'R'},
	{'R', 'L', 'L', 'R'},
	{'G', 'G', 'G', 'L', 'L', 'G', 'G', 'G'},
	{'G', 'G', 'G', 'R', 'R', 'G', 'G', 'G'},
	{'G', 'L', 'G', 'R', 'G', 'R', 'G', 'R', 'G', 'G'},
}

//TODO negative tests

func TestNormalBasic(t *testing.T) {
	for _, test := range positive {
		r := normal(&robot{3, 4, NORTH}, test)
		if r == false {
			t.Errorf("failed for '%v'", string(test))
		}
	}
}
