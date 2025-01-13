func maxFrequencyElements(nums []int) int {
    max := 0
    res := 0
    count := make(map[int]int)
    for _, v := range nums {
        count[v]++
    if count[v] > max {
        max = count[v]
        res = 0
    }
    if count[v] == max {
        res += count[v]
    }
    }
    return res
}