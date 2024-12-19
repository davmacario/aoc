package functions

// Returns true if s1 starts with s2
func StartsWith(s1, s2 string) bool {
    if len(s2) > len(s1) {
        return false
    }
    return s2 == s1[:len(s2)]
}
