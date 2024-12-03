package functions

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Sign(n int) int {
	if n < 0 {
		return -1
	} else if n > 0 {
		return +1
	}
	return 0
}

/*
 * Check all values in `v` satisfy the condition `check`
 */
func All[T int](v []T, check func(T) bool) bool {
	for _, val := range v {
		if !check(val) {
			return false
		}
	}
	return true
}

/*
 * Same as above, but return the count of False
 */
func AllCount[T int](v []T, check func(T) bool) int {
	var out int
	for _, val := range v {
		if !check(val) {
			out++
		}
	}
	return out
}

/*
 * Return index of different element in array (assuming there is only 1)
 */
func FindDifferent(v []int) int {
	var count_changes int // Count the number of changes in the values of v
	var ind int
	for i, n := range v[1:] {
		if n != v[i] {
			count_changes++
		}
		if count_changes == 2 {
			ind = i
		}
	}
	return ind
}
