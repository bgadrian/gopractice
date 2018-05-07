package nearest_time_same_digits

import "testing"

var table = map[string]string{
	//"170020": "170002",
	//"000000": "000000",
	//"235555": "235555",
	//"070708": "070807",
	"170000": "100700",
	//"120344": "120434",
	//"190059": "190509",
	//"245903": "245930",
}

var tableValidTimes = []string{
	"000000", "012344", "101010",
	"105900", "235959",
}

var tableInvalidTimes = []string{
	"240000", "237100", "120060",
}

var tableHelp = []string{
	"240000", "123456", "001100",
}

func TestHelpers(t *testing.T) {
	for _, v := range tableHelp {
		r := arrToInt(intToArr(v))
		if r == v {
			continue
		}
		t.Errorf("exp '%v', got '%v'",
			v, r)
	}
}

func TestIsValid(t *testing.T) {
	for _, v := range tableValidTimes {
		if isValid(intToArr(v)) {
			continue
		}
		t.Errorf("valid time not recognized '%v'", v)
	}

	for _, v := range tableInvalidTimes {
		if isValid(intToArr(v)) == false {
			continue
		}
		t.Errorf("Invalid time not recognized '%v'", v)
	}
}

func TestSolution(t *testing.T) {
	for in, out := range table {
		r := solution(in)
		if r == out {
			continue
		}
		t.Errorf("for '%v', exp '%v', got '%v'\n",
			in, out, r)
	}
}
