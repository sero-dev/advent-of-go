package day06

import (
	"bufio"
	"os"
)

func DemoProblem1() (int, error) {
	file, err := os.Open("2015/day06/input.txt")

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
