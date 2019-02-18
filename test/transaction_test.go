package test

import (
	"fmt"
	"net/rpc"
	"public/server/modes"
	"time"
	"testing"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/2/17
 */
func TestTransaction(t *testing.T) {
	footAdd()
}

func footAdd() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
		return
	}
	fmt.Println("连接RPC服务成功")
	var tran modes.TransactionFoot
	tran.TranId = "t123456789"
	tran.UserId = "aa1212c4503999b2ce53e8dcc8eab98c"
	tran.CashierId = "bbbbbb"
	tran.Amount=1000.01
	tran.MerchantId = "b97c6afce65859df44c3b1b0acc64dd9"
	tran.CreateAt = time.Now().Unix()
	err = client.Call("TransactionFoot.Add", &tran, &tran)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Println("调用结果:", tran)
}
