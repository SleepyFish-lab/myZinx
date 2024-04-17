package ziface

/*
路由的抽象接口：
路由里的数据都是IRequest
*/
type IRouter interface {
	//在处理conn业务之前的方法Hook
	PreHandle(request IRequest)
	//处理conn业务的方法hook
	Handle(request IRequest)
	//在处理conn业务之后的方法Hook
	PostHandle(request IRequest)
}
