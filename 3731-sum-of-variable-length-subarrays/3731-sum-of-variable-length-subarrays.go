func subarraySum(nums []int) int {
    sum := 0
    for i:=0; i<len(nums); i++ {
        start := max(0, i - nums[i])

        for j:=start; j<i+1; j++ {
            sum += nums[j]
        }
    }
    return sum
}