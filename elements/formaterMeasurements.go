package elements

import "fmt"

func FormaterMeasurements(m map[string]*Measurement) string {
	// Getting the sorted keys
	keys := SortMeasurements(m)

	for _, key := range keys {
		fmt.Printf("%s=%.1f/%.1f/%.1f, ",
			key,
			m[key].Min,
			m[key].Sum/float64(m[key].Times),
			m[key].Max,
		)
	}
	return ""
}
