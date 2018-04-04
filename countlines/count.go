package count

import (
	"bytes"
	"io"
	"os"
)

// ContLines count lines of content in the file
func CountLines(path string) int {
	var lines int
	buf := make([]byte, 32*1024)
	lineSep := []byte{'\n'}

	f, _ := os.Open(path)

	defer f.Close()

	for {
		c, err := f.Read(buf)
		lines += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF: fallthrough
		case err != nil:
			return lines
		}
	}

	return lines
}
