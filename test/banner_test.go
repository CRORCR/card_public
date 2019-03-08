package test

import (
	"card_public/server/modes"
	"fmt"
	"net/rpc"
	"testing"
	"time"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/3/8
 */
 func TestBanner(t *testing.T) {
	t.Run("addBanner", addBanner)  //添加广告
	t.Run("downLoad", downLoad)  //添加广告
}

/*添加一个广告*/
func addBanner(t *testing.T){
	t.SkipNow()
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println()
		fmt.Println("连接RPC服务失败:", err)
	}
	temp :=modes.Banner{}
	temp.AreaId=111
	temp.BannerSite="轮播图1"
	temp.PayTime=time.Now().Unix()
	temp.BannerStatus=1
	temp.TotalTimes=1000
	err = client.Call("Banner.AddBanner", &temp, &temp)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", temp)
}

/*阅览广告*/
func downLoad(t *testing.T){
	//t.SkipNow()
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println()
		fmt.Println("连接RPC服务失败:", err)
	}
	temp :=modes.Banner{}
	temp.AreaId=111
	temp.BannerSite="轮播图1"
	err = client.Call("Banner.DownShow", &temp, &temp)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", temp)
}