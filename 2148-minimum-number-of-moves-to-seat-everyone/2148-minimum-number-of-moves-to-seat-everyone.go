func minMovesToSeat(seats []int, students []int) int {
    sum := 0
    sort.Ints(seats)
    sort.Ints(students)
    for i:=0; i<len(seats); i++ {
        sum += abs(seats[i] - students[i])
    }
    return sum
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}