func countSeniors(details []string) int {
    count := 0
    for _, num := range details {
        if num[11:13] > "60" {
            count ++
        }
    }
    return count 
}