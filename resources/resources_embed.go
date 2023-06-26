//go:build embed

package resources

import (
	"embed"
)

//go:embed first-names.json last-names.json
var f embed.FS

var (
	firstNames = load[[]string]("first-names.json")
	lastNames  = load[[]string]("last-names.json")
)

// GetFirstNameList returns a list of first names.
func GetFirstNameList() []string { return firstNames }

// GetFirstNameList returns a list of last names.
func GetLastNameList() []string { return firstNames }

// LoadLocation with embedded files will always load nil since it's expected
// that the resources have been loaded into the program at compile time, and
// then parsed for read at runtime, thus embedded resources have a MustLoad
// logic to them.
func LoadLocation(s string) error { return nil }
