func finalValueAfterOperations(operations []string) int {
    x := 0
    for i:=0; i<len(operations); i++ {
        if operations[i] == "++X" || operations[i] == "X++" {
            x++
        } else if operations[i] == "--X" || operations[i] == "X--" {
            x--
        }
    }
    return x
}