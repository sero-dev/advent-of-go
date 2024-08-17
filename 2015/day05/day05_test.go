package day05

import (
	"testing"

	"serodev.com/advent-of-go/util"
)

func TestProblem1WhenShortInputIsNice(t *testing.T) {
	input := "aaa\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem1(inputString)
	expect := 1

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblem1WhenLongInputIsNice(t *testing.T) {
	input := "ugknbfddgicrmopn\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem1(inputString)
	expect := 1

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblem1WhenInputContainsNoDoubleLetter(t *testing.T) {
	input := "jchzalrnumimnmhp\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem1(inputString)
	expect := 0

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblem1WhenInputContainsAB(t *testing.T) {
	input := "haegwjzuvuyypabu\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem1(inputString)
	expect := 0

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblem1WhenInputContainsCD(t *testing.T) {
	input := "haegwjzuvuyycdyu\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem1(inputString)
	expect := 0

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblem1WhenInputContainsPQ(t *testing.T) {
	input := "haegwjzuvuyyppqu\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem1(inputString)
	expect := 0

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblem1WhenInputContainsXY(t *testing.T) {
	input := "haegwjzuvuyypxyu\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem1(inputString)
	expect := 0

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblem1WhenInputHasLessThanThreeVowels(t *testing.T) {
	input := "dvszwmarrgswjxmb\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem1(inputString)
	expect := 0

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblem2WhenInputIsNice(t *testing.T) {
	input := "qjhvhtzxzqqjkmpb\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem2(inputString)
	expect := 1

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblem2WhenInputIsNiceWithOverlapping(t *testing.T) {
	input := "xxyxx\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem2(inputString)
	expect := 1

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblem2WhenInputHasNoDoubleLettersWithLetterInTheMiddle(t *testing.T) {
	input := "uurcxstgmygtbstg\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem2(inputString)
	expect := 0

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblem2WhenInputHasNoPairRepeatingTwice(t *testing.T) {
	input := "ieodomkazucvgmuy\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem2(inputString)
	expect := 0

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}
