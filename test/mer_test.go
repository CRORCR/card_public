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

func TestMer(t *testing.T) {
	setMargin()
}

type CoordinatesPoint struct {
	Longitude float64 //经       度
	Latitude  float64 //纬       度
	Page int
	OfferSet int
}

func setMargin() {
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() {client.Close()}()
	if err != nil {
		fmt.Println("连接RPC服务失败：", err)
	}
	fmt.Println("连接RPC服务成功")
	var margin = CoordinatesPoint{114.5315555000, 36.6479443200,1,4}
	mer:=make(modes2.MerchantList,0)
	err = client.Call("Merchant.GetNearMerchant", &margin, &mer)
	if err != nil {
		fmt.Println("调用失败：", err)
	}
	fmt.Printf("调用结果:%+v\n", mer)
}