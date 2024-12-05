package functions

// Returns true if `n` is included in the slice `sl`
func InSlice(n int, sl []int) bool {
    for _, v := range sl {
        if v == n {
            return true
        }
    }
    return false
}
