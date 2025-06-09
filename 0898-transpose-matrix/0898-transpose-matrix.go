func transpose(matrix [][]int) [][]int {
    var result = make([][]int, len(matrix[0]))

    for key := range result {
        result[key] = make([]int, len(matrix))
    }
    for i:=0; i<len(matrix); i++ {
        for j:=0; j<len(matrix[i]); j++ {
            result[j][i] = matrix[i][j]
        }
    }
    return result
}