package znet

import (
	"fmt"
	"myZinx/zinx/utils"
	"myZinx/zinx/ziface"
	"net"
)

type Connection struct {
	//当前连接的socket TCP套接字
	Conn *net.TCPConn
	//连接的ID
	ConnID uint32
	//当前的连接状态
	isClosed bool

	//告知当前连接已经退出/停止 channel
	ExitChan chan bool
	//该连接处理业务的router
	Router ziface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c

}

// 处理连接的读业务
func (c *Connection) StartReader() {
	fmt.Println("Reader goroutine is running...")
	defer fmt.Println("Conn ID = ", c.ConnID, "Reader is exit", "Remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//读取客户端的数据到buf中
		buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		_, err := c.Conn.Read(buf)
		if err != nil {
			//这次读取失败，并不代表下一次读取失败
			//所以进入下一层的循环读取
			fmt.Println("recv buf err", err)
			continue
		}

		//得到当前conn数据的Request请求数据
		req := Request{
			conn: c,
			data: buf,
		}
		//执行注册的路由方法
		//从路由中，找到注册绑定的Conn对应的Router调用
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}
func (c *Connection) Start() {
	fmt.Println("Conn start().. Conn ID = ", c.ConnID)
	//启动从当前连接的读数据的业务
	go c.StartReader()
	//TODO 启动从当前连接的写数据的业务
}
func (c *Connection) Stop() {
	fmt.Println("Conn stop().. Conn ID = ", c.ConnID)
	//如果连接已经关闭，则不用关闭了
	if c.isClosed {
		return
	}
	c.isClosed = true
	//关闭tcp连接和管道，回收资源
	c.Conn.Close()
	close(c.ExitChan)
}
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()

}
func (c *Connection) Send(data []byte) error {
	return nil
}
