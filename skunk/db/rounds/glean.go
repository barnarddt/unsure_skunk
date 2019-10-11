package rounds

import "github.com/barnarddt/unsure_skunk/skunk"

//go:generate glean -table=rounds

type glean struct {
	skunk.Round
}
