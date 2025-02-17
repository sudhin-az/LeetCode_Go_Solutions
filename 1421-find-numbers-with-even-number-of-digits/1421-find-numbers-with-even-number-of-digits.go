func findNumbers(nums []int) int {
    count := 0

    for i:=0; i<len(nums); i++ {
        dig := strconv.Itoa(nums[i])
        di := len(dig)
        if di % 2 == 0 {
            count ++
        }
    }
    return count
}