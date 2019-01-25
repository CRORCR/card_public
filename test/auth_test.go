package test

import (
	"fmt"
	"testing"
)

const (
	NONE = 1 << iota
	//提现
	WITHDRAW
	//兑换
	EXCHANGE
	//二次确认
	SECOND
	//查看微名片
	FINDMER
	//添加修改微名片
	ADDMER
)

type User struct {
	Name string
	Tag  int64
}

/**
 * @desc   权限设置
 * @author Ipencil
 * @create 2019/1/25
 */
func TestAuth(t *testing.T) {
	//withdraw() //提现权限
	//exchange() //兑换权限
	//second()   //二次确认
	//findmer()  //查看微名片
	//addmer()   //添加更新微名片
	//组合测试
	commint()
}

func commint(){
	user := User{}
	user.AddAuth(ADDMER)
	//user.AddAuth(EXCHANGE)
	user.AddAuth(SECOND)
	//user.AddAuth(FINDMER)
	user.AddAuth(ADDMER)

	fmt.Println(user.IsAuth(ADDMER))
	fmt.Println(user.IsAuth(EXCHANGE))
	fmt.Println(user.IsAuth(SECOND))
	fmt.Println(user.IsAuth(FINDMER))
	fmt.Println(user.IsAuth(ADDMER))
}

func addmer() {
	user := User{}
	user.AddAuth(ADDMER)
	user.DelAuth(ADDMER)
}

func findmer() {
	user := User{}
	user.AddAuth(FINDMER)
	user.DelAuth(FINDMER)
}

func second() {
	user := User{}
	user.AddAuth(SECOND)
	user.DelAuth(SECOND)
}

func exchange() {
	user := User{}
	user.AddAuth(EXCHANGE)
	user.DelAuth(EXCHANGE)
}

func withdraw() {
	user := User{}
	fmt.Println("是否有提现权限", user.IsAuth(WITHDRAW))
	user.AddAuth(WITHDRAW)
	fmt.Println("添加提现权限")
	fmt.Println("是否有提现权限", user.IsAuth(WITHDRAW))
	user.DelAuth(WITHDRAW)
	fmt.Println("去除提现权限")
	fmt.Println("是否有提现权限", user.IsAuth(WITHDRAW))
}

func (user *User) AddAuth(auth int64) {
	user.Tag = user.Tag | auth
}

func (user *User) DelAuth(auth int64) {
	user.Tag = user.Tag ^ auth
}

func (user *User) IsAuth(auth int64) bool {
	result := user.Tag & auth
	return result == auth
}