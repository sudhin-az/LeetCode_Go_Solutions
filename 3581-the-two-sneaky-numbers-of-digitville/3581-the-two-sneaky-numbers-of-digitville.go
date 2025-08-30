func getSneakyNumbers(nums []int) []int {
    ans := []int{}
    for i:=0; i<len(nums)-1; i++  {
        for j:=i+1; j<len(nums); j++ {
            fmt.Println(nums[i] == nums[j])
            fmt.Println(nums[i], nums[j])
            if nums[i] == nums[j] {
                ans = append(ans, nums[i])
            }
        }
    }
    return ans
}