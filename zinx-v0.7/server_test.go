package main

import (
	"fmt"
	"github.com/cold-runner/zinx-learn/zinx-v0.7/ziface"
	"github.com/cold-runner/zinx-learn/zinx-v0.7/znet"
	"testing"
)

type myRegret struct {
	znet.BaseRouter
}

func (m *myRegret) Handle(request ziface.IRequest) {
	fmt.Println("==> Recv Msg: ID=", request.GetMsgID(), ", len=", request.GetDataLen(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(request.GetMsgID(), []byte("I miss you, but not you, and even is not you..."))
	if err != nil {
		fmt.Println("send msg failed!")
		return
	}
}

type another struct {
	znet.BaseRouter
}

func (m *another) Handle(request ziface.IRequest) {
	fmt.Println("==> Recv Msg: ID=", request.GetMsgID(), ", len=", request.GetDataLen(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendMsg(request.GetMsgID(), []byte("base route test"))
	if err != nil {
		fmt.Println("send msg failed!")
		return
	}
}

func TestServer(t *testing.T) {
	s := znet.NewServer()

	s.AddRouter(233, &another{})
	s.AddRouter(1204, &myRegret{})
	// 添加相同msgId进行测试
	//s.AddRouter(233, &myRegret{})
	// panic: repeated api , msgId = 233
	s.Start()
	select {}
}
