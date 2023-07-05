package ziface

type IMsgHandle interface {
	DoMsgHandler(request IRequest)          //马上以非阻塞方式处理消息
	AddRouter(msgId uint32, router IRouter) //为消息添加具体地处理逻辑
}
