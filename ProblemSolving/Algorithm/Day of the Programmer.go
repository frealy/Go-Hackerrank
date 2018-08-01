package main

import "fmt"

func main() {
	var Month = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	var N int
	fmt.Scan(&N)

	var IsLeap bool
	if 1700 <= N && N < 1919 {
		if N%4 == 0 {
			IsLeap = true
		}
	} else {
		if N%400 == 0 || (N%4 == 0 && N%100 != 0) {
			IsLeap = true
		}
	}

	sum := 0
	for i := 0; i < 8; i++ {
		sum = sum + Month[i]
	}

	if IsLeap {
		sum = sum + 1
	}

	date := 0
	if N == 1918 {
		date = 256 - sum + 13
	} else {
		date = 256 - sum
	}

	fmt.Printf("%d.09.%d", date, N)
}
