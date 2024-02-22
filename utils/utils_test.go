package utils_test

import (
	"testing"

	"github.com/SanjaySinghRajpoot/interactWrapper/utils"
)

func TestAppendValue(t *testing.T) {
	// Initialize a map
	m := make(map[string][]string)

	// Add values to the map
	utils.AppendValue(m, "key1", "value1")
	utils.AppendValue(m, "key2", "value2")

	// Check if values are correctly added
	if len(m["key1"]) != 1 || m["key1"][0] != "value1" {
		t.Errorf("expected value1 for key1 but got %v", m["key1"])
	}
	if len(m["key2"]) != 1 || m["key2"][0] != "value2" {
		t.Errorf("expected value2 for key2 but got %v", m["key2"])
	}
}

func TestFilterByTimestamp(t *testing.T) {
	// Test case 1: interactions contain the timestamp
	interactions := []string{"2024-02-22 12:00:00 - Interaction 1", "2024-02-23 13:00:00 - Interaction 2"}
	filteredInteractions := utils.FilterByTimestamp(interactions, "2024-02-22")
	if len(filteredInteractions) != 1 || filteredInteractions[0] != "2024-02-22 12:00:00 - Interaction 1" {
		t.Errorf("expected ['2024-02-22 12:00:00 - Interaction 1'] but got %v", filteredInteractions)
	}

	// Test case 2: interactions do not contain the timestamp
	filteredInteractions = utils.FilterByTimestamp(interactions, "2024-02-24")
	if len(filteredInteractions) != 0 {
		t.Errorf("expected [] but got %v", filteredInteractions)
	}
}
