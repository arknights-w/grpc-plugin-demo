package test

import (
	"fmt"
	"grpc-plugin/plugins/send"
	"grpc-plugin/tools"
	"io"
	"log"
	"os"
	"testing"
	"time"
)

func TestFirst(t *testing.T) {
	log.SetOutput(io.Discard)

	// 调用 chancel 手动关闭服务器
	clientProtocol, cancel := tools.CreateClient("./", "conf", "yml")

	raw_client, err := clientProtocol.Dispense("send")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	sm := raw_client.(send.SM)

	response := sm.Send("13388172385", "hello,world")

	fmt.Printf("\nresponse: %v\n\n", response)
	cancel()
	time.Sleep(1 * time.Second)
}
