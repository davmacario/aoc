package utils

// Returns the index of element x in slice y.
// Returns -1 if the element was not found.
func GetIndex[T comparable](x T, y []T) int {
	for i, p := range y {
		if p == x {
			return i
		}
	}
	return -1
}

// Positive modulo operation (native Go implementation returns negative values)
func Mod(a, b int) int {
	return (a%b + b) % b
}

// Replace `i`-th character of string `s` with `c`
func ReplaceCharInString(s string, i int, c rune) string {
	return s[:i] + string(c) + s[i+1:]
}

// Removes element `ind` from slice `sl`. Returns `true` on success (element
// found and removed)
func RemoveFromSlice[T comparable](sl []T, ind int) bool {
	if ind >= len(sl) {
		return false
	}
	sl = append(sl[:ind], sl[ind+1:]...)
	return true
}

// zero pad string `in` on the left until it has a length `l`
func ZeroPadLeft(in string, l int) string {
	for len(in) < l {
		in = "0" + in
	}
	return in
}

// Returns true if all elements of the slice are equal.
// If the slice contains no elements, returns true.
func AllEqualElements[T comparable](s []T) bool {
	if len(s) == 0 {
		return true
	}
	val := s[0]
	for _, v := range s[1:] {
		if v != val {
			return false
		}
	}
	return true
}

// Produce a slice by splitting a string in `numParts` equal parts
func SplitStringEqual(s string, numParts int) []string {
	// Slice of empty strings if requesting more parts than chars in s
	if numParts > len(s) {
		out := make([]string, numParts)
		for i := range numParts {
			out[i] = ""
		}
		return out
	}
	out := []string{}
	lenPart := len(s) / numParts

	for i := range numParts {
		start := i * lenPart
		end := min((i+1)*lenPart, len(s))
		out = append(out, s[start:end])
	}

	return out
}
