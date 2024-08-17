package day02

import (
	"testing"

	"serodev.com/advent-of-go/util"
)

func TestProblem1WhenInputIs2x3x4(t *testing.T) {
	inputString := "2x3x4\n"
	input := util.CreateBufferReaderFrom(inputString)

	actual, err := Problem1(input)
	expect := 58

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", inputString, actual, expect, err)
	}
}

func TestProblem1WhenInputIs1x1x10(t *testing.T) {
	inputString := "1x1x10\n"
	input := util.CreateBufferReaderFrom(inputString)

	actual, err := Problem1(input)
	expect := 43

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", inputString, actual, expect, err)
	}
}

func TestProblem2WhenInputIs2x3x4(t *testing.T) {
	inputString := "2x3x4\n"
	input := util.CreateBufferReaderFrom(inputString)

	actual, err := Problem2(input)
	expect := 34

	if actual != expect || err != nil {
		t.Fatalf("Problem2() with %s input returned %d, instead of %d, with %v error", inputString, actual, expect, err)
	}
}

func TestProblem2WhenInputIs1x1x10(t *testing.T) {
	inputString := "1x1x10\n"
	input := util.CreateBufferReaderFrom(inputString)

	actual, err := Problem2(input)
	expect := 14

	if actual != expect || err != nil {
		t.Fatalf("Problem2() with %s input returned %d, instead of %d, with %v error", inputString, actual, expect, err)
	}
}
