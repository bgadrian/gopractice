package _6_26_calculator_arithmetic_no_parant

import (
	"bytes"
	"log"
	"strconv"
	"strings"
)

type opf func(a, b *elem) float64

var op = map[string]opf{
	"+": func(a, b *elem) float64 { return a.number + b.number },
	"-": func(a, b *elem) float64 { return a.number - b.number },
	"*": func(a, b *elem) float64 { return a.number * b.number },
	"/": func(a, b *elem) float64 { return a.number / b.number },
}

//a set of the priority operators
var impOp = map[string]struct{}{"*": {}, "/": {}}

var doOperation = func(a, operator, b *elem) *elem {
	if a.isOperator || b.isOperator || operator.isOperator == false {
		log.Panicf("malformed doOperation, wrong elements")
	}
	asFloat := op[operator.text](a, b)
	asString := strconv.FormatFloat(asFloat, 'f', 2, 32)
	return newElem(asString)
}

type elem struct {
	isOperator bool
	isImpOp    bool
	text       string
	number     float64
}

func isOperator(s string) bool {
	_, isSign := op[s]
	return isSign
}

func newElem(s string) *elem {
	s = strings.TrimSpace(s)
	var e = &elem{text: s}
	e.isOperator = isOperator(s)

	if e.isOperator {
		_, e.isImpOp = impOp[s]
	} else {
		val, err := strconv.ParseFloat(e.text, 32)
		if err != nil {
			log.Println(err)
			log.Panicf("failed to format '%v'", s)
		}
		e.number = val
	}
	return e
}

type inputWrapper struct {
	original string
	index    int
}

func (w *inputWrapper) nextElem() (res *elem) {
	if w.index >= len(w.original) {
		return nil
	}

	if isOperator(string(w.original[w.index])) {
		res = newElem(string(w.original[w.index]))
		w.index++
		return
	}

	//we have a number, we keep searching until the end or a sign
	var builder bytes.Buffer
	for ; w.index < len(w.original) && isOperator(string(w.original[w.index])) == false; w.index++ {
		builder.WriteByte(w.original[w.index])
	}
	res = newElem(builder.String())
	return
}

/* 16.26 Calculator: Given an arithmetic equation consisting of positive integers, +, -, * and / (no parentheses),
compute the result.
EXAMPLE
Input: 2*3+5/6*3+15
Output: 23.5

Procedural approach with 2 structures.
*/
func solution(input string) float64 {
	if len(input) == 0 {
		return 0
	}
	if len(input) <= 2 {
		return 0 //and error
	}

	in := &inputWrapper{original: input}
	var processing = []*elem{newElem("+"), in.nextElem()}
	var result = newElem("0")

	for current := in.nextElem(); current != nil; current = in.nextElem() {
		if current.isOperator == false {
			//is a number
			//we shouldn't reach this stage because after each operator
			//we advance once more to get the operand
			//and we start from the first operator
			//this can happen when the input is malformed: 1**2
			log.Panicf("wtf '%v'", current)
		}

		if current.isImpOp {
			processing[1] = doOperation(processing[1], current, in.nextElem())
			continue
		}

		//is a +/- sign
		result = doOperation(result, processing[0], processing[1])
		processing[0] = current
		processing[1] = in.nextElem()
		continue
	}
	//one last time
	result = doOperation(result, processing[0], processing[1])

	return result.number
}
