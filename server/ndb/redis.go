package ndb

import (
	"os"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/cache/redis"
)

type RedisConf struct {
	Host string `json:"redisHost"` //
	User string `json:"redisUser"` //
	Pwd  string `json:"redisPwd"`  //
	Port string `json:"redisPort"` //
	Rndb cache.Cache
}

var Rndb *RedisConf

func GetRndbHand()*RedisConf{
	return Rndb
}

func Init( strConfigFile string )error{
	return Rndb.GetRedis( strConfigFile )
}

func (this *RedisConf)Init(configFile string) {
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

func (this *RedisConf)GetRedis(configPath string) error {
	this.Init(configPath)
	var s = fmt.Sprintf("{\"key\":\"redis\",\"conn\":\"%s:%s\",\"dbNum\":\"%d\",\"password\":\"%s\"}",
				this.Host,
				this.Port,
				0,
				"")
	fmt.Println("Redis Info:", s )
	this.Rndb = redis.NewRedisCache()
	return this.Rndb.StartAndGC(s)
}

//func main() {
//	redis, _ := getRedis()
//	redis.Put("test", 100, time.Second*200)
//}
