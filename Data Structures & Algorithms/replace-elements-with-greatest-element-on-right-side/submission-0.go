func replaceElements(arr []int) []int {
    ans := make([]int, len(arr))
    maxx := -1
    for i := len(arr) - 1; i >= 0; i-- {
        tmp := arr[i]
        ans[i] = maxx
        maxx = max(maxx, tmp)
    }
    return ans
}
