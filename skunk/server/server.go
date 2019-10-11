package server

import (
	"context"
	"github.com/luno/reflex/reflexpb"

	"github.com/luno/reflex"
	"unsure_skunk/skunk/db/events"
	pb "unsure_skunk/skunk/skunkpb"
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