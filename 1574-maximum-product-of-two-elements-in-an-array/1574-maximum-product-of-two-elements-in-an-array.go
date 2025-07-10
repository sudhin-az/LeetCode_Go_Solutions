func maxProduct(nums []int) int {
    max := 0
    for i:=0; i<len(nums); i++ {
        for j:=i+1; j<len(nums); j++ {
            product := (nums[i]-1) * (nums[j]-1)
                if product > max {
                    max = product
                }
            }
        }
    return max

}