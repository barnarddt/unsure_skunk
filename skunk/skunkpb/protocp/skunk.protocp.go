package protocp

import (
	"unsure_skunk/skunk"
	"unsure_skunk/skunk/skunkpb"
)

func PartTypeToProto(in *skunk.PartType) *skunkpb.Part {
	return &skunkpb.Part{
		Id:      in.ID,
		Part:    in.Part,
		Player:  in.Player,
		RoundId: in.RoundID,
	}
}

func PartTypeFromProto(in *skunkpb.Part) *skunk.PartType {
	return &skunk.PartType{
		ID:      in.Id,
		Part:    in.Part,
		Player:  in.Player,
		RoundID: in.RoundId,
	}
}
