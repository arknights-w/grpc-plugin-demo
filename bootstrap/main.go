package main

import (
	"grpc-plugin/config"
	"grpc-plugin/plugins"

	"github.com/hashicorp/go-plugin"
)

// 这就是一个服务端的启动器
// server-bootstrap
func main() {

	conf := config.AutoConfig()
	plugins_get := plugins.GetPlugins(conf)

	// 一定注意这里的 Plugins配置项,
	// 和 PluginMap_sm 的区别
	// PluginMap_sm 中没有参数
	// 这里的加入了 intf 的实现
	// PluginMap_sm 仅使用在前端作为插件的验证
	// 这里 Impl 的实现是 可插拔的业务
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: plugins.Handshake_sm,
		GRPCServer:      plugin.DefaultGRPCServer,

		Plugins: plugins_get,
	})
}
