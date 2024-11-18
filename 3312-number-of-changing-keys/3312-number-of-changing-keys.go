func countKeyChanges(s string) int {
    str := strings.ToLower(s)
    sb := strings.Split(str, "")
    count :=0
    for i:=0; i<len(sb)-1; i++ {
            if sb[i] != sb[i+1] {
                count++
            }
    }
    return count
}