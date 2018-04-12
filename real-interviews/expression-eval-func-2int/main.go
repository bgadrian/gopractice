package vald

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type op func(a, b int) int

var ops = map[string]op{
	"add":  func(a, b int) int { return a + b },
	"sub":  func(a, b int) int { return a - b },
	"mult": func(a, b int) int { return a * b },
}

/*Parse
Write a function that will take a simple expression and evaluate it,
 supporting only integers and the op  “add”, “sub”, and “mult”
Ex add(sub(5, 4), 1)
*/
func Parse(s string) int {
	sleft := s
	//foundOps := make(map[string]string) //op => params
	reg, err := regexp.Compile(`^([a-z]{3,4})(\(.+\))$`)
	if err != nil {
		log.Panic(err)
	}

	type sub func(s string, this sub) int

	rec := func(exp string, this sub) int {
		//if is only numeric, then is just an integer
		num, err := strconv.Atoi(strings.TrimSpace(exp))
		if err == nil {
			return num
		}

		results := reg.FindStringSubmatch(strings.TrimSpace(exp))
		if len(results) != 3 {
			log.Panicf("malformed input (b) '%v'", sleft)
		}
		//exp contains 2 groups, the op and the params
		op := results[1]
		params := strings.TrimSpace(results[2])
		params = params[1 : len(params)-1] //remove the ( and )
		splited := splitParams(params)

		if splited == nil || len(splited) != 2 {
			log.Panicf("malformed input (comma) '%v'", sleft)
		}

		a := this(splited[0], this)
		b := this(splited[1], this)

		opFunc, exists := ops[op]
		if exists == false {
			log.Panicf("malformed input (op) '%v'", sleft)
		}
		return opFunc(a, b)
	}

	return rec(s, rec)
}

func splitParams(s string) []string {
	leftParan := 0
	pivotIndex := 0

	for i, r := range s {
		if r == '(' {
			leftParan++
		}
		if r == ')' {
			leftParan--
		}
		if leftParan == 0 && r == ',' {
			pivotIndex = i
			break
		}
	}

	if pivotIndex == 0 {
		return nil
	}
	return []string{
		s[:pivotIndex],
		s[pivotIndex+1:],
	}
}
