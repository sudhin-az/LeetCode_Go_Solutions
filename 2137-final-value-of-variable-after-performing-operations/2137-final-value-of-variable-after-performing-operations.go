func finalValueAfterOperations(operations []string) int {
    x := 0
    for i:=0; i<len(operations); i++ {
        fmt.Println("sudhin",operations[i] == "++x" || operations[i] =="x++")
        fmt.Println(x + 1)
        if operations[i] == "++X" || operations[i] =="X++" {
            x++
        } else if operations[i] == "--X" || operations[i] == "X--" {
            x--
        }
        }     
    return x
}