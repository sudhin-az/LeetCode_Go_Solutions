func lengthOfLongestSubstring(s string) int {
    chars := make(map[rune]bool)
    left := 0
    maxLen := 0

    for right := 0; right < len(s); right ++ {
        ch := rune(s[right])

        for chars[ch] {
            delete(chars, rune(s[left]))
            left ++
        }
        chars[ch] = true
        currentLen := right - left + 1
        if currentLen > maxLen {
            maxLen = currentLen
        }
    }
    return maxLen
}
