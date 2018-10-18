package main

import (
    "fmt"
    "net/rpc"

    "../server/modes"
)

type RPCServer struct{
    IP      string `json:"rpc_ip"`
    Type    string `json:"rpc_type"`
    Rpc    *rpc.Client
}

func main() {
    //var unid = modes.UnionId{"shane1234567890", "unionid_android", "19803066666"}
    var user modes.Users
    client, err := rpc.Dial("tcp", "127.0.0.1:7001")
    if err != nil {
        fmt.Println("连接RPC服务失败：", err)
    }
    fmt.Println("连接RPC服务成功")
    user.Phone = "19803091863"

    err = client.Call("Users.PhoneLogin", "19803091863", &user)
    if err != nil {
        fmt.Println("调用失败：", err)
    }
    fmt.Println("调用结果：", user)
}
