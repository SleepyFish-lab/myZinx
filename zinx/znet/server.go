package znet

import (
	"fmt"
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
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP:%s, Port:%d is starting.\n", s.IP, s.Port)

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
			//已经与客户端建立连接，做一些业务：做一个基本的512字节的内容回显
			go func() {
				for {
					buf := make([]byte, 512)
					n, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf error", err)
						continue
					}
					fmt.Printf("recv buf %s,cnt=%d ", buf, n)
					//回显功能
					if _, err := conn.Write(buf[:n]); err != nil {
						fmt.Println("write buf back err", err)
						continue
					}
					fmt.Println()
				}
			}()
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
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
