package day03

import (
	"testing"

	"serodev.com/advent-of-go/util"
)

func TestProblem1WhenSantaVisitsTwoHouse(t *testing.T) {
	inputString := ">"
	input := util.CreateBufferReaderFrom(inputString)

	actual, err := Problem1(input)
	expect := 2

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", inputString, actual, expect, err)
	}
}

func TestProblem1WhenSantaVisitsHouseTwice(t *testing.T) {
	inputString := "^>v<"
	input := util.CreateBufferReaderFrom(inputString)

	actual, err := Problem1(input)
	expect := 4

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", inputString, actual, expect, err)
	}
}

func TestProblem1WhenSantaVisitsHouseMultipleTimes(t *testing.T) {
	inputString := "^v^v^v^v^v"
	input := util.CreateBufferReaderFrom(inputString)

	actual, err := Problem1(input)
	expect := 2

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", inputString, actual, expect, err)
	}
}

func TestProblem2WhenSantaVisitsThreeHouse(t *testing.T) {
	inputString := "^v"
	input := util.CreateBufferReaderFrom(inputString)

	actual, err := Problem2(input)
	expect := 3

	if actual != expect || err != nil {
		t.Fatalf("Problem2() with %s input returned %d, instead of %d, with %v error", inputString, actual, expect, err)
	}
}

func TestProblem2WhenSantaVisitsHouseTwice(t *testing.T) {
	inputString := "^>v<"
	input := util.CreateBufferReaderFrom(inputString)

	actual, err := Problem2(input)
	expect := 3

	if actual != expect || err != nil {
		t.Fatalf("Problem2() with %s input returned %d, instead of %d, with %v error", inputString, actual, expect, err)
	}
}

func TestProblem2WhenSantaVisitsHouseMultipleTimes(t *testing.T) {
	inputString := "^v^v^v^v^v"
	input := util.CreateBufferReaderFrom(inputString)

	actual, err := Problem2(input)
	expect := 11

	if actual != expect || err != nil {
		t.Fatalf("Problem2() with %s input returned %d, instead of %d, with %v error", inputString, actual, expect, err)
	}
}