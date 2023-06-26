package alias

import (
	"math/rand"
	"time"
)

// Default source of randomness. Not true random.
var (
	s = rand.NewSource(time.Now().UnixNano())
	r = rand.New(s) // initialize local pseudorandom generator
)

// RandomFromList attempts to pull a random entry out of a provided list. Does
// not take options nor does it take seeding.
func RandomFromList[T any](l []T, seed rand.Source) T {
	if seed != nil {
		return l[rand.New(seed).Intn(len(l))]
	}
	return l[r.Intn(len(l))]
}
