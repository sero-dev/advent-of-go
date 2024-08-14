package day05

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

func Problem1(bufferReader *bufio.Reader) (int, error) {
	niceCount := 0

	for {
		vowelCount := 0
		hasDoubleLetters := false

		line, err := bufferReader.ReadString('\n')

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return 0, err
		}

		for index, letter := range line {
			if isAVowel(letter) {
				vowelCount++
			}

			if index + 1 == len(line) {
				break
			}
		}
		
	}


	return niceCount, nil
}

func isAVowel(letter rune) bool {
	switch letter {
	case 'a':
	case 'e':
	case 'i':
	case 'o':
	case 'u':
		return true
	default:
		return false
	}
}