func maximumWealth(accounts [][]int) int {
    max := 0
    for i:=0; i<len(accounts); i++ {
        acc := 0
        for j:=0; j<len(accounts[i]); j++ {
            acc += accounts[i][j]
                
        }
        if acc > max {
                    max = acc
                }
    }
    return max
}