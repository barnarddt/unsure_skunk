package skunk

import (
	"context"
)

// Client defines the root engine service interface.
type Client interface {
	Ping(context.Context) error
}
