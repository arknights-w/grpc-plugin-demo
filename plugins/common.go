package plugins

import (
	"grpc-plugin/plugins/send"
	"grpc-plugin/plugins/send/service"

	"github.com/hashicorp/go-plugin"
	"github.com/spf13/viper"
)

// 用于检测 客户端/服务器 是否匹配
var Handshake_sm = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "Send_Message",
	MagicCookieValue: "hello",
}

// var PluginMap_sm = map[string]plugin.Plugin{
// 	"send": &send.SMPlugin{},
// }

func GetPlugins(config *viper.Viper) (plugins map[string]plugin.Plugin) {
	plugins = make(map[string]plugin.Plugin)
	plugins["send"] = &send.SMPlugin{Impl: &service.SM{}}
	return plugins
}
