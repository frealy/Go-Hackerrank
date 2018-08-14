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

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Complete the beautifulDays function below.
func beautifulDays(i int64, j int64, k int64) int64 {
	var count int64
	for i <= j {
		hStr := fmt.Sprint(i)
		hStr = Reverse(hStr)
		h, _ := strconv.ParseInt(hStr, 10, 64)
		if int64(math.Abs(float64(i)-float64(h)))%k == 0 {
			count++
		}
		i++
	}
	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	ijk := strings.Split(readLine(reader), " ")

	iTemp, err := strconv.ParseInt(ijk[0], 10, 64)
	checkError(err)
	i := int64(iTemp)

	jTemp, err := strconv.ParseInt(ijk[1], 10, 64)
	checkError(err)
	j := int64(jTemp)

	kTemp, err := strconv.ParseInt(ijk[2], 10, 64)
	checkError(err)
	k := int64(kTemp)

	result := beautifulDays(i, j, k)

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
