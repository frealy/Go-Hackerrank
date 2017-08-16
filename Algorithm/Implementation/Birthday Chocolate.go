package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	x := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&x[i])
	}
	var d, m int
	fmt.Scan(&d, &m)

	var count int
	var sum int
	for i := 0; i < n; i++ {
		if m == 1 && x[i] == d {
			count = count + 1
		} else if m == 2 && i < (n-2) {
			if (x[i] + x[i+1]) == d {
				count = count + 1
			}
		} else {
			for j := i; j < (i+m) && i <= (n-m); j++ {
				sum = sum + x[j]
			}
			if sum == d {
				count = count + 1
			}
		}
		sum = 0
	}

	fmt.Print(count)
}
