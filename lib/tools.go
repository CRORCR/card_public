package lib

import (
	"crypto/md5"
	"fmt"
	"io"
	"reflect"
	"regexp"
)

/*
 * 描述: MD5加密
 *
 ***************************************************************************/
func StrMd5Str(strPass string) string {
	w := md5.New()
	io.WriteString(w, strPass)
	return fmt.Sprintf("%x", w.Sum(nil))
}

/*
 * 描述: 手机号识别
 *
 ***************************************************************************/
func IsPhone(strPhone string) bool {
	regular := "^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0-9])|(19[0-9])|(17[0-9]))\\d{8}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(strPhone)
}

/*
 * 描述: struct 转换 map
 *
 ***************************************************************************/
func ToMap(obj interface{}) map[string]interface{} {

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
