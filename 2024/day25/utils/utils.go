package utils

import (
	"fmt"
)

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

func ZeroPadLeft(in string, l int) string {
	for len(in) < l {
		in = "0" + in
	}
	return in
}

// Reverse a string
func ReverseString(s string) string {
	sl := []rune(s)
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
	return string(sl)
}

func PrintMap[T comparable, V any](m map[T]V) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func SumSlice(a, b []int) []int {
	if len(a) != len(b) {
		panic("Provided slices have different length")
	}
	out := make([]int, len(a))
	for i, n := range a {
		out[i] = n + b[i]
	}
	return out
}

func All[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}