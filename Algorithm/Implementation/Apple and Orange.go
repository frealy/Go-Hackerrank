package main

import "fmt"

func main() {
	var s, t int64
	var a, b int64
	var m, n int64
	fmt.Scan(&s, &t)
	fmt.Scan(&a, &b)
	fmt.Scan(&m, &n)

	apple := make([]int64, m)
	orange := make([]int64, n)
	var i int64
	for i = 0; i < m; i++ {
		fmt.Scan(&apple[i])
	}
	for i = 0; i < n; i++ {
		fmt.Scan(&orange[i])
	}

	var AppleSum int64 = 0
	var OrangeSum int64 = 0

	for i = 0; i < m; i++ {
		if s <= apple[i]+a && apple[i]+a <= t {
			AppleSum = AppleSum + 1
		}
	}
	for i = 0; i < n; i++ {
		if s <= orange[i]+b && orange[i]+b <= t {
			OrangeSum = OrangeSum + 1
		}
	}

	fmt.Println(AppleSum)
	fmt.Println(OrangeSum)
}
