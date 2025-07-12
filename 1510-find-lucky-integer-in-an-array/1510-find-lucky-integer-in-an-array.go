func findLucky(arr []int) int {
    freq := make(map[int]int)

    for _, v := range arr {
        freq[v]++
    }
    max := -1
    for k, v := range freq {
        if k == v && k > max {
            max = k
        }
    }
    return max
}