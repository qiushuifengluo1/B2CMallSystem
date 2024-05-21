/*
这段代码定义了一个 `V1Controller` 控制器，属于 `api` 包，并使用 Beego 框架。控制器包含两个方法：

1. `Get` 方法：处理 HTTP GET 请求，返回字符串 "api v1"。
2. `Menu` 方法：处理 HTTP 请求，查询数据库中的 `Menu` 表数据，并以 JSON 格式返回。

主要步骤：
- `Get` 方法返回简单的字符串响应。
- `Menu` 方法查询数据库，获取所有 `Menu` 记录，并以 JSON 格式返回响应。
*/

package api

import (
	"B2CProject/models"
	"github.com/astaxie/beego"
)

type V1Controller struct {
	beego.Controller
}

func (c *V1Controller) Get() {
	c.Ctx.WriteString("api v1")
}

func (c *V1Controller) Menu() {
	menu := []models.Menu{}
	models.DB.Find(&menu)
	c.Data["json"] = menu
	c.ServeJSON()
}
