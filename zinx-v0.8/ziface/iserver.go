package ziface

type IServer interface {
	// Start 启动服务器方法
	Start()
	// Stop 停止服务器方法
	Stop()
	// Serve 开启业务服务方法
	Serve()

	// AddRouter AddHandler AddRouter 路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
	AddRouter(msgId uint32, router IRouter)
}
