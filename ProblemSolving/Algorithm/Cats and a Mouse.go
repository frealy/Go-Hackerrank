package main

import "fmt"

func abs(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	x := make([]int, n*3)

	for i := 0; i < n*3; i = i + 3 {
		fmt.Scan(&x[i], &x[i+1], &x[i+2])
	}

	for i := 0; i < n*3; i = i + 3 {
		a := abs(x[i+2] - x[i])
		b := abs(x[i+2] - x[i+1])
		if a < b {
			fmt.Println("Cat A")
		} else if b < a {
			fmt.Println("Cat B")
		} else {
			fmt.Println("Mouse C")
		}
	}
}
