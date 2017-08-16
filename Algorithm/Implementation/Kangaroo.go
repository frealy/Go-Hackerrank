package main

import "fmt"

func main() {
	var x1, x2, y1, y2 int64
	fmt.Scan(&x1, &x2, &y1, &y2)

	if x1 <= y1 {
		if x2 <= y2 {
			fmt.Printf("NO")
		} else if x2 > y2 {
			var i int64 = 1
			var flag bool = true
			for x1+x2*i <= y1+y2*i {
				if (x1 + x2*i) == (y1 + y2*i) {
					fmt.Print("YES")
					flag = false
				}
				i = i + 1
			}
			if flag {
				fmt.Printf("NO")
			}
		}
	} else if x1 > y1 {
		if x2 >= y2 {
			fmt.Printf("NO")
		} else if x2 < y2 {
			var i int64 = 1
			var flag bool = true
			for x1+x2*i >= y1+y2*i {
				if (x1 + x2*i) == (y1 + y2*i) {
					fmt.Print("YES")
					flag = false
				}
				i = i + 1
			}
			if flag {
				fmt.Printf("NO")
			}
		}
	}
}
