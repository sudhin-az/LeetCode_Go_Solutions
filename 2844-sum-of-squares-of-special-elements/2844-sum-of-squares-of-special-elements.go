func sumOfSquares(nums []int) (n int) {
    for i, v := range nums {
        if len(nums) % (i+1) == 0 {
            n+=v*v
        }
    }
    return 
}