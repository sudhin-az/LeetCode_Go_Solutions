func mostWordsFound(sentences []string) int {
    max := 0
    for _, sentence := range sentences {
        words := len(strings.Fields(sentence))
        if words > max {
            max = words
        }
    }
    return max
}