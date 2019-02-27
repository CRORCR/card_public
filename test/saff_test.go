package test

import (
	"fmt"
	"net/rpc"
	"card_public/server/modes"
	"testing"
	"time"
)

/**
 * @desc    员工测试
 * @author Ipencil
 * @create 2019/1/20
 */
func TestStaff(t *testing.T) {
	addStaff()
	//getStaff()
	//delStaff()

	//updateStaff()
	//getMerId()
	//addAuthority()
	//getAuthorityOfStaff()
	//getAreaNumber()
	//getUserId()

	//addAuthority()
	//showAuthority()
	//cancelAuthority()
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
		MerchantId: "dpfa6cf8d45c8bef6e0c734bb02a635f", UserId: "dcba95a8ee243cc1c440e4e52bff3384", Name: "店员3", Phone: "15156813885",
		Sex:        1, CreateAt: time.Now().Unix(), State: 0, NumberFage: 0, Authority: 1}, 310,}
	var s modes.Staff
	err = client.Call("Staff.Add", &add, &s)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", s)
}

//删除员工
func delStaff() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")
	add := modes.Staff{UserId: "53792f5ad648dfd57f4ff8efba5e3c76"}
	var s modes.Staff
	err = client.Call("Staff.Del", &add, &s)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", s)
}

//查询员工
func getStaff() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")
	sta := modes.Staff{UserId: "dcba95a8ee243cc1c440e4e52bff3384"}
	out := modes.Staff{}
	err = client.Call("Staff.Get", &sta, &out)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", out) //22
}

//更新员工信息
func updateStaff() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")
	add := modes.AddStaff{modes.Staff{
		MerchantId: "22", UserId: "22yuangogn", Name: "yg", Phone: "17600381284",
		Sex:        1}, 310,}
	var s modes.Staff
	err = client.Call("Staff.Update", &add, &s)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", s)
}

//更新员工信息
func getUserId() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")
	var userId string
	var phone = "13145213417"

	err = client.Call("Staff.GetUserId", &phone, &userId)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", userId)
}

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

//增加权限
func addAuthority() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")
	add := modes.StaffAuthority{UserId: "000", Fage: 4} //增加收银权限
	var s modes.Staff
	err = client.Call("Staff.SetAuthority", &add, &s)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", s)
}

//查询权限
func showAuthority() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")
	adds := modes.StaffAuthority{UserId: "000", Fage: 4}
	var s bool
	err = client.Call("Staff.ShowAuthority", &adds, &s)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", s)
}

//取消权限
func cancelAuthority() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")
	add := modes.Staff{UserId: "000", Authority: 3} //增加收银权限
	adds := modes.StaffAuthority{UserId: "000", Fage: 3}
	err = client.Call("Staff.CancelAuthority", &adds, &add)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", add)
}

//生成收款码
func getAuthorityOfStaff() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")
	add := modes.Staff{UserId: "000", Authority: 1}
	var s string
	err = client.Call("Staff.GetQRCode", &add, &s)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", s) //{"mid":"22","uid":"000"}
}

//测试员工属于哪个区
func getAreaNumber() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")
	var input = "1ee395a8ee243cc1c440e4e52bff3382"
	var output int64
	err = client.Call("Staff.GetAreaNumber", &input, &output)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", output)
}
