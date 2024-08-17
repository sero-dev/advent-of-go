package day06

import (
	"bufio"
	"errors"
	"io"
	"regexp"
	"strconv"
)

func Problem1(bufferReader *bufio.Reader) (uint, error) {
	turnon := func(cell *uint) { *cell = 1 }
	turnoff := func(cell *uint) { *cell = 0 }
	toggle := func(cell *uint) {
		if *cell == 1 {
			*cell = 0
		} else {
			*cell = 1
		}
	}

	return runInstructionSet(bufferReader, turnon, turnoff, toggle)
}

func Problem2(bufferReader *bufio.Reader) (uint, error) {
	turnon := func(cell *uint) { *cell += 1 }
	turnoff := func(cell *uint) {
		if *cell > 0 {
			*cell--
		}
	}
	toggle := func(cell *uint) {
		*cell += 2
	}

	return runInstructionSet(bufferReader, turnon, turnoff, toggle)
}

type point struct {
	x int
	y int
}

func runInstructionSet(bufferReader *bufio.Reader, turnon func(*uint), turnoff func(*uint), toggle func(*uint)) (uint, error) {
	lightConfig := createLightConfiguration(1000, 1000)

	for {
		line, err := bufferReader.ReadString('\n')

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return 0, err
		}

		instructionPattern := regexp.MustCompile(`(.+) (\d+),(\d+) through (\d+),(\d+)`)

		matches := instructionPattern.FindStringSubmatch(line)

		if len(matches) != 6 {
			return 0, errors.New("string " + line + " was not formatted correctly")
		}

		action := matches[1]
		sourceX, _ := strconv.Atoi(matches[2])
		sourceY, _ := strconv.Atoi(matches[3])

		destX, _ := strconv.Atoi(matches[4])
		destY, _ := strconv.Atoi(matches[5])

		actionFunction := func(*uint) {}

		switch action {
		case "turn on":
			actionFunction = turnon
		case "turn off":
			actionFunction = turnoff
		case "toggle":
			actionFunction = toggle
		}

		performInstruction(&lightConfig, actionFunction, sourceX, sourceY, destX, destY)
	}

	return getLightsOnCount(lightConfig), nil
}

func createLightConfiguration(length uint, width uint) [][]uint {
	lightConfig := make([][]uint, length)

	for i := range lightConfig {
		lightConfig[i] = make([]uint, width)

		for j := range lightConfig[i] {
			lightConfig[i][j] = 0
		}
	}

	return lightConfig
}

func performInstruction(config *[][]uint, instruction func(*uint), sourceX int, sourceY int, destX int, destY int) {
	sourcePoint, destPoint := normalizeCoordinates(sourceX, sourceY, destX, destY)

	for y := sourcePoint.y; y <= destPoint.y; y++ {
		for x := sourcePoint.x; x <= destPoint.x; x++ {
			instruction(&(*config)[y][x])
		}
	}
}

func normalizeCoordinates(sourceX int, sourceY int, destX int, destY int) (point, point) {
	sourcePoint := point{x: sourceX, y: sourceY}
	destPoint := point{x: destX, y: destY}

	if sourceX > destX {
		sourcePoint.x = destX
		destPoint.x = sourceX
	}

	if sourceY > destY {
		sourcePoint.y = destY
		destPoint.y = sourceY
	}

	return sourcePoint, destPoint
}

func getLightsOnCount(config [][]uint) uint {
	var count uint = 0

	for y := range config {
		for x := range config[y] {
			count += config[y][x]
		}
	}

	return count
}
