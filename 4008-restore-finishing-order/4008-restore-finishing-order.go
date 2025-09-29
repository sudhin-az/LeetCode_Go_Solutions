func recoverOrder(order []int, friends []int) []int {
    ans := []int{}
    for i:=0; i<len(order); i++ {
        for j:=0; j<len(friends); j++ {
            if order[i] == friends[j] {
                ans = append(ans, order[i])
            }
        }
    }
    return ans
}