package lib

import (
    "io"
    "fmt"
    "regexp"
    "crypto/md5"
)

/*
 * 描述: MD5加密
 *
 ***************************************************************************/
func StrMd5Str( strPass string )string{
    w := md5.New()
    io.WriteString( w, strPass )
    return fmt.Sprintf("%x", w.Sum(nil))
}

/*
 * 描述: 手机号识别
 *
 ***************************************************************************/
func IsPhone( strPhone string )bool{
    regular := "^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0-9])|(19[0-9])|(17[0-9]))\\d{8}$"
    reg := regexp.MustCompile(regular)
    return reg.MatchString( strPhone )
}

