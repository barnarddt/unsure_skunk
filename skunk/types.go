package skunk

import "time"

const RoundEventOffset = 100

type EventType int

func (t EventType) Valid() bool {
	return true
}

func (t EventType) ReflexType() int {
	return int(t)
}

type PartType struct {
	ID        int64
	RoundID   int64
	Player    string
	Part      int64
	CreatedAt time.Time
}
