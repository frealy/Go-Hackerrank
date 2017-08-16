package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    x := make([]int64, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&x[i])
    }
    
    var max int64 = x[0]
    var min int64 = x[0]
    var maxi, mini int64
    for i := 0; i < n; i++ {
        if max < x[i] {
            max = x[i]
            maxi = maxi + 1
        }
        if min > x[i] {
            min = x[i]
            mini = mini + 1
        }
    }
    
    fmt.Print(maxi, mini)
}