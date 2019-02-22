package test

import (
	"fmt"
	"net/rpc"
	"public/server/modes"
	"time"
)

type RPCServer struct {
	IP   string `json:"rpc_ip"`
	Type string `json:"rpc_type"`
	Rpc  *rpc.Client
}

func MerchantGet() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
		return
	}
	fmt.Println("连接RPC服务成功")
	var merc modes.Merchant
	merc.MerchantId = "123456789_319"
	err = client.Call("Merchant.Get", &merc, &merc)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", merc)
}

func MerchantAdd() {
	// 123456789 123456789_311 123456789_319
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
		return
	}
	fmt.Println("连接RPC服务成功")
	var merc modes.Merchant
	var mera modes.Merchant
	merc.Phone = "19803091863"
	merc.MerchantId = "123456789"
	merc.UserId = "aaaaaa"
	merc.UserName = "shane"
	merc.AreaNumber = 319
	merc.Longitude = 116.404
	merc.Latitude = 39.915
	//merc.Rate = 8.6
	err = client.Call("Merchant.Add", &merc, &mera)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", mera)
}

func MerchantAddBranch() {
	// 123456789 123456789_311 123456789_319
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
		return
	}
	fmt.Println("连接RPC服务成功")
	var merc modes.Merchant
	var branch modes.MerchantAddBranch
	branch.Superior = "123456789"
	branch.Lower = "123456789_311"
	err = client.Call("Merchant.AddBranch", &branch, &merc)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", merc)
}

func MerchantGetAllBranch() {
	// 123456789 123456789_311 123456789_319
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
		return
	}
	fmt.Println("连接RPC服务成功")
	var merc modes.Merchant
	var list modes.MerchantList
	merc.MerchantId = "123456789"
	err = client.Call("Merchant.GetAllBranch", &merc, &list)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", list)
}
func MerchantAskIdentity() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
		return
	}
	fmt.Println("连接RPC服务成功")
	var staff modes.Staff
	var fage bool
	staff.UserId = "aaaaaa" // aaaaaa
	err = client.Call("Staff.AskIdentity", &staff, &fage)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", fage)
}

func  StaffCancelAuthority() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
		return
	}
	fmt.Println("连接RPC服务成功")
	var auth  modes.StaffAuthority
	var fage modes.Staff
	auth.UserId = "aaaaaa" // aaaaaa
	auth.Fage   = 3 // aaaaaa
	err = client.Call("Staff.CancelAuthority", &auth, &fage )
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("%.64b\n", fage.Authority )
	fmt.Println("调用结果:", fage)
}

func  StaffShowAuthority() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
		return
	}
	fmt.Println("连接RPC服务成功")
	var auth  modes.StaffAuthority
	var fage bool
	auth.UserId = "aaaaaa" // aaaaaa
	auth.Fage   = 1 // aaaaaa
	err = client.Call("Staff.ShowAuthority", &auth, &fage )
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	//fmt.Printf("%.64b\n", fage.Authority )
	fmt.Println("调用结果:", fage)
}

func  StaffSetAuthority() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
		return
	}
	fmt.Println("连接RPC服务成功")
	var auth  modes.StaffAuthority
	var fage modes.Staff
	auth.UserId = "aaaaaa" // aaaaaa
	auth.Fage   = 3 // aaaaaa
	err = client.Call("Staff.SetAuthority", &auth, &fage )
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("%.64b\n", fage.Authority )
	fmt.Println("调用结果:", fage)
}

func  StaffGetUserId() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
		return
	}
	fmt.Println("连接RPC服务成功")
	strPhone := "19803091863"
	var fage string
	err = client.Call("Staff.GetUserId", &strPhone, &fage )
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Println("调用结果:", fage)
}

func TransactionFootAdd(){
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
		return
	}
	fmt.Println("连接RPC服务成功")
	var tran modes.TransactionFoot
	tran.TranId = "t123456789"
	tran.UserId = "s123456789"
	tran.CashierId = "bbbbbb"
	tran.MerchantId = "123456788"
	tran.CreateAt = time.Now().Unix()
	err = client.Call("TransactionFoot.Add", &tran, &tran )
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Println("调用结果:", tran)
}

func TransactionFootMerchantGetAll(){
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
		return
	}
	fmt.Println("连接RPC服务成功")
	var tran modes.TransactionList
	var mid modes.TransactionInfo
	mid.Id = "123456789"
	mid.Count = 15
	mid.Page = 0
	err = client.Call("TransactionFoot.MerchantGetAll", &mid, &tran )
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Println("调用结果:", tran)
}

func TransactionFootUserGetAll(){
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
		return
	}
	fmt.Println("连接RPC服务成功")
	var tran modes.TransactionList
	var mid = "s123456789"
	err = client.Call("TransactionFoot.UserGetAll", &mid, &tran )
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Println("调用结果:", tran)
}

func MerchantGetTarNumber(){
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
		return
	}
	fmt.Println("连接RPC服务成功")
	var tran modes.TarData
	tran.MerchantId = "123456789"
	tran.Amount     = 123456789
	var mid string
	err = client.Call("Merchant.GetTarNumber", &tran, &mid )
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Println("调用结果:", mid)
}

func TransactionFootAll(){
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
		return
	}
	fmt.Println("连接RPC服务成功")
	//var tran modes.TransactionList
	var mid = "s123456789"
	err = client.Call("TransactionFoot.All", &mid, nil )
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Println("调用结果:")
}


func main() {
	TransactionFootAll()
	//MerchantGetTarNumber()
	//StaffGetUserId()
	//TransactionFootAdd()
	//TransactionFootMerchantGetAll()
	//TransactionFootUserGetAll()
	//MerchantAdd()
	//StaffCancelAuthority()
	//StaffShowAuthority()
	//StaffSetAuthority()
	//MerchantAddBranch()
	//MerchantGet()
	//MerchantGetAllBranch()
	//MerchantAskIdentity()
	//var unid = modes.UnionId{"shane1234567890", "unionid_android", "19803066666"}
	/*client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
		return
	}
	fmt.Println("连接RPC服务成功")
	*/
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
			    fmt.Println("调用失败:", err)
			}
		        fmt.Printf("调用结果:%+v\n", merc )
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
				fmt.Println("调用失败:", err)
			}
			fmt.Printf("调用结果:%+v\n", staff )
		}
	*/
	/*
			var merc modes.Merchant
		        var sil  modes.StaffList
		        merc.MerchantId = "123456789"
		        err = client.Call("Merchant.GetStaff", &merc, &sil )
		        if err != nil {
		                fmt.Println("调用失败:", err)
		        }
		        fmt.Printf("调用结果:%+v\n", sil )

	*/
	/*
			var staff modes.Staff
			staff.UserId  = "user_id"
			var strGetQRCode string
			err = client.Call("Staff.GetQRCode", &staff, &strGetQRCode )
		        if err != nil {
		                fmt.Println("调用失败:", err)
		        }
		        fmt.Printf("调用结果:%+v\n", strGetQRCode )
	*/
}