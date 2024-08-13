package util

import (
	"bufio"
	"bytes"
)

func CreateBufferReaderFrom(input string) *bufio.Reader {
	var buffer bytes.Buffer
	buffer.WriteString(input)

	return bufio.NewReader(&buffer)
}