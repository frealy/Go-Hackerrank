package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	A := make([]int, n)
	B := make([]int, m)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	for i := 0; i < m; i++ {
		fmt.Scan(&B[i])
	}

	var min int
	for i := 0; i < n; i++ {
		if min < A[i] {
			min = A[i]
		}
	}
	var max int = B[0]
	for i := 0; i < n; i++ {
		if max < A[i] {
			max = A[i]
		}
	}

	var x int = min
	var ax, bx, X int
	for min <= x && x <= max {
		for i := 0; i < n; i++ {
			if x%A[i] == 0 {
				ax = ax + 1
			}
		}
		for i := 0; i < m; i++ {
			if B[i]%x == 0 {
				bx = bx + 1
			}
		}
		if ax == n && bx == m {
			X = X + 1
		}

		ax = 0
		bx = 0
		x = x + min
	}

	fmt.Print(X)

}
