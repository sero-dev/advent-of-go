package performance

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type benchmarkSummary struct {
	slowest time.Duration
	fastest time.Duration
	average time.Duration
}

func RunBenchmark(name string, subject func(), n uint) {
	benchmarks := make([]time.Duration, n)

	overallStartTime := time.Now()

	for index := range benchmarks {
		startTime := time.Now()
		subject()
		benchmarks[index] = time.Since(startTime)
	}

	summary, err := getSummary(benchmarks)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Number of runs: %d\n", n)
	fmt.Printf("Overall time: %v\n", time.Since(overallStartTime))
	fmt.Printf("Slowest run: %v\n", summary.slowest)
	fmt.Printf("Fastest run: %v\n", summary.fastest)
	fmt.Printf("Average run: %v\n", summary.average)
	fmt.Println()
}

func getSummary(benchmarks []time.Duration) (benchmarkSummary, error) {
	slowestRun := time.Since(time.Now())
	fastestRun := time.Since(time.Now())

	totalTime := 0

	if len(benchmarks) == 0 {
		return benchmarkSummary{
			slowest: 0,
			fastest: 0,
			average: 0,
		}, errors.New("benchmark array is empty")
	}

	for _, run := range benchmarks {
		if run > slowestRun {
			slowestRun = run
		}

		if run < fastestRun {
			fastestRun = run
		}

		totalTime += int(run.Nanoseconds())
	}

	averageRun := totalTime / len(benchmarks)

	return benchmarkSummary{
		fastest: fastestRun,
		slowest: slowestRun,
		average: time.Duration(averageRun),
	}, nil
}
