func longestConsecutive(nums []int) int {
    set := make(map[int]bool)
    for _, num := range nums {
        set[num] = true
    }

    maxLength := 0

    for num := range set {
        if !set[num-1] { 
            currentNum := num
            count := 1

            for set[currentNum+1] {
                currentNum++
                count++
            }

            if count > maxLength {
                maxLength = count
            }
        }
    }

    return maxLength
}
