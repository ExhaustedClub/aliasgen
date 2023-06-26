package resources

import (
	"encoding/json"
	"os"
	"testing"
)

// This test ensures that the resources we're embedding or shipping are
// correctly formatted lists.
func TestValidListResources(t *testing.T) {
	for _, file := range []string{
		"first-names.json",
		"last-names.json",
	} {
		t := t
		t.Run("check_"+file, func(t *testing.T) {
			f, err := os.Open(file)
			if err != nil {
				t.Fatalf("Unable to find resource file \"%s\": %v\n", file, err)
			}

			var l []string
			if err := json.NewDecoder(f).Decode(&l); err != nil {
				t.Fatalf("Failed to decode resource file \"%s\": %v\n", file, err)
			}

			if len(l) == 0 {
				t.Fatalf("Resource file \"%s\" contained no data...\n", file)
			}
		})
	}
}
