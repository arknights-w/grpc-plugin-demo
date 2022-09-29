package service

import "grpc-plugin/intf"

// 这个 SM结构体 实现 intf SM接口
// 服务器运行时，会将这个 SM结构体 注入到 SMServer的Impl 中
// 理论上来讲，我们更换运行插件时，只需要更换 SM结构体 的实现
// 其他代码都不需要动
type SM struct{}

func (SM) Send(phone string, text string) intf.Res {
	// TODO:doSomething

	return intf.Res{Reslut: true, Msg: "Succeed"}
}
