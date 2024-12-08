package functions

import (
	"errors"
    "fmt"
    "slices"
)

func ReplaceInStr(str string, ch string, ind int) string {
	return str[:ind] + ch + str[ind+1:]
}

// Returns the index if the character is inside the string
func CharInStr(ch string, str string) (int, error) {
	if len(ch) != 1 {
		return -1, errors.New("The character must be a string with length 1")
	}

	for i, c := range str {
		if string(c) == ch {
			return i, nil
		}
	}
	return -1, nil
}

// Returns false if the point `xy` is out of bounds given a height (y) and width (x)
func InsideOfBounds(xy []int, h int, w int) bool {
    if xy[0] < 0 || xy[1] < 0 || xy[0] >= w || xy[1] >= h {
        return false
    }
    return true
}

func PrintMatrix(mat []string) {
    for _, l := range mat {
        fmt.Println(l)
    }
}

func ArrInMat(s []int, m [][]int) bool {
    for _, v := range m {
        if slices.Equal(v, s) {
            return true
        }
    }
    return false
}
