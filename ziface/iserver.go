package ziface

type IServer interface {
	//开启服务器
	Start()
	//关闭服务器
	Stop()
	//开始服务
	Serve()
	//路由功能：给当前的服务注册一个路由方法，供客户端的连接处理使用
	AddRouter(router IRouter)
}
