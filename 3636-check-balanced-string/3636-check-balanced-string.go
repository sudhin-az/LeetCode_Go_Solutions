func isBalanced(num string) bool {
    evenSum := 0
    oddSum := 0
    for i, ch := range num {
        digit := int(ch - '0')
        if i % 2 == 0 {
            evenSum += digit
        } else {
            oddSum += digit
        }
    }
    return evenSum == oddSum
}