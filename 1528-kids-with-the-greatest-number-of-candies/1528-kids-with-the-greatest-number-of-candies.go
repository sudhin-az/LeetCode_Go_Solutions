func kidsWithCandies(candies []int, extraCandies int) []bool {
    ans := []bool{}
    max := 0
    for _, v := range candies {
        if v > max {
            max = v
        }
    }
    for i:=0; i<len(candies); i++ {
            if candies[i] + extraCandies >= max {
                ans = append(ans, true)
            } else {
                ans = append(ans, false)
            }
        }
    return ans
}