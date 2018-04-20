package __5one_away_edit_strings

/*

One Away: There are three types of edits that can be performed on strings: insert a character,
remove a character, or replace a character. Given two strings, write a function to check if they are
one edit (or zero edits) away.
EXAMPLE
pale, ple -> true
pales, pale -> true
pale, bale -> true
pale, bae -> false
*/

func solution(a, b []rune) bool {
	//replacement => 1 difference in strings, equal len
	//insert => 1 extra char, 1 char diff, +1 len
	//remove => insert reverse, 1 char diff, -1 len

	diff := len(a) - len(b)
	if diff < -1 || diff > 1 {
		return false
	}

	if diff == 0 {
		allowed := 1
		//only 1 change allowed, same lengths
		for i, ra := range a {
			if ra == b[i] {
				continue
			}
			allowed--
			if allowed < 0 {
				return false
			}
		}
		return true
	}

	shorter, longer := a, b //a-b=-1
	if diff == 1 {
		shorter, longer = b, a
	}

	//we seek the first change, then the rest must be equal
	for ps, pl := 0, 0; ps < len(shorter); ps, pl = ps+1, pl+1 {
		if longer[pl] == shorter[ps] {
			continue
		}

		firstDiff := ps == pl
		if firstDiff == false {
			//is the 2nd
			return false
		}

		ps-- //skip current difference by returning the shorter one to previous
	}

	return true
}
