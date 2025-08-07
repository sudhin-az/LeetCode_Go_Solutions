func runningSum(nums []int) []int {
    sum := 0
    res := []int{}
  for _, v := range nums {
        sum +=v 
        res = append(res, sum)
  }
  return res
}