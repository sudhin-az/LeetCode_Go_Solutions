func findNumbers(nums []int) int {
    count :=0
    for i:=0; i<len(nums); i++ {
        digits := strconv.Itoa(nums[i])
        dig := len(digits)
        if dig % 2 == 0 {
            count ++
        }
    }
    return count
}