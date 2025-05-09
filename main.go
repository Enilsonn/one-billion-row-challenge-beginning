package main

import (
	"fmt"
	"time"

	"github.com/Enilsonn/go-backend-starter/one-billion-row-challenge/elements"
)

func main() {
	start := time.Now()
	measurements, err := elements.GetMeasurements("measurements.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(elements.FormaterMeasurements(measurements))
	fmt.Println("Execution time:", time.Since(start))

}
