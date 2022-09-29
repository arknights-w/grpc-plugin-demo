package main

import (
	"grpc-plugin/intf"
	inner "grpc-plugin/plugin"
	"log"

	"github.com/hashicorp/go-plugin"
)

// 这个 SM结构体 实现 intf SM接口
// 服务器运行时，会将这个 SM结构体 注入到 SMServer的Impl 中
// 理论上来讲，我们更换运行插件时，只需要更换 SM结构体 的实现
// 其他代码都不需要动
type SM struct{}

func (SM) Send(phone string, text string) intf.Res {
	// TODO:doSomething

	return intf.Res{Reslut: true, Msg: "Succeed"}
}

// 这就是一个服务端的启动器
// server-bootstrap
func main() {
	log.Println("----------------1111111111111111-----------------")
	// 一定注意这里的 Plugins配置项,
	// 和 PluginMap_sm 的区别
	// PluginMap_sm 中没有参数
	// 这里的加入了 Impl 的实现
	// PluginMap_sm 仅使用在前端作为插件的验证
	// 这里 Impl 的实现是 可插拔的业务
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: inner.Handshake_sm,
		Plugins: map[string]plugin.Plugin{
			"send": &inner.SMPlugin{Impl: &SM{}},
		},

		GRPCServer: plugin.DefaultGRPCServer,
	})
}
