package functions

func ReplaceCharInString(s string, i int, c rune) string {
	return s[:i] + string(c) + s[i+1:]
}
