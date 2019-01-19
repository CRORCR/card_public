package test

import (
    "fmt"
    "net/rpc"
	"testing"
	"public/server/modes"
)

type RPCServer struct{
    IP      string `json:"rpc_ip"`
    Type    string `json:"rpc_type"`
    Rpc    *rpc.Client
}

func TestOther(t *testing.T) {
    //var unid = modes.UnionId{"shane1234567890", "unionid_android", "19803066666"}
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
		return
	}
	fmt.Println("连接RPC服务成功")
	/*
	var merc modes.Merchant
	merc.Phone = "19803091863"
	merc.MerchantId = "123456789"
	merc.UserId = "SHANE"
	merc.AreaNumber = 310
	merc.Longitude = 116.404
	merc.Latitude = 39.915
	err = client.Call("Merchant.Add", &merc, &merc)
	if err != nil {
	    fmt.Println("调用失败：", err)
	}
        fmt.Printf("调用结果：%+v\n", merc )
	*/
	/*
	for i := 0; i< 10; i++ {
		var addStaff modes.AddStaff
		var staff modes.Staff
		staff.Name = fmt.Sprintf("name_%d", i)			// 员工姓名
		staff.MerchantId = "123456789"	// 商 家 ID
		staff.Phone   = "phone"			// 员工手机号
		staff.UserId  = "user_id"		// 员 工 ID
		staff.NumberId = "number_id"		// 身份证号
		staff.Sex  = true			// 性    别
		staff.CreateAt = 15423658222		// 创建时间
		staff.State  = 0			// 状    态
		staff.NumberFage = 1			// 身份标识 
		staff.Authority  =  255			// 权    限
		staff.CreateStr  = "create_at"          // 更新时间供前端展示

		addStaff.PStaff = staff
		addStaff.AreaNumber = 310

		err = client.Call("Staff.Add", &addStaff, nil )
		if err != nil {
			fmt.Println("调用失败：", err)
		}
		fmt.Printf("调用结果：%+v\n", staff )
	}
	*/

	var merc modes.Merchant
        var sil  modes.StaffList
        merc.MerchantId = "123456789"
        err = client.Call("Merchant.GetStaff", &merc, &sil )
        if err != nil {
                fmt.Println("调用失败：", err)
        }
        fmt.Printf("调用结果：%+v\n", sil )


	var staff modes.Staff
	staff.UserId  = "user_id"
	err = client.Call("Staff.Get", &staff, &staff )
        if err != nil {
                fmt.Println("调用失败：", err)
        }
        fmt.Printf("调用结果：%+v\n", staff )
}
