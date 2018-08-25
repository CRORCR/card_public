package lib

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
	"time"
	"net/url"
	"sort"
	"crypto/sha1"
	"strings"
	"crypto/hmac"
	"encoding/base64"
	"net/http"
	"errors"
	"github.com/satori/go.uuid"
)

type smsCode struct {
	Code      string
	Message   string
	HostId    string
	Recommend string
}

type SmsSystemParam struct {
	AccessKeyId      string `json:"keyid"`
	AccSecret        string `json:"keysecret"`
	Format           string `json:"Format"`           //没传默认为JSON，可选填值：XML
	SignatureMethod  string `json:"SignatureMethod"`  //建议固定值：HMAC-SHA1
	SignatureVersion string `json:"SignatureVersion"` //建议固定值：1.0
	SignatureNonce   string                           //用于请求的防重放攻击，每次请求唯一，JAVA语言建议用：java.util.UUID.randomUUID()生成即可
	Timestamp        string                           //格式为：yyyy-MM-dd’T’HH:mm:ss’Z’；时区为：GMT
	Signature        string                           //最终生成的签名结果值
}

type SmsApplicationParam struct {
	Action        string `json:"Action"`   //API的命名，固定值，如发送短信API的值为：SendSms
	Version       string `json:"Version"`  //API的版本，固定值，如短信API的值为：2017-05-25
	RegionId      string `json:"RegionId"` //API支持的RegionID，如短信API的值为：cn-hangzhou
	PhoneNumbers  string                   //短信接收号码,支持以逗号分隔的形式进行批量调用，批量上限为1000个手机号码,批量调用相对于单条调用及时性稍有延迟,验证码类型的短信推荐使用单条调用的方式；发送国际/港澳台消息时，接收号码格式为00+国际区号+号码，如“0085200000000”
	SignName      string                   //短信签名
	TemplateCode  string                   //短信模板ID，发送国际/港澳台消息时，请使用国际/港澳台短信模版
	TemplateParam string                   //短信模板变量替换JSON串,友情提示:如果JSON中需要带换行符,请参照标准的JSON协议。 {“code”:”1234”,”product”:”ytx”}
	OutId         string                   //外部流水扩展字段
}

func (this *SmsSystemParam) ReadFile(strName string) {
	jsonFile, err := os.Open(strName)
	defer jsonFile.Close()
	if (err != nil) {
		panic("文件打开错误" + strName);
		log.Fatal(err)
	}
	jsonData, err := ioutil.ReadAll(jsonFile)
	if (err != nil) {
		panic("文件读取错误" + strName)
		log.Fatal(err)
	}
	json.Unmarshal(jsonData, &this)
}

func (this *SmsSystemParam) InitParameter(configPath string) {
	var zone = time.FixedZone("GMT", 0)
	noce, _ := uuid.NewV4()
	this.ReadFile(configPath)
	this.Timestamp = fmt.Sprintf("%s", time.Now().In(zone).Format("2006-01-02T03:04:05Z"))
	this.SignatureNonce = noce.String()
}

func (this *SmsApplicationParam) readConfig(configPath string) {
	jsonFile, err := os.Open(configPath)
	defer jsonFile.Close()
	if (err != nil) {
		panic("文件打开错误" + configPath);
		log.Fatal(err)
	}
	jsonData, err := ioutil.ReadAll(jsonFile)
	if (err != nil) {
		panic("文件读取错误" + configPath)
		log.Fatal(err)
	}
	json.Unmarshal(jsonData, &this)
}

func Sing(sys *SmsSystemParam, app *SmsApplicationParam) (string, string) {
	var keys []string
	sinArr := make(map[string]string)
	sinArr["AccessKeyId"] = sys.AccessKeyId
	sinArr["Timestamp"] = sys.Timestamp
	sinArr["Format"] = sys.Format
	sinArr["SignatureMethod"] = sys.SignatureMethod
	sinArr["SignatureVersion"] = sys.SignatureVersion
	sinArr["SignatureNonce"] = sys.SignatureNonce
	sinArr["Action"] = app.Action
	sinArr["Version"] = app.Version
	sinArr["RegionId"] = app.RegionId
	sinArr["SignName"] = app.SignName
	sinArr["TemplateCode"] = app.TemplateCode
	sinArr["PhoneNumbers"] = app.PhoneNumbers
	sinArr["TemplateParam"] = app.TemplateParam

	for i := range sinArr {
		keys = append(keys, i)
	}
	sort.Strings(keys)
	var sortQueryString string

	for _, v := range keys {
		sortQueryString = fmt.Sprintf("%s&%s=%s", sortQueryString, urlencode(v), urlencode(sinArr[v]))
	}
	stringToSign := fmt.Sprintf("GET&%s&%s", urlencode("/"), urlencode(sortQueryString[1:]))
	sinStr2 := sing(stringToSign, sys.AccSecret)
	return sinStr2, sortQueryString
}

func sing(urlstr string, secret string) (string) {
	mac := hmac.New(sha1.New, []byte(fmt.Sprintf("%s&", secret)))
	mac.Write([]byte(urlstr))
	singn := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return urlencode(singn)
}

func urlencode(str string) string {
	rep := strings.NewReplacer("+", "%20", "*", "%2A", "%7E", "~")
	return rep.Replace(url.QueryEscape(str))
}

func SendMsg(phoneNumber string, smsTmpName string, smsCodeParam string) (error) {
	var smsSys SmsSystemParam
	var smsApp SmsApplicationParam
	var configPath string = "config/sms.json"
	var sign string
	var apiURL = "http://dysmsapi.aliyuncs.com/?Signature="
	smsSys.InitParameter(configPath)
	smsApp.readConfig(configPath)
	smsApp.PhoneNumbers = phoneNumber
	smsApp.TemplateCode = smsTmpName
	smsApp.TemplateParam = "{code:'" + smsCodeParam + "'}"
	sign, httpparams := Sing(&smsSys, &smsApp)
	//send
	client := &http.Client{}
	var req *http.Request
	req, err := http.NewRequest("GET", apiURL+sign+"&"+httpparams, strings.NewReader("test"))
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("HTTP GET ERRROR :", err)
		return err
	}
	defer resp.Body.Close()
	var bodys []byte
	bodys, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("http readall body error ", err)
		return err
	}
	//fmt.Println(string(bodys))
	var res smsCode
	json.Unmarshal(bodys, &res)
	//fmt.Println(res.Code)
	if (res.Code != "OK") {
		return errors.New(res.Code)
	}
	return nil
}

//func main() {
//	err := sendMsg("17601028625", "SMS_131345043", "1234")
//	if (err != nil) {
//		fmt.Println(err)
//	} else {
//		fmt.Println("message already sended ..")
//	}
//}
