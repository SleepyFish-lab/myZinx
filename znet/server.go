package znet

import (
	"fmt"
	"myZinx/zinx/utils"
	"myZinx/zinx/ziface"
	"net"
)

// 实现IServer接口
type Server struct {
	//服务器的名称
	Name string
	//服务器的IP版本
	IPVersion string
	//服务器的ip地址
	IP string
	//服务器的端口
	Port int
	//当前的server添加一个router，server注册的连接对应的处理业务
	//可以设置为数组，记录多个Router，现在只设置为一个router
	Router ziface.IRouter
}

func (s *Server) Start() {
	fmt.Printf("[Zinx] Server Name: %s, Listenner at IP: %s, Port:%d is starting\n",
		utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[Zinx] Version: %s, MaxConn: %d, MaxPackageSize: %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPackageSize)
	go func() {
		//1.获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("net.ResolveTCPAddr error", err)
			return
		}
		//2.监听服务器地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("net.ListenTCP error", err)
			return
		}
		fmt.Println("start Zinx server succ,", s.Name, "succ, listenning....")

		var cid uint32 = 0

		//3.阻塞的等待客户端的连接，处理客户端的连接业务（读写）
		for {
			//如果有客户的连接，阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("listenner accept error", err)
				//连接失败是需要放弃这次连接，而不是放弃所有连接
				//即继续接收其他的连接
				continue
			}
			//解耦合，使用connection模块的handler处理
			//将处理业务的方法与conn进行绑定，得到我们的连接模块
			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			//启动当前的连接业务处理
			go dealConn.Start()

		}
	}()

}
func (s *Server) Stop() {
	//TODO 将服务器的一些资源、状态或者已经开辟的链接信息 进行停止或者回收
}
func (s *Server) Serve() {
	//调用start方法
	s.Start()

	//TODO 做一些启动服务器之后的额外业务

	//由于start方法中没有阻塞，这样会导致主协程直接退出，整个程序也直接退出了
	//所以要在此处进入阻塞状态
	select {}
}
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("Add Router Succ!!")
}
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}
	return s
}
