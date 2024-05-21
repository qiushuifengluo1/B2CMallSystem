// Redis缓存模型
package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
	"time"
)

var redisClient cache.Cache
var enableRedis, _ = beego.AppConfig.Bool("enableRedis")
var redisTime, _ = beego.AppConfig.Int("redisTime")
var YzmClient cache.Cache

func init() {
	if enableRedis {
		config := map[string]string{
			"key":      beego.AppConfig.String("redisKey"),
			"conn":     beego.AppConfig.String("redisConn"),
			"dbNum":    beego.AppConfig.String("redisDbNum"),
			"password": beego.AppConfig.String("redisPwd"),
		}
		bytes, _ := json.Marshal(config)
		//创建Redis客户端,创建Redis缓存对象
		redisClient, err = cache.NewCache("redis", string(bytes))
		YzmClient, _ = cache.NewCache("redis", string(bytes))
		if err != nil {
			logs.Error("连接redis数据库失败")
		} else {
			logs.Info("连接redis数据库成功")
		}

	}
}

type cacheDb struct{}

var CacheDb = &cacheDb{} // 创建一个全局的缓存对象,用于调用缓存方法

// Set 写入数据到Redis缓存的方法
func (c cacheDb) Set(key string, value interface{}) {
	if enableRedis {
		bytes, _ := json.Marshal(value)
		err := redisClient.Put(key, string(bytes), time.Second*time.Duration(redisTime))
		if err != nil {
			return
		}
	}
}

// Get 从Redis缓存中获取数据的方法
func (c cacheDb) Get(key string, obj interface{}) bool {
	if enableRedis {
		if redisStr := redisClient.Get(key); redisStr != nil {
			fmt.Println("在redis里面读取数据...")
			redisValue, ok := redisStr.([]uint8)
			if !ok {
				fmt.Println("获取redis数据失败")
				return false
			}
			err := json.Unmarshal([]byte(redisValue), obj)
			if err != nil {
				return false
			}
			return true
		}
		return false
	}
	return false
}
