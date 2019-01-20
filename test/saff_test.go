package test

import (
	"fmt"
	"net/rpc"
	"public/server/modes"
	"testing"
	"time"
)

/**
 * @desc    员工测试
 * @author Ipencil
 * @create 2019/1/20
 */
func TestStaff(t *testing.T) {
	//addStaff()
	//getStaff()
	//getMerId()
}

//添加员工
func addStaff() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")
	add := modes.AddStaff{modes.Staff{
		MerchantId: "22", UserId: "22yuangogn", Name: "yaungong", Phone: "17600381284",
		Sex:        false, CreateAt: time.Now().Unix(), State: 0, NumberFage: 0, Authority: 1}, 310,}
	var s modes.Staff
	err = client.Call("Staff.Add", &add, &s)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", s)
}

func getStaff(){
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")
	sta := modes.Staff{UserId: "22yuangogn"}
	err = client.Call("Staff.Get", &sta, &sta)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", sta) //22
}
/*
{Id:2 Name:yaungong MerchantId:22 Phone:17600381284 UserId:22yuangogn NumberId: Sex:false
CreateAt:1547950866 State:0 NumberFage:0 Authority:1 CreateStr:}
*/

func getMerId() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")
	sta := modes.Staff{UserId: "22yuangogn"}
	var s string
	err = client.Call("Staff.GetMerchantId", &sta, &s)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", s) //22
}
