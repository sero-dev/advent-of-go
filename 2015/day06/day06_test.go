package day06

import (
	"testing"

	"serodev.com/advent-of-go/util"
)

func TestProblem1WhenInstructionTurnsOnAllLights(t *testing.T) {
	input := "turn on 0,0 through 999,999\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem1(inputString)
	expect := 1000000

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblem1WhenInstructionTurnsOnFirstLineOfLights(t *testing.T) {
	input := "toggle 0,0 through 999,0\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem1(inputString)
	expect := 1000

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblem1WhenInstructionTurnsOffMiddleLight(t *testing.T) {
	input := "turn on 0,0 through 999,999\nturn off 499,499 through 500,500\n"
	inputString := util.CreateBufferReaderFrom(input)

	actual, err := Problem1(inputString)
	expect := 1000000 - 4

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}
