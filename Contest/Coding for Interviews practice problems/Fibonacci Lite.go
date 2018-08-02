package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int64
	fmt.Scan(&n)

	var n1, n2 int64 = 0, 1
	for i := 0; i < int(n); i++ {
		temp := n1
		n1 = n2
		n2 = n2 + temp
	}

	if n == 0 {
		n1 = 0
	} else if n == 1 {
		n1 = 1
	}

	fmt.Println(n1)
}
