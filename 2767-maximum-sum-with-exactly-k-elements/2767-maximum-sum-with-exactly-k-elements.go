func maximizeSum(nums []int, k int) int {
    max := 0
    sum :=0
    for _, v := range nums  {
        if v > max {
            max = v
        }
    }
    for i:= 1; i<=k; i++ {
        sum += max
        max++
    }
    return sum
}