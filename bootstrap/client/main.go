package main

import (
	"fmt"
	"grpc-plugin/intf"
	inner "grpc-plugin/plugin"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-plugin"
)

const(
	TENCENT = "./tencent"
	HUAWEI = "./huawei"
)

func main(){
	// We don't want to see the plugin logs.
	log.SetOutput(io.Discard)

	// We're a host. Start by launching the plugin process.
	// 这个 client 可用于管理 插件子进程 的生命周期
	// HandshakeConfig	通讯密钥
	// Plugins			插件集，选择需要的插件进行加载
	// Cmd				rpc命令，可以选择服务器的二进制程序，这里仅用 TENCENT 举例
	// AllowedProtocols	允许协议
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: inner.Handshake_sm,
		Plugins:         inner.PluginMap_sm,
		Cmd:             exec.Command("sh", "-c", TENCENT),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})
	defer client.Kill()

	// 启动服务器端,同时返回客户端
	// 注意服务器是在客户端中启起来的
	// client-bootstrap
	clientProtocol, err := client.Client()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	// 这里选择客户端所需要的插件
	// 返回的 raw 是一个实际类型为 SMClient 的客户端
	// 这个客户端实现了 intf 中的 SM 接口，因此我们仅使用上层接口进行调用
	raw_client, err := clientProtocol.Dispense("send")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	sm := raw_client.(intf.SM)

	response := sm.Send("13388172385", "hello,world")

	fmt.Printf("response: %v\n", response)

	os.Exit(0)

}