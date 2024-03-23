package main

import "myZinx/zinx/znet"

// 基于zinx框架来开发，服务器端程序
func main() {

	//1.创建一个Server句柄，调用zinx的api
	s := znet.NewServer("[zinx V0.2]")
	//启动Server
	s.Serve()
}
