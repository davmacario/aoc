package functions

import "log"

func SwapChars(str []string, i int, j int) {
	if i >= len(str) || j >= len(str) || i < 0 || j < 0 {
        log.Fatal("Index out of bounds; i: ", i, ", j: ", j)
	}
	if i == j {
		return
	}
	tmp := str[j]
	str[j] = str[i]
	str[i] = tmp
}
