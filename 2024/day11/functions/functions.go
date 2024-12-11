package functions

import (
	"log"
	"strconv"
)

func SumSlices(s1 []int, s2 []int) []int {
	if len(s1) != len(s2) {
		log.Fatal("Slice lengths must match! ", len(s1), " ", len(s2))
	}
	out := make([]int, len(s1))
	for i := 0; i < len(s1); i++ {
		out[i] = s1[i] + s2[i]
	}
	return out
}

func StrToInt(s string) int {
	out, err := strconv.Atoi(s)
	if err != nil {
        log.Fatal("Cannot convert to int:", s)
	}
	return out
}
