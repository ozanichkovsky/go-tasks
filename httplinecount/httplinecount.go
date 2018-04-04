package main

import (
	"net/http"
	"fmt"
	"log"
	"os"
	"io"
	"bufio"
)

var dir string

// CountLines - count the lines in from Reader
func CountLines(r io.Reader) (int, error) {
	sc := bufio.NewScanner(r)
	var lines int
	for sc.Scan() {
		lines++
	}
	return lines, sc.Err()
}

// CountFile write to stdout number of lines in the file
func CountFile(path string) int {
	f, _ := os.Open(path)

	defer f.Close()

	r := bufio.NewReader(f)
	l, _ := CountLines(r)

	return l
}

func handler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Path[1:]
	fmt.Fprintf(w, "%s line count is %d!", filename, CountFile(dir + filename))
}

func main() {
	dir = os.Args[1]

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
