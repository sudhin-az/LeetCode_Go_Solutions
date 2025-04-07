func canAliceWin(nums []int) bool {
    totalSum := 0
    singleSum := 0
    doubleSum := 0
    for _, num := range nums {
        totalSum += num
    
    if num >=1 && num <= 9 {
        singleSum += num
    } else if num >=10 && num <= 99 {
        doubleSum += num
    }
    }
    if singleSum > totalSum - singleSum || doubleSum > totalSum - doubleSum {
        return true
    }
    return false
}