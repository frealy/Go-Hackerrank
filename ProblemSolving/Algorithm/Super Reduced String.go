package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the superReducedString function below.
func superReducedString(s string) string {
	N := len(s)
	for i := 0; i < N-1; {
		if s[i] == s[i+1] {
			s = strings.Replace(s, string(s[i])+string(s[i+1]), "", -1)
			N = len(s)
			i = 0
		} else {
			i++
		}
		fmt.Println(s, N, i)
	}

	if s == "" {
		s = "Empty String"
	}
	return s

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := superReducedString(s)

	fmt.Fprintf(writer, "%s\n", result)

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
