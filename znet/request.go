package znet

import "myZinx/zinx/ziface"

type Request struct {
	//已经建立好连接的connection
	conn ziface.IConnection
	//客户端请求的数据
	data []byte
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}
func (r *Request) GetData() []byte {
	return r.data
}
