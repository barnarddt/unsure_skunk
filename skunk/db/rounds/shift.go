package rounds

import (
	"github.com/luno/shift"

	"unsure_skunk/skunk"
	"unsure_skunk/skunk/db/events"
)

//go:generate shiftgen -inserter=ready -updaters=joined,update -table=rounds

var roundFSM = shift.NewFSM(events.GetTable()).
	Insert(skunk.RoundStatusJoin, ready{}, skunk.RoundStatusJoined).
	Update(skunk.RoundStatusJoined, joined{}, skunk.RoundStatusCollect).
	Update(skunk.RoundStatusCollect, update{}, skunk.RoundStatusCollected).
	Update(skunk.RoundStatusCollected, update{}, skunk.RoundStatusSubmit).
	Update(skunk.RoundStatusSubmit, update{}, skunk.RoundStatusSubmitted).
	Update(skunk.RoundStatusSubmitted, update{}, skunk.RoundStatusSuccess,
		skunk.RoundStatusFailed).
	Build()


type ready struct {
	Player string
}

type joined struct {
	ID int64
	ExternalID int64
	Player string
	Rank int
}

type update struct{
	ID int64
}