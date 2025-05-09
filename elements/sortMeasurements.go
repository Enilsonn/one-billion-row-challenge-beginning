package elements

import "sort"

func SortMeasurements(Measurements map[string]*Measurement) []string {
	// Create a slice to store the names of the measurements
	keys := make([]string, 0, len(Measurements))

	for key := range Measurements {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}
