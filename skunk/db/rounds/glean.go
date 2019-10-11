package rounds

import "unsure_skunk/skunk"

//go:generate glean -table=rounds

type glean struct {
	skunk.Round
}
