package main

import (
	"fmt"

	"github.com/Enilsonn/go-backend-starter/one-billion-row-challenge/elements"
)

func main() {
	measurements, err := elements.GetMeasurements("measurements.txt")
	if err != nil {
		panic(err)
	}

	for nome, m := range measurements {
		fmt.Printf("%s => min=%.2f, max=%.2f, avg=%.2f, count=%d\n", nome, m.Min, m.Max, m.Sum/float64(m.Times), m.Times)
	}

}
