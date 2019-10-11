package ops

import (
	"time"
)

const (
	maxRounds          = 10
	roundStatusTimeout = time.Minute // Timeout a round after this duration in single state.

	consumerStartRound    = "engine_start_round_%d"
	consumerTimeoutRound  = "engine_timeout_round_%s"
	consumerAdvanceRound  = "engine_advance_round_%s"
	consumerMatchComplete = "engine_complete_match"
)

func StartLoops(b Backends) {

}
