package main

import "fmt"
import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var str string
	for scanner.Scan() {
		str = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	n := len(str)
	if n%2 == 1 {
		fmt.Println("False")
		return
	}

	var open, close int
	for i := 0; i < n; i++ {
		a := str[i]
		if a == 40 || a == 91 || a == 123 {
			open++
		}
		if a == 41 || a == 93 || a == 125 {
			close++
		}
	}

	if open == close {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
