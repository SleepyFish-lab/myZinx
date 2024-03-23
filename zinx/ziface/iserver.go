package ziface

type IServer interface {
	//开启服务器
	Start()
	//关闭服务器
	Stop()
	//开始服务
	Serve()
}
