package server

import (
	"context"

	"github.com/corverroos/unsure/skunk/db/events"
	pb "github.com/corverroos/unsure/skunk/skunkpb"
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
		stream:  events.ToStream(b.EngineDB().DB),
	}
}

func (srv *Server) Stop() {
	srv.rserver.Stop()
}

func (srv *Server) Ping(ctx context.Context, req *pb.Empty) (*pb.Empty, error) {
	return req, nil
}
