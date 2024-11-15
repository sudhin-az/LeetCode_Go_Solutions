func getSneakyNumbers(nums []int) []int {
    nestedMap := make(map[int]int)
    var res []int

    for _, v := range nums{
        nestedMap[v]++
        if nestedMap[v] > 1 {
            res = append(res, v)
        }
    }
    return res
}