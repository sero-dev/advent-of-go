package day01

import (
	"bufio"
	"bytes"
	"testing"
)

func TestProblem1WhenInputIsBalanced(t *testing.T) {
	inputString := "(())"
	input := createBufferReaderFrom(inputString)

	actual, err := Problem1(input)
	expect := 0

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", inputString, actual, expect, err)
	}
}

func TestProblem1WhenInputHasMoreOpeningParentheses(t *testing.T) {
	inputString := "))((((("
	input := createBufferReaderFrom(inputString)

	actual, err := Problem1(input)
	expect := 3

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", inputString, actual, expect, err)
	}
}

func TestProblem1WhenInputHasMoreClosingParentheses(t *testing.T) {
	inputString := ")())())"
	input := createBufferReaderFrom(inputString)

	actual, err := Problem1(input)
	expect := -3

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", inputString, actual, expect, err)
	}
}

func TestProblem2WhenCountIsNegativeAtFirstPosition(t *testing.T) {
	inputString := ")"
	input := createBufferReaderFrom(inputString)

	actual, err := Problem2(input)
	expect := 1

	if actual != expect || err != nil {
		t.Fatalf("Problem2() with %s input returned %d, instead of %d, with %v error", inputString, actual, expect, err)
	}
}

func TestProblem2WhenCountIsNegativeAtFifthPosition(t *testing.T) {
	inputString := "()())"
	input := createBufferReaderFrom(inputString)

	actual, err := Problem2(input)
	expect := 5

	if actual != expect || err != nil {
		t.Fatalf("Problem2() with %s input returned %d, instead of %d, with %v error", inputString, actual, expect, err)
	}
}

func createBufferReaderFrom(input string) *bufio.Reader {
	var buffer bytes.Buffer
	buffer.WriteString(input)

	return bufio.NewReader(&buffer)
}