func transformArray(nums []int) []int {
    ans := []int{}
    odd := 0
    even := 0
    for i:=0; i<len(nums); i++ {
        if nums[i] % 2 == 0 {
            even = 0
            ans = append(ans, even)
        } else {
            odd = 1
            ans = append(ans, odd)
        }
        sort.Ints(ans)
    }
    return ans
}