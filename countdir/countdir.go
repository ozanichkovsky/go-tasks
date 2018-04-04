package main

import (
	"io"
	"os"
	"bufio"
	"fmt"
)

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
func CountFile(path string) {
	f, _ := os.Open(path)

	defer f.Close()

	r := bufio.NewReader(f)
	l, _ := CountLines(r)

	fmt.Println(path, l)
}

// CountDir - write to stdout numbers of lines for each files in directory
func CountDir(dir string) {
	d, err := os.Open(dir)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer d.Close()
	fi, err := d.Readdir(-1)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, fi := range fi {
		name := fi.Name()

		if string(name[len(name)-4:]) == ".txt" {
			CountFile(dir + fi.Name())
		}
	}
}

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		CountDir(arg)
	}
}
