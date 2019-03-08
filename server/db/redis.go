package db

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"io/ioutil"
	"os"
)

type RedisServer struct {
	IPPort   string `json:"ip_port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

var g_RedisHand *redis.Client

func GetRedis() *redis.Client {
	return g_RedisHand
}

func (this *RedisServer) ReadConfigFile(strName string) error {
	jsonFile, err := os.Open(strName)
	if err != nil {
		panic("打开文件错误，请查看:" + strName)
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic("读取文件错误:" + strName)
	}
	return json.Unmarshal(jsonData, this)
}

func (this *RedisServer) Start(strName string) error {
	err := this.ReadConfigFile(strName)
	if err != nil {
		return err
	}
	g_RedisHand = redis.NewClient(&redis.Options{
		Addr:     this.IPPort,
		Password: this.Password, // no password set
		DB:       this.DB,       // use default DB
	})
	pong, err := g_RedisHand.Ping().Result()
	fmt.Println(pong, err)
	return err
}
