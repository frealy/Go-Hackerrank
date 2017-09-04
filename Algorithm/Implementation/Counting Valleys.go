package main

import "fmt"

func main() {
	var n, count, valley int
	var flag bool = false
	fmt.Scan(&n)

	val := make([]byte, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%c", &val[i])
	}

	for i := 0; i < n; i++ {
		if val[i] == 'U' {
			valley = valley + 1
			if valley > 0 {
				flag = false
			}
		} else if val[i] == 'D' {
			valley = valley - 1
			if valley < 0 {
				flag = true
			}
		}

		if valley == 0 && flag {
			count = count + 1
		}
		//fmt.Println(valley, count, flag)
	}
	fmt.Printf("%d", count)
}
