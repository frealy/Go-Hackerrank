package main

import "fmt"

func main() {
	var s, n, m, count, left int
	fmt.Scan(&s, &n, &m)
	x := make([]int, n)
	y := make([]int, m)

	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&y[i])
	}

	left = s
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if s/(x[i]+y[j]) == 1 {
				if s%(x[i]+y[j]) < left {
					left = s % (x[i] + y[j])
					count = x[i] + y[j]
				}
			}
		}
	}

	if count > 0 {
		fmt.Print(count)
	} else {
		fmt.Print(-1)
	}
}
