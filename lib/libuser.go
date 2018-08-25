package lib

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type UserLib struct {
	HeadIcon	string	`json:"UserHeadImage"`		// 用户默认头像
	MerchantIcon	string	`json:"MerchantIcon"`		// 商家默认头像
}

var g_user UserLib

func GetUserLib()*UserLib{
	return &g_user
}

func ( this *UserLib )InitUser( strPathName string ){
	jsonFile , err := os.Open( strPathName )
	if err != nil {
	        panic("打开文件错误，请查看:" + strPathName )
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll( jsonFile )
	if err != nil {
		panic("读取文件错误:" + strPathName )
	}
	json.Unmarshal( jsonData, &this )
	fmt.Println("用户默认头像:", this.HeadIcon )
	fmt.Println("商家默认头像:", this.MerchantIcon )
}

