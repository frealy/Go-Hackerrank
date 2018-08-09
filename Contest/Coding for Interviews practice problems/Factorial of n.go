package main

import "fmt"
import (
	"bufio"
	"math/big"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var n int64
	for scanner.Scan() {
		n, _ = strconv.ParseInt(scanner.Text(), 10, 64)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	sum := new(big.Int).SetInt64(1)
	for n > 0 {
		x := new(big.Int).SetInt64(n)
		sum = sum.Mul(sum, x)
		n--
	}

	fmt.Println(sum)
}
