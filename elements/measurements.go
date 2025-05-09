package elements

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Times int64
}

func (m *Measurement) Add(value float64) {
	if m.Times == 0 {
		m.Min = value
		m.Max = value
	} else {
		if value < m.Min {
			m.Min = value
		}
		if value > m.Max {
			m.Max = value
		}
	}
	m.Sum += value
	m.Times++
}

func GetMeasurements(path string) (map[string]*Measurement, error) {
	// Create a map to store the measurements
	// The key is the location and the value is a pointer to the Measurement struct
	measurements := make(map[string]*Measurement)

	// Open the file
	measurement, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	defer measurement.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(measurement)

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		semicolon := strings.Index(line, ";")

		if semicolon == -1 {
			fmt.Println("Linha mal formatada:", line)
			continue
		}

		// For each line, split the line into location and value
		location := line[:semicolon]
		value := line[semicolon+1:]

		// If the location is not in the map, add it
		if _, ok := measurements[location]; !ok {
			measurements[location] = &Measurement{
				Times: 0,
			}
		}

		// Convert the value to a float64
		measurementValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			fmt.Println("Error converting value to float64:", err)
			panic(err)
		}

		// Add the value to the measurement
		measurements[location].Add(measurementValue)
	}

	return measurements, nil
}
