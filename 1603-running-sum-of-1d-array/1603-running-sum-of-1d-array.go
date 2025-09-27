func runningSum(nums []int) []int {
    ans := []int{}
    sum := 0
    for i:=0; i<len(nums); i++ {
        sum += nums[i]
        ans = append(ans, sum)
    }
    return ans
}