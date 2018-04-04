package main

import (
	"io"
	"os"
	"bufio"
	"fmt"
)

// CountLines count lines in the reader
func CountLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	var lines int
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

// CountFile write count of lines in the file to the stdout
func CountFile(path string) {
	f, _ := os.Open(path)

	defer f.Close()

	r := bufio.NewReader(f)
	l, _ := CountLines(r)

	fmt.Println(path, l)
}

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		CountFile(arg)
	}
}
