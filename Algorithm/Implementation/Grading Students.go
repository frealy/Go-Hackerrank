package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    grade := make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&grade[i])
    }
    for i := 0; i < n; i++ {
        if 5 - (grade[i] % 5) < 3 && grade[i] > 37 {
            grade[i] = grade[i] + (5 - grade[i] % 5)
        }
        fmt.Println(grade[i])
    }
}