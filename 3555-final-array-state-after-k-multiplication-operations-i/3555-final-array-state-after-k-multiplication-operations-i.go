func getFinalState(nums []int, k int, multiplier int) []int { 
   for i:=0; i<k; i++ {
    min := 0
    for j:=1; j<len(nums); j++ {
        if nums[j] < nums[min] {
            min = j
        }
    }
    nums[min] = nums[min] * multiplier
   }
   return nums
}