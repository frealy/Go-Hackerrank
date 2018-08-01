package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}

	var count int
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if (x[i]+x[j])%m == 0 {
				count = count + 1
			}
		}
	}

	fmt.Print(count)
}
