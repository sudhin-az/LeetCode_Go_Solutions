func minimumAverage(nums []int) float64 {
   sort.Ints(nums)
   averages := []float64{}

   for len(nums)>0 {
    min := nums[0]
    max := nums[len(nums)-1]
    avg := float64(min+max)/2
    averages = append(averages, avg)
    nums = nums[1 :len(nums)-1]
   }
   minaverages := averages[0]
   for _, value := range averages {
    if value < minaverages {
        minaverages = value
    }
   }
   return minaverages
}