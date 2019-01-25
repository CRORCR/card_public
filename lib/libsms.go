package lib

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"time"
	//"io/ioutil"
	//"encoding/json"
)

const (
	SMS_TIMEOUT int64 = 300 // 短信超时时长5分钟
)

type SmsQueue struct {
	strPhone string //
	nFage    uint8  //
	nTime    int64  //
	nCode    uint32 //
}

var g_smsqueue []SmsQueue

func RandInt64(min, max int64) int64 {
	maxBigInt := big.NewInt(max)
	i, _ := rand.Int(rand.Reader, maxBigInt)
	if i.Int64() < min {
		RandInt64(min, max)
	}
	return i.Int64()
}

func getConfig(nfage uint8) string {
	if 0 == nfage {
		return "SMS_131345043"
	}
	return "SMS_131345043"
}

/*
 * 描述：发送短信验证码
 *
 * strPhone 	: 发送的手机号
 * nFage	: 来源标志
 * nTime	: 发放时间戳
 *
 *****************************************************************************************/
func SendSMS(strPhone string, nFage uint8, nTime int64) uint8 {

	nNumber := RandInt64(100000, 999999)
	strNumber := strconv.FormatInt(nNumber, 10)
	fmt.Println(strPhone, "  ", nNumber)
	err := SendMsg(strPhone, getConfig(nFage), strNumber)
	if err == nil {
		g_smsqueue = append(g_smsqueue, SmsQueue{strPhone, nFage, nTime, uint32(nNumber)})
		fmt.Println("SMSQueue:", g_smsqueue)
		return 0
	}
	fmt.Println("err", err)
	return 1
}

/*
 * 描述：验证短信验证码
 *
 * strPhone 	: 发送的手机号
 * nCode	: 短信验证码
 *
 * return 	: 0 正确，1 手机号不存在， 2 超时，3 验证码不正确
 *
 *****************************************************************************************/
func CheckSMS(strPhone string, nCode uint32) uint8 {
	// 从队列中删除数据的条件是，超时 和验证码正确。
	fmt.Println("Phone", strPhone, " checkcode ", nCode)
	nTime := time.Now().Unix()
	nLen := len(g_smsqueue)
	fmt.Println("dd8888888888888888888888", g_smsqueue)
	var nFage uint8 = 9
	var i int
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaa")
	for i := nLen - 1; i >= 0; i-- {
		fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbbbbb")
		fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbbbbb", g_smsqueue[i].strPhone)
		if g_smsqueue[i].strPhone == strPhone {
			if (nTime - g_smsqueue[i].nTime) > SMS_TIMEOUT {
				fmt.Println("cccccccccccccccccccccccccccccccccccccccc")
				nFage = 2 //验证码超时
				break
			}
			if g_smsqueue[i].nCode == nCode {
				fmt.Println("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
				nFage = 1 // 验证码正确
				break
			}
		}
	}
	switch nFage {
	case 2:
		fmt.Println("333333333333333333333333333333333333333333")
		g_smsqueue = append(g_smsqueue[:i], g_smsqueue[i+1:]...)
		return 2
	case 1:
		fmt.Println("2222222222222222222222222222222222222222")
		g_smsqueue = append(g_smsqueue[:i], g_smsqueue[i+1:]...)
		return 0

	}
	fmt.Println("11111111111111111111111111111111111111")
	return 1
}
