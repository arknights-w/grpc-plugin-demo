package plugin

import (
	"context"

	"google.golang.org/grpc"

	"grpc-plugin/intf"
	"grpc-plugin/plugin/client"
	"grpc-plugin/plugin/server"
	"grpc-plugin/proto"

	"github.com/hashicorp/go-plugin"
)

// 用于检测 客户端/服务器 是否匹配
var Handshake_sm = plugin.HandshakeConfig{
	ProtocolVersion: 1,
	MagicCookieKey: "Send_Message",
	MagicCookieValue: "hello",
}

var PluginMap_sm = map[string]plugin.Plugin{
	"send": &SMPlugin{},
}

type SMPlugin struct{
	plugin.Plugin

	// 这里使用上层接口
	// 为了工程结构与下层实现分离
	Impl intf.SM
}

func (sm *SMPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterSendMessageServer(s,&server.SMServer{Impl: sm.Impl})
	return nil
}

func (sm *SMPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &client.SMClient{Client: proto.NewSendMessageClient(c)},nil
}