package znet

import (
	"errors"
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

// 定义当前客户端连接所绑定的handle api （目前这个handle是写死的，以后优化应该由用户自定义）
func CallBackToClient(conn *net.TCPConn, data []byte, n int) error {
	//回显的业务
	fmt.Println("[Conn Handel] CallBackToClient ......")
	if _, err := conn.Write(data[:n]); err != nil {
		fmt.Println("Write back buf err", err)
		return errors.New("CallBackToClient error")
	}
	return nil
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
			dealConn := NewConnection(conn, cid, CallBackToClient)
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
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
