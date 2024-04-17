package utils

import (
	"encoding/json"
	"myZinx/zinx/ziface"
	"os"
)

/*
存储一切有关zinx框架的全局参数，供其他模块使用
一些参数可以通过zinx.json由用户进行配置
*/
type GlobalObj struct {

	/*
		Server
	*/
	TcpServer ziface.IServer //当前Zinx全局的Server对象
	Host      string         //当前服务器主机监听的IP
	TcpPort   int            //当前服务器主机监听的port
	Name      string         //当前服务器的名字

	/*
		Zinx
	*/
	Version        string //当前Zinx的版本号
	MaxConn        int    //当前服务器主机运行的最大连接数量
	MaxPackageSize uint32 //当前Zinx框架数据包的最大值
}

/*
定义好一个全局的对外 globalobj
*/
var GlobalObject *GlobalObj

/*
从zinx.json去加载用户自定义的参数
*/
func (g *GlobalObj) Reload() {
	data, err := os.ReadFile("conf/zinx.json")

	if err != nil {
		panic(err)
	}
	//将json文件数据解析到struct中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

/*
提供一个init方法，初始化当前的GlobalObject
*/
func init() {

	//如果配置文件没有加载，默认的值
	GlobalObject = &GlobalObj{
		Name:           "ZinxServer",
		Version:        "V0.4",
		TcpPort:        8999,
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}

	//尝试从config/zinx.json去加载一些用户自定义的参数
	GlobalObject.Reload()
}
