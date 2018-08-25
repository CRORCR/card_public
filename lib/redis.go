package lib

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/astaxie/beego/cache/redis"
	"fmt"
	//"reflect"
	"github.com/astaxie/beego/cache"
	//"time"
	"errors"
)

type RedisConf struct {
	Host string `json:"redisHost"` //
	User string `json:"redisUser"` //
	Pwd  string `json:"redisPwd"`  //
	Port string `json:"redisPort"` //
}

func (this *RedisConf) init(configFile string) {
	jsonFile, err := os.Open(configFile)
	defer jsonFile.Close()
	if (err != nil) {
		panic("配置文件读取错误")
	}
	jsonData, err := ioutil.ReadAll(jsonFile)
	if (err != nil) {
		panic("配置文件读取失败")
	}
	json.Unmarshal(jsonData, &this)
}

func (this *RedisConf) GetRedis(configPath string) (cache.Cache, error) {
	this.init(configPath)
	var s = fmt.Sprintf("{\"key\":\"%s\",\"conn\":\"%s:%s\",\"dbNum\":\"%d\",\"password\":\"%s\"}", "redis", this.Host, this.Port, 0, "")
	res := redis.NewRedisCache()
	err := res.StartAndGC(s)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (this *RedisConf) GetRedisOnString(cache2 cache.Cache, key string) (string, error) {
	exists := cache2.IsExist(key)
	if !exists {
		fmt.Println(key , "not found..")
		return "", errors.New("not found")
	}
	b := cache2.Get(key)
	by := b.([]byte)
	return string(by), nil
}

//func main() {
//	redis, _ := getRedis()
//	redis.Put("test", 100, time.Second*200)
//}
