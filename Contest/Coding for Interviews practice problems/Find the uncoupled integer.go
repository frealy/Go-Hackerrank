package main

import "fmt"
import (
	"bufio"
	"os"
	"strings"
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

	N := strings.Split(str, ", ")
	n := len(N)
	m := make(map[string]int)
	for i := 0; i < n; i++ {
		if m[string(N[i])] == 0 {
			m[string(N[i])] = 1
		} else if m[string(N[i])] == 1 {
			delete(m, string(N[i]))
		}
	}

	for val, _ := range m {
		fmt.Println(val)
	}
}
