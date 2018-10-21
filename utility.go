package utility

import "math/rand"

// ItemInCollection checks to see if a string is in an array of strings
func ItemInCollection(item string, collection []string) bool {
	for _, element := range collection {
		if item == element {
			return true
		}
	}
	return false
}

// Min returns the smaller of two integers
func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// RandomItem returns a random string from an array of strings
func RandomItem(items []string) string {
	return items[rand.Intn(len(items))]
}

// RandomItemFromThresholdMap returns a random weighted string
func RandomItemFromThresholdMap(items map[string]int) string {
	result := ""
	ceiling := 0
	start := 0
	var thresholds = make(map[string]int)

	for item, weight := range items {
		ceiling += weight
		thresholds[item] = start
		start += weight
	}

	randomValue := rand.Intn(ceiling)

	for item, threshold := range thresholds {
		if threshold <= randomValue {
			result = item
		}
	}

	return result
}
