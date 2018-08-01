package main

import "fmt"

func main() {
	var n int
	var x1, x2, x3, x4, x5 int = 1, 2, 3, 4, 5
	fmt.Scan(&n)

	val := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&val[i])
	}
	value := make([]int, 5)

	for i := 0; i < n; i++ {
		if val[i] == x1 {
			value[0] = value[0] + 1
		} else if val[i] == x2 {
			value[1] = value[1] + 1
		} else if val[i] == x3 {
			value[2] = value[2] + 1
		} else if val[i] == x4 {
			value[3] = value[3] + 1
		} else if val[i] == x5 {
			value[4] = value[4] + 1
		}
	}

	var max int
	var index int
	for i := 0; i < len(value); i++ {
		if max < value[i] {
			max = value[i]
			index = i + 1
		}
	}

	fmt.Print(index)
}
