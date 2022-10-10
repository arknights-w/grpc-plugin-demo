package send

import (
	"context"

	"google.golang.org/grpc"

	"grpc-plugin/plugins/send/proto"


	"github.com/hashicorp/go-plugin"
)

type SMPlugin struct {
	plugin.Plugin

	// 这里使用上层接口
	// 为了工程结构与下层实现分离
	Impl SM
}

func (sm *SMPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterSendMessageServer(s, &SMServer{Impl: sm.Impl})
	return nil
}

func (sm *SMPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &SMClient{Client: proto.NewSendMessageClient(c)}, nil
}
