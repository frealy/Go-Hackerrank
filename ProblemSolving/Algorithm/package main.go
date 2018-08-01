package main

import "fmt"

func main() {
	var n, m int64
	fmt.Scan(&n, &m)

	value := make([]int64, n)
	var i int64
	for i = 0; i < n; i++ {
		fmt.Scan(&value[i])
	}

	var val int64
	fmt.Scan(&val)

	var sum int64 = 0
	for i = 0; i < n; i++ {
		if m != i {
			sum = sum + value[i]
		}
	}

	if val == sum/2 {
		fmt.Print("Bon Appetit")
	} else {
		fmt.Print(val - sum/2)
	}
}
