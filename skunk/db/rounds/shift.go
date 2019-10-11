package rounds

import (
	"github.com/luno/shift"

	"unsure_skunk/skunk"
	"unsure_skunk/skunk/db/events"
)

//go:generate shiftgen -inserter=ready -updaters=joined,collected,empty -table=rounds

var roundFSM = shift.NewFSM(events.GetTable()).
	Insert(skunk.RoundStatusJoin, ready{}, skunk.RoundStatusJoined).
	Update(skunk.RoundStatusJoined, joined{}, skunk.RoundStatusCollect).
	Update(skunk.RoundStatusCollect, empty{}, skunk.RoundStatusCollected).
	Update(skunk.RoundStatusCollected, collected{}, skunk.RoundStatusSubmit).
	Update(skunk.RoundStatusSubmit, empty{}, skunk.RoundStatusSubmitted).
	Update(skunk.RoundStatusSubmitted, empty{}, skunk.RoundStatusSuccess,
		skunk.RoundStatusFailed).
	Build()


type ready struct {
	Player string
	ExternalID int64
}

type joined struct {
	ID int64
}

type collected struct{
	ID int64
	Rank int64
}

type empty struct {
	ID int64
}