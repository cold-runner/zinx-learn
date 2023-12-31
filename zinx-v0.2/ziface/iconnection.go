package ziface

import "net"

// IConnection 定义连接接⼝口
type IConnection interface {
	// Start 启动连接，让当前连接开始⼯工作
	Start()
	// Stop 停⽌止连接，结束当前连接状态M
	Stop()
	// GetConnID 从当前连接获取原始的socket TCPConn GetTCPConnection() *net.TCPConn //获取当前连接ID
	GetConnID() uint32
}

// HandFunc 定义⼀一个统⼀一处理理链接业务的接⼝口
type HandFunc func(*net.TCPConn, []byte, int) error
