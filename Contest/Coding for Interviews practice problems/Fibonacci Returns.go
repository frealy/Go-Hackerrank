package main

import "fmt"
import (
	"bufio"
	"math/big"
	"os"
	"strconv"
)

func fib(table []*big.Int, N int64) *big.Int {
	var x big.Int
	if table[N] != nil {
		if table[N].Cmp(&x) == 1 {
			return table[N]
		}
	}
	for i := int64(2); i <= N; i++ {
		table[i] = new(big.Int).Add(table[i-1], table[i-2])
	}

	return table[N]
}

func main() {
	var val []*big.Int
	val = make([]*big.Int, 100+1)
	val[0] = new(big.Int).SetInt64(0)
	val[1] = new(big.Int).SetInt64(1)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		temp, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		fmt.Println(fib(val, temp))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
