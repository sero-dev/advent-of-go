package day07

import (
	"testing"

	"serodev.com/advent-of-go/util"
)

func TestProblem1WhenUsingDirectSignal(t *testing.T) {
	input := "123 -> x\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem1(inputString, "x")
	var expect uint16 = 123

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblem1WhenUsingAndGate(t *testing.T) {
	input := "123 -> x\n456 -> y\nx AND y -> d\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem1(inputString, "d")
	var expect uint16 = 72

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblem1WhenUsingOrGate(t *testing.T) {
	input := "123 -> x\n456 -> y\nx OR y -> e\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem1(inputString, "e")
	var expect uint16 = 507

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblem1WhenUsingLShiftGate(t *testing.T) {
	input := "123 -> x\nx LSHIFT 2 -> f\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem1(inputString, "f")
	var expect uint16 = 492

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblem1WhenUsingRShiftGate(t *testing.T) {
	input := "456 -> y\ny RSHIFT 2 -> g\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem1(inputString, "g")
	var expect uint16 = 114

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblem1WhenUsingNotGate(t *testing.T) {
	input := "456 -> y\nNOT y -> i\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem1(inputString, "i")
	var expect uint16 = 65079

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}
