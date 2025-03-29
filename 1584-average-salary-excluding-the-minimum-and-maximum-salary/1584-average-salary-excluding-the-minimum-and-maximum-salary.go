func average(salary []int) float64 {
    sum :=0
    min := salary[0]
    max := salary[0]
    for _, v := range salary {
        sum += v
    
    if v > max {
        max = v
    }
    if v < min {
        min = v
    }}
    return float64(sum-min-max) / float64(len(salary)-2)
}