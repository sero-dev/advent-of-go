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

func Problem2(bufferReader *bufio.Reader) (int, error) {
	count := 0
	index := 1

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

		if count < 0 {
			return index, nil
		}

		index++
	}

	return -1, errors.New("count was never negative")
}
