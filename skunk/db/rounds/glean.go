package rounds

import "unsure_skunk/skunk"

//go:generate glean -table=rounds -src=$GOPATH/src/github.com/barnarddt

type glean struct {
	skunk.Round
}
