func minMovesToSeat(seats []int, students []int) int {
    sort.Ints(seats)
    sort.Ints(students)
    sum := 0
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