package test

import (
	"fmt"
	"net/rpc"
	"public/server/modes"
	"testing"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/2/18
 */

var rateList []*modes.YoawoRate


func TestRate(t *testing.T) {
	//getTopList()
	//getList()
	//getListByName()
	getRateBy()
	//update()
}

/* 获得所有 level_id= 0 行业,拿到行业classid,根据classid去查询所有的数据*/
func getTopList() error {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")

	this:= &modes.YoawoRate{}
	err = client.Call("YoawoRate.GetTopList", this, &rateList)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Println(len(rateList))
	fmt.Printf("%+v\n",rateList)
	return err
}

/*根据classid去查询所有的数据*/
func getList() error {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")

	this:= &modes.YoawoRate{ClasId:1}
	err = client.Call("YoawoRate.GetList", this, &rateList)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Println(len(rateList))
	fmt.Printf("%+v\n",rateList)
	return err
}

/*根据name模糊查询*/
func getListByName() error {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")

	this:= &modes.YoawoRate{Name:"中"}
	err = client.Call("YoawoRate.GetListByName", this, &rateList)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Println(len(rateList))
	fmt.Printf("%+v\n",rateList)
	return err
}

/*获得一条数据*/
func getRateBy() error {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")

	this:= &modes.YoawoRate{Id:2}
	err = client.Call("YoawoRate.GetOne", this, this)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("%+v\n",this)
	return err
}

/*更新数据*/
func update() error {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")

	this:= &modes.YoawoRate{Id:2,Rate1:100}
	err = client.Call("YoawoRate.Update", this, this)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Println(len(rateList))
	fmt.Printf("%+v\n",this)
	return err
}