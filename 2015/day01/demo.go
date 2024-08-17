package day01

import (
	"bufio"
	"os"
)

func DemoProblem1() (int, error) {
	file, err := os.Open("2015/day01/input.txt")

	if err != nil {
		return 0, err
	}

	bufferReader := bufio.NewReader(file)
	answer, err := Problem1(bufferReader)

	if err != nil {
		return 0, err
	}

	return answer, nil
}

func DemoProblem2() (int, error) {
	file, err := os.Open("2015/day01/input.txt")

	if err != nil {
		return 0, err
	}

	bufferReader := bufio.NewReader(file)
	answer, err := Problem2(bufferReader)

	if err != nil {
		return 0, err
	}

	return answer, nil
}
