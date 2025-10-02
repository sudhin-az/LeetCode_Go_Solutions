func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, num := range nums {
        diff := target - num
        if j, found := m[diff]; found {
            return []int{i, j}
        }
        m[num] = i
    }
    return nil
}