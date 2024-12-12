func sumOfSquares(nums []int) (c int) {
   for i, j := range nums {
    if len(nums) % (i+1) == 0 {
        c += j * j
    }
   }
   return
}