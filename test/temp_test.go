package test

import (
	"card_public/server/modes"
	"fmt"
	"net/rpc"
	"testing"
)

/**
 * @desc    TODO
 * @author Ipencil
 * @create 2019/3/7
 */
func TestTemp(t *testing.T) {
	t.Run("tempGet", tempGet)  //模板获取
	t.Run("tempSet", tempSet) //模板设置
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