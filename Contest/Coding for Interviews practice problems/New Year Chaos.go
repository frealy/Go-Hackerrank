package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int64) {
	n := int64(len(q))
	bribe := int64(0)
	for i := n - 1; i >= 0; i-- {
		m := i + 1
		if q[i]-m > 2 {
			fmt.Println("Too chaotic")
			return
		}

		for j := int64(math.Max(0, float64(q[i]-2))); j < i; j++ {
			if q[j] > q[i] {
				bribe++
			}
		}
	}
	fmt.Println(bribe)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int64(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int64(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int64

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int64(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
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
