import (
	"fmt"
	"sort"
)
func trimMean(arr []int) float64 {
    sort.Ints(arr)
    n := len(arr)
    remove := n / 20

    sum :=0
    count :=0
    for i:=remove; i<n - remove; i++ {
        sum += arr[i]
        count ++
    }
    return float64(sum) / float64(count)
}