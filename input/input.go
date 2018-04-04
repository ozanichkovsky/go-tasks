package main

import (
	"io"
	"io/ioutil"
)

// ReadAll reads all the lines of text from r and returns
// all the data read as a string
func ReadAll(r io.Reader) string {
	if b, err := ioutil.ReadAll(r); err == nil {
		return string(b)
	}

	return ""
}
