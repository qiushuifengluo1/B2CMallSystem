// Cookie模型
package models

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

type cookie struct{}

// Set 写入数据的方法
func (c cookie) Set(ctx *context.Context, key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	ctx.SetSecureCookie(beego.AppConfig.String("secureCookie"), key, string(bytes), 3600*24*30, "/", beego.AppConfig.String("domain"), nil, true)

}

// Remove 删除数据的方法
func (c cookie) Remove(ctx *context.Context, key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	ctx.SetSecureCookie(beego.AppConfig.String("secureCookie"), key, string(bytes), -1, "/", beego.AppConfig.String("domain"), nil, true)

}

// Get 获取数据的方法
func (c cookie) Get(ctx *context.Context, key string, obj interface{}) bool {
	tempData, ok := ctx.GetSecureCookie(beego.AppConfig.String("secureCookie"), key)
	if !ok {
		return false
	}
	err := json.Unmarshal([]byte(tempData), obj)
	if err != nil {
		return false
	}
	return true

}

// Cookie 实例化结构体
var Cookie = &cookie{}
