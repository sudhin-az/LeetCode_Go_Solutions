func largestLocal(grid [][]int) [][]int {
    n := len(grid)
    result := [][]int{}
    for i:=0; i<n-2; i++ {
        row := []int{}
        for j:=0; j<n-2; j++ {
            max := grid[i][j]

            for x:=i; x<=i+2; x++ {
                for y:=j; y<=j+2; y++ {
                    if grid[x][y] > max {
                        max = grid[x][y]
                    }
                }
            }
            row = append(row, max)
        }
        result = append(result, row)
    }
    return result
}