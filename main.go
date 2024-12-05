package main

import (
	"fmt"
	"math"

	"github.com/gavgrego/advent-2024-go/config"
)

func main() {
	var total int
	for i := 0; i < len(config.FirstList); i++ {
		if config.FirstList[i] != config.SecondList[i] {
			difference := math.Abs(float64(config.FirstList[i] - config.SecondList[i]))
			total += int(difference)
		}
	}

	fmt.Println("Total", total)
}


