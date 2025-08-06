func findDifference(nums1 []int, nums2 []int) [][]int {
    map1 := make(map[int]bool)
    map2 := make(map[int]bool)
    for _, num := range nums1 {
        map1[num] = true
    }
    for _, num := range nums2 {
        map2[num] = true
    }
    res1 := []int{}
    res2 := []int{}
    for num := range map1 {
        if !map2[num] {
            res1 = append(res1, num)
        }
    }
    for num := range map2 {
        if !map1[num] {
            res2 = append(res2, num)
        }
    }
    return [][]int{res1, res2}
}