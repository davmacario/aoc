package functions

import "log"


func SwapChars(str *string, i int, j int) {
	if i >= len(*str) || j >= len(*str) {
		log.Fatal("Index out of bounds")
	}
	tmpCh := str[i]
	(*str)[i] = *str[j]

}
