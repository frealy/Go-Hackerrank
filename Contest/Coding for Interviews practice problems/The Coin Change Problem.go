package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the getWays function below.
func getWays(n int64, c []int64) int64 {
	var table []int64 = make([]int64, n+1)
	table[0] = 1
	m := int64(len(c))

	for i := int64(0); i < m; i++ {
		for j := c[i]; j <= n; j++ {
			table[j] += table[j-c[i]]
		}
	}

	return table[n]

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int64(nTemp)

	mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	checkError(err)
	m := int64(mTemp)

	cTemp := strings.Split(readLine(reader), " ")

	var c []int64

	for i := 0; i < int(m); i++ {
		cItem, err := strconv.ParseInt(cTemp[i], 10, 64)
		checkError(err)
		c = append(c, cItem)
	}

	// Print the number of ways of making change for 'n' units using coins having the values given by 'c'

	ways := getWays(n, c)

	fmt.Fprintf(writer, "%d\n", ways)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
