//go:build !embed

package resources

import "fmt"

// LoadLocation will look in a directory for "first-names.json" as well as
// "last-names.json" when the resources are not embedded in the binary.
func LoadLocation(s string) error {
	return fmt.Errorf("non embedded resource loading is not yet supported")
}

// TODO:
// - Ensure that it loads it from the local filesystem.

// GetFirstNameList returns a list of first names.
func GetFirstNameList() []string { return nil }

// GetFirstNameList returns a list of last names.
func GetLastNameList() []string { return nil }
