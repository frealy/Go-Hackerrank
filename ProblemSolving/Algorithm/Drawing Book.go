package main

import "fmt"

func main() {
	var n, p int
	fmt.Scan(&n, &p)

	var min1, min2, length int

	for i := 1; i < p; i = i + 2 {
		min1 = min1 + 1
	}

	if n%2 == 0 {
		length = n
	} else {
		length = n - 1
	}
	for i := length; i > p; i = i - 2 {
		min2 = min2 + 1
	}

	if p == 1 {
		fmt.Print(0)
	} else if p == n && n%2 == 0 {
		fmt.Print(0)
	} else if (p == n-1 || p == n) && n%2 == 1 {
		fmt.Print(0)
	} else if min1 <= min2 {
		fmt.Print(min1)
	} else {
		fmt.Print(min2)
	}
}
