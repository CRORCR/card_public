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
	t.Run("addBanner", addBanner) //添加广告
	t.Run("downLoad", downLoad)   //阅览广告
	t.Run("findBanner", findBanner)   //查询广告
	t.Run("updateBanner", updateBanner)   //更新广告数据
	t.Run("QueryBannerShowInfo", QueryBannerShowInfo)   //查询指定区域,指定广告位
	//模板测试
	t.Run("tempGet", tempGet)   //模板获取
	t.Run("tempSet", tempSet)   //模板设置
}

/*添加一个广告*/
func addBanner(t *testing.T) {
	t.SkipNow()
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close()}()
	if err != nil {
		fmt.Println()
		fmt.Println("连接RPC服务失败:", err)
	}
	temp := modes.Banner{}
	temp.AreaId = 111
	temp.BannerSite = "轮播图1"
	temp.PayTime = time.Now().Unix()
	temp.BannerStatus = 1
	temp.TotalTimes = 1000
	err = client.Call("Banner.AddBanner", &temp, &temp)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", temp)
}

/*阅览广告*/
func downLoad(t *testing.T) {
	//t.SkipNow()
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println()
		fmt.Println("连接RPC服务失败:", err)
	}
	temp := modes.Banner{}
	temp.AreaId = 130322
	temp.BannerSite = "轮播1"
	err = client.Call("Banner.DownShow", &temp, &temp)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", temp)
}

/*根据条件查询历史记录*/
func findBanner(t *testing.T) {
	t.SkipNow()
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println()
		fmt.Println("连接RPC服务失败:", err)
	}
	temp := modes.Where{}
	temp.OffSet=0
	temp.Sum=2
	temp.SQL="id>0"
	result:=modes.ResultBanner{}
	err = client.Call("Banner.FindBanner", &temp, &result)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Println("total",result.Total)
	fmt.Println("total",result.Error)
	for _, value := range result.BannerResultList {
		fmt.Printf("调用结果:%+v\n", value)
	}
}

/*根据条件查询历史记录*/
func updateBanner(t *testing.T) {
	t.SkipNow()
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println()
		fmt.Println("连接RPC服务失败:", err)
	}
	temp := modes.Banner{}
	temp.AreaId = 112
	temp.ID=1
	temp.BannerSite = "轮播图3"
	err = client.Call("Banner.UpdateBanner", &temp, &temp)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
}

/*查询指定县,指定广告位数据*/
func QueryBannerShowInfo(t *testing.T) {
	t.SkipNow()
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println()
		fmt.Println("连接RPC服务失败:", err)
	}
	temp := modes.BannerShow{}
	temp2:=make([]*modes.BannerShow,0)
	temp.AreaID=111
	err = client.Call("Banner.QueryBannerShowInfo", &temp, &temp2)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	for _, value := range temp2 {
		fmt.Printf("%+v\n",value)
	}
}

/*模板获取*/
func tempGet(t *testing.T){
	t.SkipNow()
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	temp :=modes.Temp{}
	err = client.Call("TemplateBanner.Get", &temp, &temp)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", temp)
}

/*模板设置*/
func tempSet(t *testing.T){
	t.SkipNow()
	client, err := rpc.Dial("tcp", "127.0.0.1:7003")
	defer func() { client.Close() }()
	if err != nil {
		fmt.Println("连接RPC服务失败:", err)
	}
	temps :=modes.Temp{}
	temp:=modes.TemplateBanner{}
	temp.Name="轮播图1"
	temp.Pri=make([]modes.TempPrice,0)
	pri:=modes.TempPrice{}
	pri.Count=10
	pri.Price=110
	temp.Pri=append(temp.Pri,pri)
	pri.Count=11
	pri.Price=100
	temp.Pri=append(temp.Pri,pri)
	temps.Url="https://192.17"
	temps.Temps=append(temps.Temps,temp)
	temp.Name="轮播图2"
	temp.Pri[0].Count=20
	temp.Pri[0].Price=20
	temps.Temps=append(temps.Temps,temp)

	err = client.Call("TemplateBanner.Set", &temps, &temps)
	if err != nil {
		fmt.Println("调用失败:", err)
	}
	fmt.Printf("调用结果:%+v\n", temp)
}