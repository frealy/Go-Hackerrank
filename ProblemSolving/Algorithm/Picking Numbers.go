package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the pickingNumbers function below.
func pickingNumbers(a []int32) int32 {
	var val []int32
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	for i := 0; i < len(a); i++ {
		count := []int32{a[i]}
		for j := i + 1; j < len(a); j++ {
			if math.Abs(float64(a[i]-a[j])) <= 1 {
				count = append(count, a[j])
			}
		}
		val = append(val, int32(len(count)))
	}
	var value int32 = val[0]
	for _, e := range val {
		if e > value {
			value = e
		}
	}
	fmt.Println(value)
	return value

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	aTemp := strings.Split(readLine(reader), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := pickingNumbers(a)

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
