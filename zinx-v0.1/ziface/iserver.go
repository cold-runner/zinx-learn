package ziface

type IServer interface {
	// Start 启动服务器方法
	Start()
	// Stop 停止服务器方法
	Stop()
	// Serve 开启业务服务方法
	Serve()
}
