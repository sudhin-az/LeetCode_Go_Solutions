func calPoints(operations []string) int {
    record :=[]int{}
    for _, op := range operations {
        if op == "C" {
            record = record[:len(record)-1]
        } else if op == "D" {
            record = append(record, 2 *record[len(record)-1])
        } else if op== "+" {
            record = append(record, record[len(record)-1] + record[len(record)-2])
        } else {
            num, _ := strconv.Atoi(op)
            record = append(record, num)
        }
    }
    total := 0
    for _, v := range record {
        total += v
    }
    return total
}