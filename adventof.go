package main

import (
	"serodev.com/advent-of-go/2015/day01"
	"serodev.com/advent-of-go/performance"
)

func main() {
	performance.RunBenchmark(func() {
		day01.DemoProblem1()
	}, 100000)
}