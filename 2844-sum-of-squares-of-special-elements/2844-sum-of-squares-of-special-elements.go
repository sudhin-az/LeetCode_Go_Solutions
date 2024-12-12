func sumOfSquares(nums []int) (c int) {
   for i, v := range nums {
    if len(nums) % (i+1) == 0 {
        c += v * v
    }
   }
   return
}