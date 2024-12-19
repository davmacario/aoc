package functions

func Mod(a, b int) int {
	return (a%b + b) % b
}

func ReplaceCharInString(s string, i int, c rune) string {
	return s[:i] + string(c) + s[i+1:]
}

func FindInMat(m []string, c rune) (i, j int) {
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			if rune(m[i][j]) == c {
				return i, j
			}
		}
	}
	return -1, -1
}

// Returns the index of x in slice y. Returns -1 if the element is not found in
// the slice
func GetIndex[T comparable](x T, y []T) int {
	for i, p := range y {
		if p == x {
			return i
		}
	}
	return -1
}
