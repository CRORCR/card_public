package test

import (
	"fmt"
	"net/rpc"
	modes2 "public/server/modes"
	"testing"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/1/19
 */

type CoordinatesPoint struct {
	Longitude float64 //经       度
	Latitude  float64 //纬       度
	Page      int
	OfferSet  int
}

func TestMer(t *testing.T) {
	//addMer() //添加商家信息
	//getMer()  //获得商家信息
	//findBranch()
	//getMerStaff() //查询商家所有员工
	//UpdateTrust() //更新锘豆
	//UpdateStatus() //更新商家状态 审核 未认证过:0  审核通过:1
	//getMargin()    //获得20公里以内的所有商家
}

//添加商家信息
func addMer() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")
	mer := modes2.Merchant{
		Phone:      "19803091863",
		MerchantId: "22",
		UserId:     "SHANE",
		AreaNumber: 310,
		Longitude:  116.404,
		Latitude:   39.915,
	}
	err = client.Call("Merchant.Add", &mer, &mer)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", mer)
}

/* 根据商家id 获取商家   必填参数 商家id 必须使用rpc添加的用户才能知道区号*/
func getMer() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")

	mer := modes2.Merchant{
		MerchantId: "33",
	}
	mer2:=modes2.Merchant{}
	err = client.Call("Merchant.Get", &mer, &mer2)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", mer2)
}

func findBranch() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")

	mer := modes2.Merchant{
		MerchantId: "33",
		FID:"22",
	}
	mer2:=modes2.MerchantList{}
	err = client.Call("Merchant.FindBranch", &mer, &mer2)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", mer2)
}

//获得商家所有员工   给定商家id
func getMerStaff() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")

	var merc modes2.Merchant
	var mer modes2.StaffList
	merc.MerchantId = "33"
	err = client.Call("Merchant.GetStaff", &merc, &mer)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", mer)
}


//获得商家所有没有权限员工   给定商家id
func getMerStaffNotAuth() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")

	var merc modes2.Merchant
	var mer modes2.StaffList
	merc.MerchantId = "33"
	err = client.Call("Merchant.GetStaffNotAuth", &merc, &mer)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", mer)
}

//更新锘豆  给定商家id和锘豆数量(上层处理)
func UpdateTrust() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")

	var merc = modes2.Merchant{MerchantId: "33", Trust: 100}
	var mer modes2.Merchant
	err = client.Call("Merchant.UpdateTrust", &merc, &mer)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", mer)
}

//更新商家状态
func UpdateStatus() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")

	var merc = modes2.Merchant{MerchantId: "33", Status: 1}
	var mer modes2.Merchant
	err = client.Call("Merchant.UpdateStatus", &merc, &mer)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", mer)
}

func getMargin() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")
	var margin = CoordinatesPoint{114.5315555000, 36.6479443200, 1, 4}
	mer := make(modes2.MerchantList, 0)
	err = client.Call("Merchant.GetNearMerchant", &margin, &mer)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", mer)
}
