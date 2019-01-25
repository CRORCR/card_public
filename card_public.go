package main

import (
	"public/server/db"
	"public/server/mainserver"
	"fmt"
	"public/server/modes"
)

func main() {
	var rpcserver mainserver.RPCServer
	if nil != db.InitDB() {
		panic("数据库打开失败...")
	}
	var redis db.RedisServer
	if err := redis.Start("./config/redis.json"); nil != err {
		fmt.Println("Redis Error: ", err)
		return
	}
	if err := modes.InviteInit("./config/output.txt"); nil != err {
		fmt.Println("File Error:", err)
		return
	}
	fmt.Println("REDIS")
	fmt.Println("数据库打开成功...")
	fmt.Println("启动RPC服务")
	rpcserver.Start()
}

