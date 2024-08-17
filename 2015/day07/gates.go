package day07

import (
	"errors"
	"strconv"
)

type Gate interface {
	Process(*map[string]interface{}) (uint16, error)
}

type AndGate struct {
	wire1Name string
	wire2Name string
}

func (g AndGate) Process(wires *map[string]interface{}) (uint16, error) {

	wire1Int, err := strconv.Atoi(g.wire1Name)
	wire1Value := uint16(wire1Int)

	if err != nil {
		value, err := FindSignal(wires, g.wire1Name)
		wire1Value = value

		if err != nil {
			return 0, err
		}
	}

	wire2Signal, err := FindSignal(wires, g.wire2Name)

	if err != nil {
		return 0, err
	}

	return wire1Value & wire2Signal, nil
}

type OrGate struct {
	wire1Name string
	wire2Name string
}

func (g OrGate) Process(wires *map[string]interface{}) (uint16, error) {

	wire1Signal, err := FindSignal(wires, g.wire1Name)

	if err != nil {
		return 0, err
	}

	wire2Signal, err := FindSignal(wires, g.wire2Name)

	if err != nil {
		return 0, err
	}

	return wire1Signal | wire2Signal, nil
}

type ShiftGate struct {
	wireName     string
	shiftBy      int
	isRightShift bool
}

func (g ShiftGate) Process(wires *map[string]interface{}) (uint16, error) {

	wireSignal, err := FindSignal(wires, g.wireName)

	if err != nil {
		return 0, err
	}

	if g.isRightShift {
		return wireSignal >> g.shiftBy, nil
	} else {
		return wireSignal << g.shiftBy, nil
	}
}

type NotGate struct {
	wireName string
}

func (g NotGate) Process(wires *map[string]interface{}) (uint16, error) {

	wireSignal, err := FindSignal(wires, g.wireName)

	if err != nil {
		return 0, err
	}

	return ^wireSignal, nil
}

func FindSignal(wires *map[string]interface{}, wireName string) (uint16, error) {
	wire, exists := (*wires)[wireName]

	if !exists {
		return 0, errors.New("signal cannot be found")
	}

	var signal uint16
	var err error

	switch value := wire.(type) {
	case uint16:
		signal = value
	case string:
		signal, err = FindSignal(wires, value)
	case Gate:
		signal, err = value.Process(wires)
	default:
		err = errors.New("unsupported wire type")
	}

	if err != nil {
		return 0, err
	}

	(*wires)[wireName] = signal
	return signal, nil
}
