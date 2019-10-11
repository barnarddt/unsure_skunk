package skunk

import "github.com/luno/reflex"

type consumer = reflex.ConsumerName

const (
	// ConsumerJoinRounds consumes local shift events for records moving into
	// RoundStatusJoin and attempts to join the next round.
	ConsumerJoinRounds consumer = "join_rounds"
)
