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

func main() {
	// We don't want to see the plugin logs.
	log.SetOutput(io.Discard)

	// We're a host. Start by launching the plugin process.
	// 这个 client 可用于管理 插件子进程 的生命周期
	// HandshakeConfig	通讯密钥
	// Plugins			插件集，选择需要的插件进行加载
	// Cmd				rpc命令，可以选择服务器的二进制程序，同时参数选择服务指令 get,put,send...
	// AllowedProtocols	允许协议
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: inner.Handshake,
		Plugins:         inner.PluginMap,
		Cmd:             exec.Command("sh", "-c", os.Getenv("KV_PLUGIN")),
		AllowedProtocols: []plugin.Protocol{
			plugin.ProtocolNetRPC, plugin.ProtocolGRPC},
	})
	defer client.Kill()

	// Connect via RPC
	// 启动服务器端,同时返回客户端
	// 注意服务器是在客户端中启起来的
	// client-bootstrap
	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// Request the plugin
	// 这里选择客户端所需要的插件
	// 返回的 raw 是一个实际类型为 KVGRPCClient 的客户端
	// 这个客户端实现了 intf 中的 KV 接口，因此我们仅使用上层接口进行调用
	// 这里可以做一些抽象，比如 switch case 进行多条件选择执行我们想要的插件
	raw, err := rpcClient.Dispense("kv_grpc")
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	// We should have a KV store now! This feels like a normal interface
	// implementation but is in fact over an RPC connection.
	kv := raw.(intf.KV)
	os.Args = os.Args[1:]
	switch os.Args[0] {
	case "get":
		result, err := kv.Get(os.Args[1])
		if err != nil {
			fmt.Println("Error:", err.Error())
			os.Exit(1)
		}

		fmt.Println(string(result))

	case "put":
		err := kv.Put(os.Args[1], []byte(os.Args[2]))
		if err != nil {
			fmt.Println("Error:", err.Error())
			os.Exit(1)
		}

	default:
		fmt.Printf("Please only use 'get' or 'put', given: %q", os.Args[0])
		os.Exit(1)
	}
	os.Exit(0)
}
