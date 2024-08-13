package day03

import (
	"bufio"
	"errors"
	"io"
)

type point struct {
	x int
	y int
}

func Problem1(bufferReader *bufio.Reader) (int, error) {
	x := 0
	y := 0

	coordinates := make(map[point] bool)
	coordinates[point{x: x, y: y}] = true

	for {
		b, err := bufferReader.ReadByte()

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return 0, err
		}

		if b == 94 {
			y += 1
		} else if b == 62 {
			x += 1
		} else if b == 118 {
			y -= 1
		} else if b == 60 {
			x -= 1
		}

		coordinates[point{x: x, y: y}] = true
	}

	return len(coordinates), nil
}

func Problem2(bufferReader *bufio.Reader) (int, error) {
	santaX := 0
	santaY := 0

	robotX := 0
	robotY := 0

	isSantasTurn := true
	coordinates := make(map[point] bool)
	coordinates[point{x: santaX, y: santaY}] = true

	for {
		b, err := bufferReader.ReadByte()

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return 0, err
		}

		if b == 94 {
			if isSantasTurn {
				santaY++
			} else {
				robotY++
			}
		} else if b == 62 {
			if isSantasTurn {
				santaX++
			} else {
				robotX++
			}
		} else if b == 118 {
			if isSantasTurn {
				santaY--
			} else {
				robotY--
			}
		} else if b == 60 {
			if isSantasTurn {
				santaX--
			} else {
				robotX--
			}
		}
 
		if (isSantasTurn) {
			coordinates[point{x: santaX, y: santaY}] = true
		} else {
			coordinates[point{x: robotX, y: robotY}] = true
		}

		isSantasTurn = !isSantasTurn
	}

	return len(coordinates), nil
}