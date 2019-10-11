package main

import (
	"flag"
	"net/http"
	"unsure_skunk/skunk/ops"
	"unsure_skunk/skunk/server"
	"unsure_skunk/skunk/skunkpb"

	"github.com/corverroos/unsure"
	"unsure_skunk/skunk/state"

	"github.com/luno/jettison/errors"
)

var (
	httpAddress = flag.String("http_address", ":12047", "skunk healthcheck address")
	grpcAddress = flag.String("grpc_address", ":12048", "skunk grpc server address")
)

func main() {
	unsure.Bootstrap()

	s, err := state.New()
	if err != nil {
		unsure.Fatal(errors.Wrap(err, "new state error"))
	}

	go serveGRPCForever(s)

	ops.StartLoops(s)

	http.HandleFunc("/health", makeHealthCheckHandler())
	go unsure.ListenAndServeForever(*httpAddress, nil)

	unsure.WaitForShutdown()
}

func serveGRPCForever(s *state.State) {
	grpcServer, err := unsure.NewServer(*grpcAddress)
	if err != nil {
		unsure.Fatal(errors.Wrap(err, "new grpctls server"))
	}

	skunkSrv := server.New(s)
	skunkpb.RegisterSkunkServer(grpcServer.GRPCServer(), skunkSrv)

	unsure.RegisterNoErr(func() {
		skunkSrv.Stop()
		grpcServer.Stop()
	})

	unsure.Fatal(grpcServer.ServeForever())
}

func makeHealthCheckHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	}
}
