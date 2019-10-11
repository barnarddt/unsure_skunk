package parts

import "unsure_skunk/skunk"

//go:generate glean -table=parts -src=$GOPATH/repos
type glean struct {
	skunk.PartType
}
