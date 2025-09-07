func countConsistentStrings(allowed string, words []string) int {
    count := 0
   for _, s := range words {
   for k, v := range s {
    if strings.Contains(allowed, string(v)) == false {
        break
    }
    if k == len(s) - 1 {
        count ++ 
    }}
   }
   return count
}