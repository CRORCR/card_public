package test

import (
	"fmt"
	"net/rpc"
	"card_public/server/modes"
)

import (
	"testing"
)

/**
 * @desc    提现记录表
 * @author Ipencil
 * @create 2019/2/20
 */

func TestMW(t *testing.T) {
	//mwithList()
	//getSuccess()
}

func mwithList() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")

	mer := modes.WithdrawalWhere{}
	mer.Where = "wit_id = 11111"
	mer.Page = 0   //偏移量
	mer.Count = 20 //多少条记录

	out := make([]modes.MWithdrawalFoot, 0)
	err = client.Call("MWithdrawalFoot.QueryWhere", &mer, &out)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", out[0])
}

func getSuccess() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")

	mer := modes.MWithdrawalFoot{}
	mer.MerchantId="33"

	var out float64
	err = client.Call("MWithdrawalFoot.GetSuccess", &mer, &out)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", out)
}
