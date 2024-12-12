func splitWordsBySeparator(words []string, separator byte) []string {
    arr := []string{}
    for _, v := range words {
        split := strings.Split(v, string(separator))
        for _, j := range split {
            if j != ""{
            arr = append(arr, j)
        }
    }}
    return arr
}