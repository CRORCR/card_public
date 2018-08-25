package main

import(
	"fmt"
	"./server/db"
	"./server/mainserver"
)


func main(){
	var rpcserver mainserver.RPCServer
	if nil != db.InitDB() {
		panic("数据库打开失败...")
	}
	fmt.Println("数据库打开成功...")
	rpcserver.Start()
}
