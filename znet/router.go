package znet

import "myZinx/zinx/ziface"

/*
实现Router时，我们可以预先实现一个基本的Router，
如果用户觉得不好，可以自己去实现irouter或者继承BaseRouter
*/

/*
设计思想：
1. 如果不设置BaseRouter，那么用户需要用router必须实现irouter的所有方法
2. 如果用户只想实现Handle的话，也要实现另外两个方法
3. 加了BaseRouter的话，因为BaseRouter是irouter的一个实现，如果你只需要实现某一个方法
你只需要继承BaseRouter，只实现一个方法即可，这样你也是实现了irouter，但是你只实现了一个方法
这样极大的减少了程序员的代码量，不用每次都重写多个方法，且用不着
*/
type BaseRouter struct {
}

func (br *BaseRouter) PreHandle(request ziface.IRequest) {

}
func (br *BaseRouter) Handle(request ziface.IRequest) {

}
func (br *BaseRouter) PostHandle(request ziface.IRequest) {

}
