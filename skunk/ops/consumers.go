package ops

import (
	"context"
	"github.com/luno/fate"
	"github.com/luno/reflex"
)

func makeConsume(b Backends) reflex.Consumer {
	fn := func(ctx context.Context, fate fate.Fate, e *reflex.Event) error {

		return fate.Tempt()
	}

	return reflex.NewConsumer(reflex.ConsumerName("skunk_consumer"), fn)
}
