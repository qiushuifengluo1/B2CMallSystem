/*
LoginController 提供了用户登录管理的基本功能，包括：
显示登录页面 (Get)
处理登录请求 (GoLogin)
处理退出登录请求 (LoginOut)
这些方法共同构成了一个完整的用户登录管理模块，便于在后台系统中对用户进行身份验证和会话管理。
*/

package backend

import (
	"B2CProject/common"
	"B2CProject/models"
	"strings"
)

type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	c.TplName = "backend/login/login.html"
}

func (c *LoginController) GoLogin() {
	//验证码验证
	var flag = models.Cpt.VerifyReq(c.Ctx.Request)
	if flag {
		username := strings.Trim(c.GetString("username"), "")
		password := common.Md5(strings.Trim(c.GetString("password"), ""))
		administrator := []models.Administrator{}
		models.DB.Where("username=? AND password=? AND status=1", username, password).Find(&administrator)
		if len(administrator) == 1 {
			c.SetSession("userinfo", administrator[0])
			c.Success("登陆成功", "/")
		} else {
			c.Error("无登陆权限或用户名密码错误", "/login")
		}
	} else {
		c.Error("验证码错误", "/login")
	}
}

func (c *LoginController) LoginOut() {
	c.DelSession("userinfo")
	c.Success("退出登录成功,将返回登陆页面！", "/login")
}
