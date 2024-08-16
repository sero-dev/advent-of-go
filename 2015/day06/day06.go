package day06

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

func Problem1(bufferReader *bufio.Reader) (int, error) {
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

		actionFunction := func(*bool) {}

		switch action {
		case "turn on":
			actionFunction = turnon
		case "turn off":
			actionFunction = turnoff
		case "toggle":
			actionFunction = toggle
		}

		performInstruction(&lightConfig, actionFunction, sourceX, sourceY, destX, destY)
		fmt.Printf("Lights that are on: %d\n", getLightsOnCount(lightConfig))
	}

	return getLightsOnCount(lightConfig), nil
}

type point struct {
	x int
	y int
}

func createLightConfiguration(length uint, width uint) [][]bool {
	lightConfig := make([][]bool, length)

	for i := range lightConfig {
		lightConfig[i] = make([]bool, width)

		for j := range lightConfig[i] {
			lightConfig[i][j] = false
		}
	}

	return lightConfig
}

func performInstruction(config *[][]bool, action func(*bool), sourceX int, sourceY int, destX int, destY int) {
	sourcePoint, destPoint := normalizeCoordinates(sourceX, sourceY, destX, destY)

	for y := sourcePoint.y; y <= destPoint.y; y++ {
		for x := sourcePoint.x; x <= destPoint.x; x++ {
			action(&(*config)[y][x])
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

	if sourceX > destY {
		sourcePoint.y = destY
		destPoint.y = sourceY
	}

	return sourcePoint, destPoint
}

func turnon(cell *bool) {
	*cell = true
}

func turnoff(cell *bool) {
	*cell = false
}

func toggle(cell *bool) {
	*cell = !*cell
}

func getLightsOnCount(config [][]bool) int {
	count := 0

	for y := range config {
		for x := range config[y] {
			if config[y][x] {
				count++
			}
		}
	}

	return count
}
