func applyOperations(nums []int) []int {
    n :=len(nums)
    for i:=0; i<n-1; i++ {
        if nums[i] == nums[i+1] {
            nums[i] *=2
                nums[i+1] = 0
        }
    }
    result := make([]int, 0, n)
    zerocount := 0

    for _, num := range nums {
        if num != 0 {
            result = append(result, num)
        } else {
        zerocount ++
    }
    }
    for i:=0; i<zerocount; i++ {
        result = append(result, 0)
    }
    return result
}