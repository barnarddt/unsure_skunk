package skunk

import "github.com/luno/reflex"

type consumer = reflex.ConsumerName

const (
	// ConsumerJoinRounds consumes local shift events for records moving into
	// RoundStatusJoin and attempts to join the next round.
	ConsumerJoinRounds consumer = "join_rounds"

	// ConsumerCollectParts consumes local shift events for records moving into
	// RoundStatusCollect and attempts to collect parts from the engine.
	ConsumerCollectParts consumer = "collect_parts"

	// ConsumerSkipLocalJoined consumes local shift events for records moving
	// into RoundStatusJoined and automatically shifts them into
	// RoundStatusCollect.
	ConsumerSkipLocalJoined consumer = "skip_local_joined"

	// ConsumerSubmitParts consumes local shift events for records moving into
	// RoundStatusSubmit and sends their parts to the engine.
	ConsumerSubmitParts consumer = "submit_parts"
)
