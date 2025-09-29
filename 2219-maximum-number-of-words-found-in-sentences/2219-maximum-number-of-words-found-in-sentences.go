func mostWordsFound(sentences []string) int {
    maxWords := 0
    for _, sentence := range sentences {
        count := 1
        for _, ch := range sentence {
            if ch == ' ' {
                count ++
            }
        }
        if count > maxWords {
            maxWords = count
        }
    }
    return maxWords
}