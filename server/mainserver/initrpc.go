package mainserver

import (
	"fmt"
	"net"
	"os"
	//"errors"
	"net/rpc"

	"../modes"
)

type RPCServer struct {
	strPort string
}

func (this *RPCServer) Start() {

	user := new(modes.Users)
	rpc.Register(user)

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":7001")

	if err != nil {
		fmt.Println("错误了哦")
		os.Exit(1)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	for {
		//需要自己控制连接，当有客户端连接上来后，我们需要把这个连接交给rpc 来处理
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		rpc.ServeConn(conn)
	}
}
