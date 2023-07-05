package main

import (
	"fmt"
	"github.com/cold-runner/zinx-learn/zinx-v0.6/znet"
	"net"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	//客户端goroutine，负责模拟粘包的数据，然后进行发送
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		t.Fatal("client dial err:", err)
		return
	}
	for {
		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMsgPackage(0, []byte("I'm HuaiYu!")))
		if err != nil {
			fmt.Println("pack msg failed!", err)
			return
		}
		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}
