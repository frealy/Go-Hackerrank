package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	val := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&val[i])
	}

	SockMerchant := make(map[int]int)
	for i := 0; i < n; i++ {
		SockMerchant[val[i]] = SockMerchant[val[i]] + 1
	}
	sum := 0
	for _, v := range SockMerchant {
		sum = sum + v/2
	}

	fmt.Print(sum)
}
