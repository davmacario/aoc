package utils

import (
	"fmt"
	"math/rand/v2"
	"strconv"
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
func RemoveFromSlice[T comparable](sl []T, ind int) ([]T, bool) {
	if ind >= len(sl) {
		return sl, false
	}
	sl = append(sl[:ind], sl[ind+1:]...)
	return sl, true
}

func ZeroPadLeft(in string, l int) string {
	for len(in) < l {
		in = "0" + in
	}
	return in
}

func PrintMap[T comparable, V any](m map[T]V) {
    for k, v := range m {
        fmt.Println(k, v)
    }
}

func CopyMap[K comparable, V any](source map[K]V) map[K]V {
	out := make(map[K]V)
	for k, v := range source {
		out[k] = v
	}
	return out
}

func RandIntRange(m, M int) int {
	return rand.IntN(M-m) + m
}

func IntToBin(a int) string {
	return strconv.FormatInt(int64(a), 2)
}

func BinToInt(a string) int {
	out, err := strconv.ParseInt(a, 2, 0)
	if err != nil {
		panic(err)
	}
	return int(out)
}

// Reverse a string
func ReverseString(s string) string {
	sl := []rune(s)
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
	return string(sl)
}

func FindDifferences(a, b string) []int {
	var diffs []int
	if len(a) != len(b) {
		panic("The provided strings must have the same length")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diffs = append(diffs, i)
		}
	}
	return diffs
}
