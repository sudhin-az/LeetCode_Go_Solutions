func numOfUnplacedFruits(fruits []int, baskets []int) int {
    placed := len(fruits)

    for i := range fruits {
        for j := range baskets {
            if baskets[j] >= fruits[i] {
                baskets[j] = 0
                placed -= 1
                break
            }
        }
    }
    return placed
}