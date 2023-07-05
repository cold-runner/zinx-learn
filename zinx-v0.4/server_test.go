package zinx_v0_3

import (
	"fmt"
	"github.com/cold-runner/zinx-learn/zinx-v0.4/ziface"
	"github.com/cold-runner/zinx-learn/zinx-v0.4/znet"
	"testing"
)

// ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter //一定要先基础BaseRouter
}

// Test PreHandle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("Before ping .....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

// Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ping ping.....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

// Test PostHandle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Print("Call Router PostHandle\n\n")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping .....\n\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func TestServer(t *testing.T) {
	//创建一个server句柄
	s := znet.NewServer("[zinx V0.3]")

	s.AddRouter(&PingRouter{})

	//2 开启服务
	s.Serve()
}
