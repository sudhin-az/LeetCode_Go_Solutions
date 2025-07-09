func maxDivScore(nums []int, divisors []int) int {
   maxCount := -1
   bestDivisor := -1
   for i:=0; i<len(divisors); i++ {
    count := 0
    for j:=0; j<len(nums); j++ {
        if nums[j] % divisors[i] == 0 {
            count ++
        }
        if count > maxCount || (count == maxCount && divisors[i] < bestDivisor) {
            maxCount = count
            bestDivisor = divisors[i]
        }
    }
   }
   return bestDivisor
}