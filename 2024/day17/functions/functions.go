package functions

func Mod(a, b int) int {
	return (a%b + b) % b
}

func ZeroPadLeft(in string, l int) string {
    for len(in) < l {
        in = "0" + in
    }
    return in
}

func MultString(in string, t int) string {
    out := ""
    for _ = range t {
        out += in
    }
    return out
}
