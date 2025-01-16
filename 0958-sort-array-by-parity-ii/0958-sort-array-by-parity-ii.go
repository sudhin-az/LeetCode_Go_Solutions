func sortArrayByParityII(nums []int) []int {
    res := make([]int, len(nums))
    even, odd := 0, 1
    for _, num := range nums {
        if num % 2 == 0 {
            res[even] = num
            even+=2
        } else {
            res[odd]=num
            odd += 2
        }
    }
    return res
}