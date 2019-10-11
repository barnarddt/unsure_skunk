package skunk

import "time"

//go:generate stringer -type=RoundStatus -trimprefix=RoundStatus

// RoundStatus defines the current status in the state machine.
type RoundStatus int

// Enum satisfied the shift.Status interface.
func (rs RoundStatus) Enum() int {
	return int(rs)
}

// ReflexType satisfied the shift.Status interface.
func (rs RoundStatus) ReflexType() int {
	return int(rs)
}

// ShiftStatus satisfied the shift.Status interface.
func (rs RoundStatus) ShiftStatus(){}

const (
	// RoundStatusUnknown is an invalid status which usually indicates missing
	// data.
	RoundStatusUnknown    RoundStatus = 0

	// RoundStatusJoin indicates that a Peer is ready to join a round.
	RoundStatusJoin      RoundStatus = 1

	// RoundStatusJoined indicates that a Peer has successfully joined a round.
	RoundStatusJoined    RoundStatus = 2

	// RoundStatusCollect indicates that a Peer is ready to collect their parts.
	RoundStatusCollect   RoundStatus = 3

	// RoundStatusCollected indicates that a Peer has successfully collected
	// their parts.
	RoundStatusCollected RoundStatus = 4

	// RoundStatusSubmit indicates that a Peer is ready to submit their parts.
	RoundStatusSubmit    RoundStatus = 5

	// RoundStatusSubmitted indicates that a Peer has successfully submitted
	// their parts.
	RoundStatusSubmitted RoundStatus = 6

	// RoundStatusSuccess indicates that the round was successful.
	RoundStatusSuccess   RoundStatus = 7

	// RoundStatusSuccess indicates that the round was failed.
	RoundStatusFailed    RoundStatus = 8

	// must be last.
	roundStatusSentinel  RoundStatus = 9
)

// Valid returns whether "rs" is a declared RoundStatus constant.
func (rs RoundStatus) Valid() bool {
	return rs > RoundStatusUnknown && rs < roundStatusSentinel
}

type Round struct {
	ID int64
	// The RoundID as known on Engine.
	ExternalID int64
	// Unique name for each Peer.
	Player string
	// Rank for the Peer as provided by Engine.
	Rank int
	// Current status for the unique Peer.
	Status RoundStatus
	CreatedAt time.Time
	UpdatedAt time.Time
	
}
