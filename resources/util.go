//go:build embed

package resources

import (
	"encoding/json"
)

// Since these embedded files are required to be correct, load will panic if it
// fails to decode them. I find this rather annoying that go:embed doesn't let
// us specify an embedded encoder so we may ourselves ensure that the data being
// embedded is correct, especially as it would suck to embed it at compile time,
// distribute it, and then find out during runtime that the files are incorrect.
// MUST rely on integeration tests for embed.
func load[T any](name string) (t T) {
	fs, err := f.Open(name)
	if err != nil {
		panic("Failed to open embedded file " + name + ":" + err.Error())
	}
	defer fs.Close()

	if err := json.NewDecoder(fs).Decode(&t); err != nil {
		panic("Failed to decode embedded file " + name + ":" + err.Error())
	}

	return t
}
