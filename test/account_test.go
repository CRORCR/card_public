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
 * @create 2019/2/22
 */

func TestAccount(t *testing.T) {
	findAccount() //查询商家所有员工
}

//添加商家信息
func findAccount() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}

	sql := fmt.Sprintf("merchant_id in (%v)", 33)
	mer := modes.WithdrawalWhere{}
	mer.Where = sql
	mer.Count = 20
	mer.Page = 0 //偏移量
	out := make([]*modes.WithdrawalAccount, 0)
	err = client.Call("WithdrawalAccount.QueryWhere", &mer, &out)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", out)
}