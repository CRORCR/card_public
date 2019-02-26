package main

import (
	"fmt"
	"net/rpc"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/1/19
 */

type Merchant struct {
	Id           int64   `json:"id" xorm:"id"`                       //商家表ID
	FID          string  `json:"fid" xorm:"fid"`                     //父商家id
	MerchantId   string  `json:"merchant_id" xorm:"merchant_id"`     //本商家Id
	UserId       string  `json:"user_id" xorm:"user_id"`             //用户分享ID
	InviteCode   string  `json:"invite_code" xorm:"invite_code"`     //商家邀请码
	MerchantType int64   `json:"merchant_type" xorm:"merchant_type"` //商家 行业 类型
	MerchantRate float64 `json:"rate" xorm:"rate"`                   //商家 行业 利率
	//TrustStatus  bool    `json:"trust_status" xorm:"trust_status"`   //是否诺商家 0 否 1 是
	AreaNumber  int64   `json:"area_number" xorm:"area_number"` //市 I   D
	AreaId      int64   `json:"area_id" xorm:"area_id"`         //县 I   D
	CreateAt    int64   `json:"-" xorm:"create_at"`             //创 建 时 间
	CreateAtStr string  `json:"create_at" xorm:"-"`             //创 建 时 间
	Describea   string  `json:"describea" xorm:"describea"`     //描       述
	Address     string  `json:"address" xorm:"address"`         //地       址
	UserName    string  `json:"name" xorm:"name"`               //店铺名称
	Status      int64   `json:"status" xorm:"status"`           //状       态 //1 认证中 2 认证未通过 3 认证通过， 5 删除
	Phone       string  `json:"phone" xorm:"phone"`             //手  机   号
	Icon        string  `json:"icon" xorm:"icon"`               //商 家 头 像
	LoopImg     string  `json:"loopimg" xorm:"loopimg"`         //商家轮播图
	InfoImg     string  `json:"infoimg" xorm:"infoimg"`         //商家详情图
	Video       string  `json:"video" xorm:"video"`             //商家视频介绍
	CheckDesc   string  `json:"checkdesc" xorm:"checkdesc"`     //认证失败描述
	Business    string  `json:"business" xorm:"business"`       //营 业  执 照
	NumberId    string  `json:"number_id" xorm:"number_id"`     //身   份   证
	Industry    string  `json:"industry" xorm:"industry"`       //行 业许可 证
	Longitude   float64 `json:"longitude" xorm:"longitude"`     //经        度
	Latitude    float64 `json:"latitude" xorm:"latitude"`       //纬        度
	Cash        float64 `json:"cash" xorm:"cash"`               //现        金
	Trust       float64 `json:"trust" xorm:"trust"`             //鍩        分
	Credits     float64 `json:"credits" xorm:"credits"`         //积        分
	Distance    float64 `json:"distance" xorm:"-"`              //商家与用户的距离
}

type CoordinatesPoint struct {
	Longitude float64 //经       度
	Latitude  float64 //纬       度
	Page      int
	OfferSet  int
}

func main() {
	//addMer() //添加商家信息
	//getMer() //获得商家信息
	addBranch()
	//getMerStaff() //查询商家所有员工
}

//添加商家信息
func addMer() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")
	mer := Merchant{
		Phone:      "17600381201",
		MerchantId: "aaaa6cf8d45c8bef6e0c734bb02a635f",
		UserName:   "奇葩1号店",
		UserId:     "3dfa6cf8d45c8bef6e0c734bb02aaaaa",
		AreaNumber: 319,
		Longitude:  116.404,
		Latitude:   39.915,
	}
	err = client.Call("Merchant.Add", &mer, &mer)
	fmt.Printf("addMer 调用结果:%+v\n", mer)

	mer = Merchant{
		Phone:      "17600381203",
		MerchantId: "bbbb6cf8d45c8bef6e0c734bb02a635f",
		UserName:   "奇葩2号店",
		UserId:     "3dfa6cf8d45c8bef6e0c734bb02abbbb",
		AreaNumber: 319,
		Longitude:  116.404,
		Latitude:   39.915,
	}
	err = client.Call("Merchant.Add", &mer, &mer)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("addMer 调用结果:%+v\n", mer)
}

/* 根据商家id 获取商家   必填参数 商家id 必须使用rpc添加的用户才能知道区号*/
func getMer() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")

	mer := Merchant{
		MerchantId: "aaaa6cf8d45c8bef6e0c734bb02a635f",
	}
	mer2 := Merchant{}
	err = client.Call("Merchant.Get", &mer, &mer2)
	fmt.Printf("getMer 调用结果:%+v\n", mer2)

	mer = Merchant{
		MerchantId: "bbbb6cf8d45c8bef6e0c734bb02a635f",
	}
	err = client.Call("Merchant.Get", &mer, &mer2)
	fmt.Printf("调用结果:%+v\n", mer2)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("getMer 调用结果:%+v\n", mer2)
}


type MerchantAddBranch struct {
	Superior string `json:"superior"`// 上级ID, 总店
	Lower    string `json:"lower"`// 下级ID, 本店
}

//绑定
func addBranch() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")

	mer := MerchantAddBranch{Superior: "3f594b15fc30d04d78aa007f382acf13", Lower: "cccc0505d58a3980d1d6049feea62122"}
	mer2 := Merchant{}
	err = client.Call("Merchant.AddBranch", &mer, &mer2)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("addBranch 调用结果:%+v\n", mer2)
}


type Staff struct {
	Id         int    `json:"id" xorm:"id"`                   // id主键
	Name       string `json:"name" xorm:"name"`               // 员工姓名
	MerchantId string `json:"merchant_id" xorm:"merchant_id"` // 商 家 ID
	Phone      string `json:"phone" xorm:"phone"`             // 员工手机号
	UserId     string `json:"user_id" xorm:"user_id"`         // 员 工 ID
	NumberId   string `json:"number_id" xorm:"number_id"`     // 身份证号
	Sex        int    `json:"sex" xorm:"sex"`                 // 性    别
	CreateAt   int64  `json:"-" xorm:"create_at"`             // 创建时间
	State      int64  `json:"state" xorm:"state"`             // 状    态
	NumberFage int64  `json:"number_fage" xorm:"number_fage"` // 身份标识
	Authority  uint64 `json:"authority" xorm:"authority"`     // 权    限
	CreateStr  string `json:"create_at" xorm:"-"`             // 更新时间供前端展示
}

type StaffList []Staff

//获得商家所有员工   给定商家id
func getMerStaff() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	fmt.Println("连接RPC服务成功")

	var merc Merchant
	var mer StaffList
	merc.MerchantId = "aaaa6cf8d45c8bef6e0c734bb02a635f"
	err = client.Call("Merchant.GetStaff", &merc, &mer)
	fmt.Printf("getMerStaff 调用结果:%+v\n", mer)
	merc.MerchantId = "bbbb6cf8d45c8bef6e0c734bb02a635f"
	err = client.Call("Merchant.GetStaff", &merc, &mer)
	fmt.Printf("getMerStaff 调用结果:%+v\n", mer)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
}