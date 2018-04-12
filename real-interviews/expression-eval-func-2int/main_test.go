package vald

import "testing"

var tableBasic = map[string]int{
	"add(2,3)":            5,
	"sub(3,2)":            1,
	"mult(2,3)":           6,
	"add(222000222000,1)": 222000222001,
	"add( 1 , 2 )":        3,
	"add( -1 , 200 )":     199,
}

func TestParseBasic(t *testing.T) {
	for str, res := range tableBasic {
		o := Parse(str)
		if o != res {
			t.Errorf("exp %v got %v, for '%v'",
				res, o, str)
		}
	}
}

var tableNested = map[string]int{
	"add(add(1,3) , 5)":        9,
	"add(4 , add(1,4))":        9,
	"add(add(1,3) , add(1,4))": 9,

	"add(add(add(add(1,2),add(1,2)),1) , add(1,4))": 12,
}

func TestParseNested(t *testing.T) {
	for str, res := range tableNested {
		o := Parse(str)
		if o != res {
			t.Errorf("exp %v got %v, for '%v'",
				res, o, str)
		}
	}
}
