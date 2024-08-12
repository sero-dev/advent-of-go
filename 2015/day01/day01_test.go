package day01

import (
	"bufio"
	"bytes"
	"log"
	"testing"
)

func TestProblem1WhenInputIsBalanced(t *testing.T) {
	inputString := "(())"
	input := createBufferReaderFrom(inputString)

	actual, err := Problem1(input)
	expect := 0

	if actual != expect || err != nil {
		log.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", inputString, actual, expect, err)
	}
}

func TestProblem1WhenInputHasMoreOpeningParentheses(t *testing.T) {
	inputString := "))((((("
	input := createBufferReaderFrom(inputString)

	actual, err := Problem1(input)
	expect := 3

	if actual != expect || err != nil {
		log.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", inputString, actual, expect, err)
	}
}

func TestProblem1WhenInputHasMoreClosingParentheses(t *testing.T) {
	inputString := ")())())"
	input := createBufferReaderFrom(inputString)

	actual, err := Problem1(input)
	expect := -3

	if actual != expect || err != nil {
		log.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", inputString, actual, expect, err)
	}
}

func createBufferReaderFrom(input string) *bufio.Reader {
	var buffer bytes.Buffer
	buffer.WriteString(input)

	return bufio.NewReader(&buffer)
}