package day07

import (
	"bufio"
	"errors"
	"io"
	"maps"
	"regexp"
	"strconv"
	"strings"
)

func Problem1(bufferReader *bufio.Reader, wireName string) (uint16, error) {
	wires := make(map[string]interface{})

	for {
		line, err := bufferReader.ReadString('\n')

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return 0, err
		}

		if err := processString(&wires, line); err != nil {
			return 0, err
		}
	}

	return FindSignal(&wires, wireName)
}

func Problem2(bufferReader *bufio.Reader, wireName string) (uint16, error) {
	wires := make(map[string]interface{})

	for {
		line, err := bufferReader.ReadString('\n')

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return 0, err
		}

		if err := processString(&wires, line); err != nil {
			return 0, err
		}
	}

	wires2 := maps.Clone(wires)
	firstValue, err := FindSignal(&wires, wireName)

	if err != nil {
		return 0, nil
	}

	wires2["b"] = firstValue

	return FindSignal(&wires2, wireName)
}

func processString(wires *map[string]interface{}, line string) error {
	instructionPattern := regexp.MustCompile(`(.+) -> (\w+)`)

	matches := instructionPattern.FindStringSubmatch(line)

	if len(matches) != 3 {
		return errors.New("string " + line + " was not formatted correctly")
	}

	instruction := matches[1]
	wireName := matches[2]

	if strings.Contains(instruction, "AND") {
		gate, err := ParseAndGate(line)
		if err != nil {
			return err
		}
		(*wires)[wireName] = gate
	} else if strings.Contains(instruction, "OR") {
		gate, err := ParseOrGate(line)
		if err != nil {
			return err
		}
		(*wires)[wireName] = gate
	} else if strings.Contains(instruction, "LSHIFT") {
		gate, err := ParseLShiftGate(line)
		if err != nil {
			return err
		}
		(*wires)[wireName] = gate
	} else if strings.Contains(instruction, "RSHIFT") {
		gate, err := ParseRShiftGate(line)
		if err != nil {
			return err
		}
		(*wires)[wireName] = gate
	} else if strings.Contains(instruction, "NOT") {
		gate, err := ParseNotGate(line)
		if err != nil {
			return err
		}
		(*wires)[wireName] = gate
	} else {
		signal, err := strconv.Atoi(instruction)
		if err != nil {
			(*wires)[wireName] = instruction
		} else {
			(*wires)[wireName] = uint16(signal)
		}
	}

	return nil
}

func ParseAndGate(instruction string) (Gate, error) {
	instructionPattern := regexp.MustCompile(`(\w+) AND (\w+)`)

	matches := instructionPattern.FindStringSubmatch(instruction)

	if len(matches) != 3 {
		return AndGate{}, instructionMalformedError(instruction)
	}

	return AndGate{
		wire1Name: matches[1],
		wire2Name: matches[2],
	}, nil
}

func ParseOrGate(instruction string) (Gate, error) {
	instructionPattern := regexp.MustCompile(`(\w+) OR (\w+)`)

	matches := instructionPattern.FindStringSubmatch(instruction)

	if len(matches) != 3 {
		return OrGate{}, instructionMalformedError(instruction)
	}

	return OrGate{
		wire1Name: matches[1],
		wire2Name: matches[2],
	}, nil
}

func ParseLShiftGate(instruction string) (Gate, error) {
	instructionPattern := regexp.MustCompile(`(\w+) LSHIFT (\d+)`)

	matches := instructionPattern.FindStringSubmatch(instruction)

	if len(matches) != 3 {
		return ShiftGate{}, instructionMalformedError(instruction)
	}

	shiftBy, err := strconv.Atoi(matches[2])

	if err != nil {
		return ShiftGate{}, instructionMalformedError(instruction)
	}

	return ShiftGate{
		wireName:     matches[1],
		shiftBy:      shiftBy,
		isRightShift: false,
	}, nil
}

func ParseRShiftGate(instruction string) (Gate, error) {
	instructionPattern := regexp.MustCompile(`(\w+) RSHIFT (\d+)`)

	matches := instructionPattern.FindStringSubmatch(instruction)

	if len(matches) != 3 {
		return ShiftGate{}, instructionMalformedError(instruction)
	}

	shiftBy, err := strconv.Atoi(matches[2])

	if err != nil {
		return ShiftGate{}, instructionMalformedError(instruction)
	}

	return ShiftGate{
		wireName:     matches[1],
		shiftBy:      shiftBy,
		isRightShift: true,
	}, nil
}

func ParseNotGate(instruction string) (Gate, error) {
	instructionPattern := regexp.MustCompile(`NOT (\w+)`)

	matches := instructionPattern.FindStringSubmatch(instruction)

	if len(matches) != 2 {
		return NotGate{}, instructionMalformedError(instruction)
	}

	return NotGate{
		wireName: matches[1],
	}, nil
}

func instructionMalformedError(instruction string) error {
	return errors.New("instruction " + instruction + " was not formatted correctly")
}
