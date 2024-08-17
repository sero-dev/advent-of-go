package day04

import (
	"testing"
)

func TestProblemWhenSecretIsABCDEF(t *testing.T) {
	input := "abcdef"

	actual, err := Problem(input, "00000")
	var expect uint = 609043

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}

func TestProblemWhenSecretIsPQRSTUV(t *testing.T) {
	input := "pqrstuv"

	actual, err := Problem(input, "00000")
	var expect uint = 1048970

	if actual != expect || err != nil {
		t.Fatalf("Problem1() with %s input returned %d, instead of %d, with %v error", input, actual, expect, err)
	}
}
