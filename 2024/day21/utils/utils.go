package utils

import "unicode"

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

func GetNumCharsOnly(s string) (out string) {
	for _, c := range s {
		if unicode.IsDigit(c) {
			out += string(c)
		}
	}
	return out
}

func Sign(a int) int {
	if a > 0 {
		return +1
	} else if a < 0 {
		return -1
	} else {
		return 0
	}
}

func IntAbs(a int) int {
	return a * Sign(a)
}

// Reverse a string
func ReverseString(s string) string {
	sl := []rune(s)
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
	return string(sl)
}

// Recursively generate all
func PermutationsOfString(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}
	out := make([]string, 0)
	for i, c := range s {
		remaining := s[:i] + s[i+1:]
		perm_red := PermutationsOfString(remaining)
		for _, s := range perm_red {
			out = append(out, string(c)+s)
		}
	}
	return out
}
