package errorhandling

import (
	"os"
	"bytes"
	"io"
)

// CountLines count lines in the file
func CountLines(path string) (int, error) {
	var lines int
	buf := make([]byte, 32*1024)
	lineSep := []byte{'\n'}

	f, err := os.Open(path)

	defer f.Close()

	if err != nil {
		return 0, err
	}

	for {
		c, err := f.Read(buf)
		lines += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF: fallthrough
		case err != nil:
			return lines, nil
		}
	}

	return lines, nil
}
