func largestGoodInteger(num string) string {
    max := ""
    for i:=0; i<len(num)-2; i++ {
        if num[i] == num[i+1] && num[i+1] == num[i+2] {
            triplet := num[i:i+3]
        
        if triplet > max {
            max = triplet
        }
    }}
    return max
}
