package client

import (
	"context"
	"unsure_skunk/skunk/skunkpb/protocp"

	pb "unsure_skunk/skunk/skunkpb"

	"github.com/corverroos/unsure"
	"github.com/luno/reflex"
	"github.com/luno/reflex/reflexpb"
	"google.golang.org/grpc"

	"unsure_skunk/skunk"
)

var _ skunk.Client = (*client)(nil)

type client struct {
	address   string
	rpcConn   *grpc.ClientConn
	rpcClient pb.SkunkClient
}

type option func(*client)

func WithAddress(address string) option {
	return func(c *client) {
		c.address = address
	}
}

func New(opts ...option) (*client, error) {
	c := client{}
	for _, o := range opts {
		o(&c)
	}

	var err error
	c.rpcConn, err = unsure.NewClient(c.address)
	if err != nil {
		return nil, err
	}

	c.rpcClient = pb.NewSkunkClient(c.rpcConn)

	return &c, nil
}

func (c *client) Ping(ctx context.Context) error {
	_, err := c.rpcClient.Ping(ctx, &pb.Empty{})
	return err
}

func (c *client) Stream(ctx context.Context, after string, opts ...reflex.StreamOption) (reflex.StreamClient, error) {
	sFn := reflex.WrapStreamPB(func(ctx context.Context,
		req *reflexpb.StreamRequest) (reflex.StreamClientPB, error) {
		return c.rpcClient.Stream(ctx, req)
	})

	return sFn(ctx, after, opts...)
}

func (c *client) GetParts(ctx context.Context, roundId int64, player string) ([]skunk.PartType, int64, error) {
	res, err := c.rpcClient.GetData(ctx, &pb.GetDataReq{RoundId: roundId, Player: player})
	if err != nil {
		return nil, 0, err
	}

	var pt []skunk.PartType
	for _, ss := range res.Part {
		pt = append(pt, *protocp.PartTypeFromProto(ss))
	}

	return pt, res.Rank, nil
}
