package functions

/*
 * Perform binary search on an array `lst` to return the index of the element
 * that is immediately smaller or equal than `val`.
 */
func BinSearch(lst []int, val int) int {
	if len(lst) <= 1 {
		return 0
	}

	pos := len(lst) / 2 // Integer division

	if val > lst[pos] {
		return pos + BinSearch(lst[pos:], val)
	} else if val < lst[pos] {
		return 0 + BinSearch(lst[:pos], val)
	} else {
		return pos
	}
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

/*
 * Print count of values in list
 */
func UniqueValuesCount(lst []int) map[int]int {
	dict := make(map[int]int)
	for _, num := range lst {
		dict[num] = dict[num] + 1
	}
	return dict
}
