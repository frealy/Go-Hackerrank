package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
func Min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func nonDivisibleSubset(k int64, S []int64) int64 {
	n := int64(len(S))
	count := int64(0)
	s := make([]int64, k)

	for i := int64(0); i < n; i++ {
		s[S[i]%k] += 1
	}

	count = Min(s[0], 1)
	for i := int64(1); i < k/2+1; i++ {
		if i != k-i {
			count += Max(s[i], s[k-i])
		}
	}
	if k%2 == 0 {
		count += 1
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int64(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int64(kTemp)

	STemp := strings.Split(readLine(reader), " ")

	var S []int64

	for i := 0; i < int(n); i++ {
		SItemTemp, err := strconv.ParseInt(STemp[i], 10, 64)
		checkError(err)
		SItem := int64(SItemTemp)
		S = append(S, SItem)
	}

	result := nonDivisibleSubset(k, S)

	fmt.Fprintf(writer, "%d\n", result)

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
