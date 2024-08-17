package main

import (
	"fmt"
	"log"

	"serodev.com/advent-of-go/2015/day06"
)

func main() {
	answer, err := day06.DemoProblem2()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(answer)
}
