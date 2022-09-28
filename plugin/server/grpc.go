package server

import (
	"grpc-plugin/intf"
	"grpc-plugin/proto"

	"golang.org/x/net/context"
)

type GRPCServer struct {
	// This is the real implementation
	Impl intf.KV
	proto.UnimplementedKVServer
}

func (m *GRPCServer) Put(
	ctx context.Context,
	req *proto.PutRequest) (*proto.Empty, error) {
	return &proto.Empty{}, m.Impl.Put(req.Key, req.Value)
}

func (m *GRPCServer) Get(
	ctx context.Context,
	req *proto.GetRequest) (*proto.GetResponse, error) {
	v, err := m.Impl.Get(req.Key)
	return &proto.GetResponse{Value: v}, err
}
