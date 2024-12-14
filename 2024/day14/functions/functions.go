package functions

// Returns the *positive* modulo a%b
func Mod(a, b int) int {
	return (a%b + b) % b
}

// Replace character in position `i` of string `in` with `r`
func ReplaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
