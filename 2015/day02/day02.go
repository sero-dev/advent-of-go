package day02

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strconv"
)

func Problem1(bufferReader *bufio.Reader) (int, error) {

	totalArea := 0

	for {
		line, err := bufferReader.ReadString('\n')

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return 0, err
		}

		dimensionPattern := regexp.MustCompile(`(\d+)x(\d+)x(\d+)`)

		matches := dimensionPattern.FindStringSubmatch(line)

		if (len(matches) != 4) {
			return 0, errors.New("string " + line + " was not formatted correctly")
		}

		length, _ := strconv.Atoi(matches[1])
		width, _ := strconv.Atoi(matches[2])
		height, _ := strconv.Atoi(matches[3])

		smallestSide := min(length * width, width * height, height * length)

		totalArea += getSurfaceArea(length, width, height) + smallestSide
	}

	return totalArea, nil
}

func Problem2(bufferReader *bufio.Reader) (int, error) {
	totalRibbonNeeded := 0

	for {
		line, err := bufferReader.ReadString('\n')

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return 0, err
		}

		dimensionPattern := regexp.MustCompile(`(\d+)x(\d+)x(\d+)`)

		matches := dimensionPattern.FindStringSubmatch(line)

		if (len(matches) != 4) {
			return 0, errors.New("string " + line + " was not formatted correctly")
		}

		length, _ := strconv.Atoi(matches[1])
		width, _ := strconv.Atoi(matches[2])
		height, _ := strconv.Atoi(matches[3])

		smallestSide := min(2 * (length + width), 2 * (width + height), 2 * (height + length))

		totalRibbonNeeded += length * width * height + smallestSide
	}

	return totalRibbonNeeded, nil
}

func getSurfaceArea(length int, width int, height int) int {
	return 2 * length * width + 2 * width * height + 2 * height * length
} 