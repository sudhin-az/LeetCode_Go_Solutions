func smallestEvenMultiple(n int) int {
    for i:=2; i<=2*n; i++ {
        if i % n == 0 && i % 2 == 0 {
            return i
        }
    }
    return n
}