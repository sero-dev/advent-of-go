package main

import (
	"serodev.com/advent-of-go/2015/day01"
	"serodev.com/advent-of-go/2015/day02"
	"serodev.com/advent-of-go/2015/day03"
	"serodev.com/advent-of-go/2015/day04"
	"serodev.com/advent-of-go/2015/day05"
	"serodev.com/advent-of-go/2015/day06"
	"serodev.com/advent-of-go/2015/day07"
	"serodev.com/advent-of-go/performance"
)

func main() {
	runBenchmarks()
}

func runBenchmarks() {
	performance.RunBenchmark("Day 01 - Problem 1", func() { day01.DemoProblem1() }, 10000)
	performance.RunBenchmark("Day 01 - Problem 2", func() { day01.DemoProblem2() }, 10000)
	performance.RunBenchmark("Day 02 - Problem 1", func() { day02.DemoProblem1() }, 100)
	performance.RunBenchmark("Day 02 - Problem 2", func() { day02.DemoProblem2() }, 100)
	performance.RunBenchmark("Day 03 - Problem 1", func() { day03.DemoProblem1() }, 1000)
	performance.RunBenchmark("Day 03 - Problem 2", func() { day03.DemoProblem2() }, 1000)
	performance.RunBenchmark("Day 04 - Problem 1", func() { day04.DemoProblem1() }, 10)
	performance.RunBenchmark("Day 04 - Problem 2", func() { day04.DemoProblem2() }, 1)
	performance.RunBenchmark("Day 05 - Problem 1", func() { day05.DemoProblem1() }, 1000)
	performance.RunBenchmark("Day 05 - Problem 2", func() { day05.DemoProblem2() }, 1000)
	performance.RunBenchmark("Day 06 - Problem 1", func() { day06.DemoProblem1() }, 1)
	performance.RunBenchmark("Day 06 - Problem 2", func() { day06.DemoProblem2() }, 1)
	performance.RunBenchmark("Day 07 - Problem 1", func() { day07.DemoProblem1() }, 1000)
	performance.RunBenchmark("Day 07 - Problem 2", func() { day07.DemoProblem2() }, 1000)
}
