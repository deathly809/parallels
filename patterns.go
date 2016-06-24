package parallels

import (
	"github.com/deathly809/gotypes"
)

// Map takes in a function and how many
func Map(data []gotypes.Value, f func(gotypes.Value) gotypes.Value) {
	myFunc := func(i int) bool {
		data[i] = f(data[i])
		return false
	}
	Foreach(myFunc, len(data))
}
