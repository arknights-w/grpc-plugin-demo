package server

import (
	"grpc-plugin/intf"
	"grpc-plugin/proto"

	"golang.org/x/net/context"
)

// 这是 proto Server的一个实现
type SMServer struct {
	Impl intf.SM
	proto.UnimplementedSendMessageServer
}

// 注意这个 Send 是 proto Send 接口的实现
// 不是 Impl Send 接口的实现
// 一定注意 context 是 net 包中的 context, 不是标准库的
func (sm *SMServer) Send(ctx context.Context, req *proto.Msg) (*proto.Res, error) {
	r := sm.Impl.Send(req.Phone, req.Text)
	res := proto.Res{
		Result: r.Reslut,
		Msg:    r.Msg,
	}
	return &res, nil
}
