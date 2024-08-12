package day01

import (
	"bufio"
	"errors"
	"io"
)

func Problem1(bufferReader *bufio.Reader) (int, error) {
	count := 0

	for {
		b, err := bufferReader.ReadByte()

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return 0, err
		}

		if b == 40 {
			count++
		} else if b == 41 {
			count--
		}
	}

	return count, nil
}