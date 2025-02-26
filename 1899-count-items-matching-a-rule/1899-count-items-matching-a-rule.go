func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    hashMap := map[string]int{
        "type":0,
        "color":1,
        "name":2,
    }
    var count int
    key := hashMap[ruleKey]
    for _, item := range items {
        if item[key] == ruleValue {
            count ++
        }
    }
    return count
}