package test

import (
	"fmt"
	"net/rpc"
	"card_public/server/modes"
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
	//getMerStaff() //查询商家所有员工
	//UpdateTrust() //更新锘豆
	//UpdateStatus() //更新商家状态 审核 未认证过:0  审核通过:1
	getMargin()    //获得20公里以内的所有商家
	//addBranch()
	//getAllBranch()
}

//添加商家信息
func addMer() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")
	mer := modes.Merchant{
		Phone:      "17600381202",
		MerchantId: "bbbb6cf8d45c8bef6e0c734bb02a635f",
		UserName:   "奇葩1号店",
		UserId:     "3dfa6cf8d45c8bef6e0c734bb02abbbb",
		AreaNumber: 319,
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

	mer := modes.Merchant{
		MerchantId: "bbbb6cf8d45c8bef6e0c734bb02a635f",
	}
	mer2 := modes.Merchant{}
	err = client.Call("Merchant.Get", &mer, &mer2)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", mer2)
}

//绑定
func addBranch() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")

	mer:=modes.MerchantAddBranch{Superior:"aaaa6cf8d45c8bef6e0c734bb02a635f",Lower:"bbbb6cf8d45c8bef6e0c734bb02a635f"}
	mer2 := modes.Merchant{}
	err = client.Call("Merchant.AddBranch", &mer, &mer2)
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

	var merc modes.Merchant
	var mer modes.StaffList
	merc.MerchantId = "dpfa6cf8d45c8bef6e0c734bb02a635f"
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

	var merc modes.Merchant
	var mer modes.StaffList
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

	var merc = modes.Merchant{MerchantId: "33", Trust: 100}
	var mer modes.Merchant
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

	var merc = modes.Merchant{MerchantId: "33", Status: 1}
	var mer modes.Merchant
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
	var margin = CoordinatesPoint{116.3990880000,39.9096820000, 3, 2}
	mer := modes.MerchantList{}
	err = client.Call("Merchant.GetNearMerchant", &margin, &mer)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", mer)
}

func getAllBranch() {
	// 123456789 123456789_311 123456789_319
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
		return
	}
	fmt.Println("连接RPC服务成功")
	var merc modes.Merchant
	var list modes.MerchantList
	merc.MerchantId = "b97c6afce65859df44c3b1b0acc64dd9"
	err = client.Call("Merchant.GetAllBranch", &merc, &list)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", list)
}