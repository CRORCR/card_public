package main

import (
	"./server/db"
	"./server/ndb"
	"./server/mainserver"
	"fmt"
)

func main() {
	var rpcserver mainserver.RPCServer
	if nil != db.InitDB() {
		panic("数据库打开失败...")
	}
	fmt.Println("数据库打开成功...")
	err := ndb.Init("./config/redis.json")
	if nil != err {
		panic("Redis Error:",err )
	}
	rpcserver.Start()
}
