package send

import (
	"grpc-plugin/plugins/send/proto"

	"log"

	"golang.org/x/net/context"
)

type SMClient struct{
	Client proto.SendMessageClient
}

// 注意这里的 Send 是实现 intf 中的
func (sm *SMClient)Send(phone string,text string) Res{
	// 在这里面调用 grpc 的方法请求后端
	// 一定注意 context 是 net 包中的 context, 不是标准库的
	res,err :=sm.Client.Send(context.Background(),&proto.Msg{Phone: phone,Text: text})
	if err != nil {
		log.Println("--------------------------------------------------------------------")
		log.Printf("err: %v\n", err)
		panic("client send fail")
	}
	return Res{Reslut: res.Result,Msg: res.Msg}
}