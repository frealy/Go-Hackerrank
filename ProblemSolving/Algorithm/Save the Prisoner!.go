package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the saveThePrisoner function below.
func saveThePrisoner(n int64, m int64, s int64) int64 {
	var ans int64

	if (m+s-1)%n == 0 {
		ans = n
	} else {
		ans = (s + m - 1) % n
	}
	return ans
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024*3)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024*3)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int64(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nms := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nms[0], 10, 64)
		checkError(err)
		n := int64(nTemp)

		mTemp, err := strconv.ParseInt(nms[1], 10, 64)
		checkError(err)
		m := int64(mTemp)

		sTemp, err := strconv.ParseInt(nms[2], 10, 64)
		checkError(err)
		s := int64(sTemp)

		result := saveThePrisoner(n, m, s)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
