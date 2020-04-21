package main

import (
	"errors"

	gpb "github.com/gogo/protobuf/types"
	pb "github.com/kata-containers/agent/protocols/grpc"
	"golang.org/x/net/context"
)

type bublAgentGRPC struct {
	*agentGRPC
}

func (a *bublAgentGRPC) WriteStdin(ctx context.Context, req *pb.WriteStreamRequest) (*pb.WriteStreamResponse, error) {
	return nil, errors.New("not allowed")
}

func (a *bublAgentGRPC) ReadStdout(ctx context.Context, req *pb.ReadStreamRequest) (*pb.ReadStreamResponse, error) {
	return nil, errors.New("not allowed")
}

func (a *bublAgentGRPC) ReadStderr(ctx context.Context, req *pb.ReadStreamRequest) (*pb.ReadStreamResponse, error) {
	return nil, errors.New("not allowed")
}

func (a *bublAgentGRPC) CloseStdin(ctx context.Context, req *pb.CloseStdinRequest) (*gpb.Empty, error) {
	return nil, errors.New("not allowed")
}

func (a *bublAgentGRPC) TtyWinResize(ctx context.Context, req *pb.TtyWinResizeRequest) (*gpb.Empty, error) {
	return nil, errors.New("not allowed")
}
