func unequalTriplets(nums []int) int {
    count := 0
    for i:=0; i<len(nums); i++ {
        for j:=i+1; j<len(nums); j++ {
            for k :=j+1; k<len(nums); k++ {
                if nums[i] != nums[j] && nums[i] != nums[k] && nums[j] != nums[k] {
                    count ++
                }
            }
        }
    }
    return count
}