func findWordsContaining(words []string, x byte) []int {
    ans := []int{}
    run := string(x)

    for i:=0; i<len(words); i++ {
        if strings.Contains(words[i], run) {
            ans = append(ans, i)
        }
    }
    return ans
}