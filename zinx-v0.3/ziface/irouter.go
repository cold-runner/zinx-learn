package ziface

/*
这里之所以BaseRouter的方法都为空，
是因为有的Router不希望有PreHandle或PostHandle
所以Router全部继承BaseRouter的好处是，不需要实现PreHandle和PostHandle也可以实例化
*/
type IRouter interface {
	PreHandle(request IRequest)  //在处理conn业务之前的钩子方法
	Handle(request IRequest)     //处理conn业务的方法
	PostHandle(request IRequest) //处理conn业务之后的钩子方法
}
