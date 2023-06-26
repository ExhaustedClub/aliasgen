//go:build embed

package resources

import (
	"testing"
)

// This test is a highly moot point. I wanted it to help test the embed, but to
// have to manually specify a tag for testing defeats the purpose imho. Need to
// look into better testing, maybe.

func TestLoad(t *testing.T) {
	for name, test := range map[string]struct {
		Input    string
		Contents bool
		Error    interface{}
	}{
		"valid_first_name_list": {
			Input:    "first-names.json",
			Contents: true,
		},
		"valid_last_name_list": {
			Input:    "last-names.json",
			Contents: true,
		},
		"missing_file": {
			Input: "not-existing.json",
			Error: "Failed to open embedded file not-existing.json:open not-existing.json: file does not exist",
		},
		// This test is excluded since it is impossible to get a failed JSON
		// read without the first two test failing, since that's the only JSON
		// files we have.
		// "not_json": {
		// 	Input: "...",
		// 	Error: "...",
		// },
	} {
		t := t
		t.Run(name, func(t *testing.T) {
			var output []string
			var err interface{}
			func() {
				defer func() { err = recover() }()
				output = load[[]string](test.Input)
			}()
			if err != test.Error {
				t.Fatalf("Expected Error %v got %v\n", test.Error, err)
			}
			if (len(output) != 0) != test.Contents {
				t.Fatalf("Expected Contents(%v) got (%v)", test.Contents, len(output) != 0)
			}
		})
	}
}
