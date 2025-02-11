func isPrefixString(s string, words []string) bool {
    var word string
   for i:= 0; i<len(words); i++ {
    word += words[i]
    if s == word {
        return true
    }
}
return false
}