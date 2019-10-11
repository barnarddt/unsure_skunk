package skunk

const RoundEventOffset = 100

type EventType int

func (t EventType) Valid() bool {
	return true
}

func (t EventType) ReflexType() int {
	return int(t)
}
