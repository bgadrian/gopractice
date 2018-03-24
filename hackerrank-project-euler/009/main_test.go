package main

import "testing"

var table = map[int]int{
	12: 60, //3,4,5
	4:  -1, //
}

//https://en.wikipedia.org/wiki/Pythagorean_triple#Generating_a_triple
//var primitives = [][]int{
//	//{20, 99, 101}, {60, 91, 109},
//	{15, 112, 113},
//	//{44, 117, 125},
//	//{88, 105, 137},
//	//{17, 144, 145}, {24, 143, 145}, {51, 140, 149},
//	//{85, 132, 157}, {119, 120, 169}, {52, 165, 173}, {19, 180, 181},
//	//{57, 176, 185}, {104, 153, 185}, {95, 168, 193}, {28, 195, 197},
//	//{84, 187, 205}, {133, 156, 205}, {21, 220, 221}, {140, 171, 221},
//	//{60, 221, 229}, {105, 208, 233}, {120, 209, 241}, {32, 255, 257},
//	//{23, 264, 265}, {96, 247, 265}, {69, 260, 269}, {115, 252, 277},
//	//{160, 231, 281}, {161, 240, 289}, {68, 285, 293},
//}

func TestResult(t *testing.T) {
	for in, out := range table {
		r := getResult(in)
		if out != r {
			t.Errorf("for %v exp %v got %v", in, out, r)
		}
	}
}

//
//func TestResultPrimitives(t *testing.T) {
//	for _, prim := range primitives {
//		for m := 1; m < 2; m++ { //primitives can be used to generate other triplets
//			a := prim[0] * m
//			b := prim[1] * m
//			c := prim[2] * m
//			out := getResult(a + b + c)
//			exp := a * b * c
//			if exp != out {
//				t.Errorf("for %v exp %v got %v", []int{a, b, c}, exp, out)
//			}
//		}
//	}
//}
