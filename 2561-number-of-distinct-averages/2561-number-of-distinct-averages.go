func distinctAverages(nums []int) int {
    averages := make(map[float64]bool)
    sort.Ints(nums)

    for len(nums) > 1 {
        min := nums[0]
        max := nums[len(nums)-1]

        avg := float64(min+max)/2
        averages[avg]=true

        nums = nums[1:len(nums)-1]
    }
    return len(averages)
}