func firstPalindrome(words []string) string {
    for _, v := range words {
        if Palindrome(v) {
        return v
    }}
    return ""
} 
func Palindrome(words string) bool {
   left, right := 0, len(words)-1
   for left < right {
   if words[left] != words[right] {
    return false
   }
   left++
   right--
   }
   return true
}