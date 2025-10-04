func createTargetArray(nums []int, index []int) []int {
    ans := []int{}
    for i:=0; i<len(nums); i++ {
        idx := index[i]
        val := nums[i]
        if idx == len(ans) {
            ans = append(ans, val)
        } else {
            ans = append(ans, 0)
            copy(ans[idx+1:], ans[idx:])
            ans[idx] = val
        }
    }
    return ans
}