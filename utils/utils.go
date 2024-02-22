package utils

import "strings"

// Function to append a value to the map for the given key
func AppendValue(m map[string][]string, key string, value string) {
	m[key] = append(m[key], value)
}

// filterByTimestamp filters interactions based on the provided timestamp
func FilterByTimestamp(interactions []string, timestamp string) []string {
	var filteredInteractions []string

	// Iterate through interactions and filter based on timestamp
	for _, interaction := range interactions {
		if strings.Contains(interaction, timestamp) {
			filteredInteractions = append(filteredInteractions, interaction)
		}
	}

	return filteredInteractions
}
