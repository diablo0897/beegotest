package main

import (
	_ "beegotest/routers"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/cache"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/beego/beego/v2/server/web/session/redis"
	"github.com/go-redis/redis"
)
var cc cache.Cache

var redisClient *redis.Client
func initRedis() {



	defer func() {
		if r := recover(); r != nil {
			cc = nil
		}
	}()

	host, _ := beego.AppConfig.String("cache::redis_host")
	password, _ := beego.AppConfig.String("cache::redis_passwd")
	config := make(map[string]interface{})
	config["conn"] = host
	if "" != password {
		config["password"] = password
	}

	config["key"] = "zgomart_test"
	configJson, _ := json.Marshal(config)
	fmt.Println("连接参数:",string(configJson))
	cc, _ = cache.NewCache("redis", string(configJson))

	redisClient = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     host,
		Password: password,
	})
}
func main() {
	initRedis()
	beego.Run()
}

