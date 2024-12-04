package functions

import (
	"fmt"
	"log"
)

func DispMatrix(mat []string) {
	for _, row := range mat {
        fmt.Println(row)
	}
}

// Sum two arrays element-wise (checks equal length as well)
func SumArr(a1 []int, a2 []int) []int {
	if len(a1) != len(a2) {
		log.Fatal("The arrays to be added should have the same length")
	}
	out := make([]int, len(a1))
	for i := 0; i < len(a1); i++ {
		out[i] = a1[i] + a2[i]
	}
	return out
}

func MultArray(a1 []int, times int) []int {
	out := make([]int, len(a1))
	for i := 0; i < len(a1); i++ {
		out[i] = a1[i] * times
	}
	return out
}

// Returns true if n is in the range [rng[0], rng[1])
func InRange(n int, rng [2]int) bool {
    if n >= rng[0] && n < rng[1] {
        return true
    }
    return false
}
