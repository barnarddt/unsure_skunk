package server

import (
	"context"
	"unsure_skunk/skunk/ops"
	"unsure_skunk/skunk/skunkpb/protocp"

	"github.com/luno/reflex/reflexpb"

	"unsure_skunk/skunk/db/events"
	pb "unsure_skunk/skunk/skunkpb"

	"github.com/luno/reflex"
)

var _ pb.SkunkServer = (*Server)(nil)

// Server implements the engine grpc server.
type Server struct {
	b       Backends
	rserver *reflex.Server
	stream  reflex.StreamFunc
}

// New returns a new server instance.
func New(b Backends) *Server {
	return &Server{
		b:       b,
		rserver: reflex.NewServer(),
		stream:  events.ToStream(b.SkunkDB().DB),
	}
}

func (srv *Server) Stop() {
	srv.rserver.Stop()
}

func (srv *Server) Ping(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	return req, nil
}

func (srv *Server) Stream(req *reflexpb.StreamRequest, ss pb.Skunk_StreamServer) error {
	return srv.rserver.Stream(srv.stream, req, ss)
}

func (srv *Server) GetData(ctx context.Context, req *pb.GetDataReq) (*pb.GetDataRes, error) {

	data1, rank, err := ops.LookUpData(ctx, srv.b, req.RoundId)
	if err != nil {
		return nil, err
	}

	res := &pb.GetDataRes{Rank: int32(rank)}
	for _, dt := range data1 {
		res.Part = append(res.Part, protocp.PartTypeToProto(&dt))
	}

	return res, nil
}
