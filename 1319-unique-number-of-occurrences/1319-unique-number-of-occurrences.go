func uniqueOccurrences(arr []int) bool {
  freq := make(map[int]int)
  seen := make(map[int]bool)
  for _, v := range arr {
    freq[v]++
  }
  for _, v := range freq {
    if _, ok := seen[v]; ok {
        return false
    }
    seen[v] = true
  }
  return true
}