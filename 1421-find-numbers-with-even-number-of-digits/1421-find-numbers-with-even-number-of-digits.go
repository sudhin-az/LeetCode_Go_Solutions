func findNumbers(nums []int) int {
  count :=0
  for _, num := range nums {
    digitcount :=0
    for num >0 {
        digitcount ++
        num = num/10
    }
    if digitcount % 2 == 0 {
        count ++
    }
  }
  return count
}