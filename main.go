package main

import (
	"time"
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
		panic("Redis Error" )
	}
	fmt.Println("Redis服务连接成功...")
	rdb := ndb.GetRndbHand()
	//rdb.Do("shane", "aaaaaaaaaaaa")
	rdb.Put("astaxie", 10000, 10*time.Second*20)
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	rpcserver.Start()
}
