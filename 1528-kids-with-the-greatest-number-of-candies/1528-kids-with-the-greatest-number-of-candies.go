func kidsWithCandies(candies []int, extraCandies int) []bool {
    max := 0
    for _, v := range candies {
        if v > max {
            max = v
        }
    }
    
    res := make([]bool, len(candies))
    for i:=0; i<len(candies); i++ {
        if candies[i] + extraCandies >= max {
            res[i] = true
        }
    }
    return res
}