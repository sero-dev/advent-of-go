package day05

import (
	"bufio"
	"errors"
	"io"
)

func Problem1(bufferReader *bufio.Reader) (int, error) {
	niceCount := 0

	for {
		line, err := bufferReader.ReadString('\n')

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return 0, err
		}

		if isStringNice(line) {
			niceCount++
		}
	}

	return niceCount, nil
}

func Problem2(bufferReader *bufio.Reader) (int, error) {
	niceCount := 0

	for {
		line, err := bufferReader.ReadString('\n')

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return 0, err
		}

		if isStringReallyNice(line) {
			niceCount++
		}
	}

	return niceCount, nil
}

func isStringNice(str string) bool {
	vowelCount := 0
	hasDoubleLetters := false

	for i := 0; i < len(str)-1; i++ {
		if isAVowel(str[i]) {
			vowelCount++
		}

		if i+1 == len(str) {
			break
		}

		if str[i] == str[i+1] {
			hasDoubleLetters = true
		}

		if isForbiddenSequeue(str[i], str[i+1]) {
			return false
		}
	}

	return hasDoubleLetters && vowelCount >= 3
}

func isStringReallyNice(str string) bool {
	hasRepeatingPair := false
	hasDoubleLetters := false

	pairMap := make(map[string]int)

	if len(str) < 4 {
		return false
	}

	for i := 0; i < len(str)-1; i++ {

		if i > len(str)-3 && !hasDoubleLetters {
			return false
		}

		if i > len(str)-2 && !hasRepeatingPair {
			return false
		}

		if !hasRepeatingPair {
			key := str[i : i+2]
			index, exists := pairMap[key]

			if exists {
				if index != i-1 {
					hasRepeatingPair = true
				}
			} else {
				pairMap[key] = i
			}
		}

		if !hasDoubleLetters && str[i] == str[i+2] {
			hasDoubleLetters = true
		}

		if hasDoubleLetters && hasRepeatingPair {
			break
		}
	}

	return hasDoubleLetters && hasRepeatingPair
}

func isForbiddenSequeue(first byte, second byte) bool {
	forbiddenSequence := []string{
		"ab",
		"cd",
		"pq",
		"xy",
	}

	for _, sequence := range forbiddenSequence {
		if first == sequence[0] && second == sequence[1] {
			return true
		}
	}

	return false
}

func isAVowel(letter byte) bool {
	switch letter {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}

	return false
}
