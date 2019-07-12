package tcppage

// 接收器

import (
	"fmt"
	"net"
	"sync"
)

type Acceptor struct {
	// 保存监听器
	l net.Listener

	//监听器停止同步
	wg sync.WaitGroup

	// 链接的数据回调
	OnSessionData func(net.Conn, []byte) bool
}

func (a *Acceptor) Start(addr string)  {
	go a.Listen(addr)
}

func (a *Acceptor) Listen(addr string)  {
	// 监听开始
	a.wg.Add(1)

	// 延迟 结束监听
	defer a.wg.Done()

	var err error

	// 监听给定地址
	a.l, err = net.Listen("tcp", addr)

	if err != nil{
		fmt.Println(err)
		return
	}

	for {
		// 新链接没到来之前是阻塞的
		conn,err := a.l.Accept()

		if err != nil{
			fmt.Println(err)
			break
		}

		// 并发  开启会话
		go  handleSession(conn, a.OnSessionData)

	}

}

// 停止监听
func(a *Acceptor) Stop()  {
	a.l.Close()
}

// 等待监听完全停止
func (a *Acceptor) Wait()  {
	a.wg.Wait()
}

// 实例化监听器
func NewAcceptor() *Acceptor  {
	return &Acceptor{}
}