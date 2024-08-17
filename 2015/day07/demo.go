package day07

import (
	"bufio"
	"os"
)

func DemoProblem1() (uint16, error) {
	file, err := os.Open("2015/day07/input.txt")

	if err != nil {
		return 0, err
	}

	bufferReader := bufio.NewReader(file)
	answer, err := Problem1(bufferReader, "a")

	if err != nil {
		return 0, err
	}

	return answer, nil
}

func DemoProblem2() (uint16, error) {
	file, err := os.Open("2015/day07/input.txt")

	if err != nil {
		return 0, err
	}

	bufferReader := bufio.NewReader(file)
	answer, err := Problem2(bufferReader, "a")

	if err != nil {
		return 0, err
	}

	return answer, nil
}
