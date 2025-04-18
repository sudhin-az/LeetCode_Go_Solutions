func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    s1 := strings.Split(s, "")
    t1 := strings.Split(t, "")
    sort.Strings(s1)
    sort.Strings(t1)
    for i := range s1 {
        if s1[i] != t1[i] {
            return false
        }
    }
    return true
}