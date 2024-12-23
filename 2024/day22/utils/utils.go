package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
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

func SliceToStr(s []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(s)), ","), "[]")
}

// Reverses `SliceToString`
func StrToSlice(s string) []int {
	out := make([]int, 0)
	for _, v := range strings.Split(s, ",") {
		v_int, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
		out = append(out, v_int)
	}
	return out
}
