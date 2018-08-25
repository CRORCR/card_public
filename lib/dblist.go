package lib

import (
    "os"
    "fmt"
    "io/ioutil"
    "encoding/json"
)

type DBList struct{
    DBUser  string `json:"db_user"`
    DBHome  string `json:"db_home"`
    DBPort  uint32 `json:"db_port"`
    DBName  string `json:"db_name"`
    DBPass  string `json:"db_pass"`
}

func ReadDBConfig( strName string )[]DBList{
    var dblist []DBList

    fmt.Println(strName)
    jsonFile , err := os.Open(strName)
    if err != nil {
        panic("打开文件错误，请查看:" + strName )
    }
    defer jsonFile.Close()

    jsonData, err := ioutil.ReadAll( jsonFile )
    if err != nil {
        panic("读取文件错误:" + strName )
    }

    json.Unmarshal( jsonData, &dblist )
    return dblist
}

/*
func main(){
    dblist := ReadDBConfig("../config/db.json")
    fmt.Println(dblist)
}
*/
