func removeDuplicates(nums []int) int {
    count := 0
    for i:=0; i<len(nums); i++ {
        if nums[i] != nums[count] {
            count ++ 
            nums[count] = nums[i]
        }
    }
    return count + 1
}