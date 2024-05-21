/*
	该代码实现了后台管理系统的权限控制，确保只有经过认证且具有适当权限的用户才能访问相应的管理界面。未登录用户会被重定向到登录页面，

没有权限的用户会收到相应的提示。
*/
package common

import (
	"B2CProject/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"net/url"
	"strings"
)

// BackendAuth 后台权限判断
func BackendAuth(ctx *context.Context) { //接受一个context.Context类型的参数，用于处理HTTP请求和响应。
	pathname := ctx.Request.URL.String()
	//userinfo尝试从会话中获取userinfo对象，并将其转换为models.Administrator类型。
	userinfo, ok := ctx.Input.Session("userinfo").(models.Administrator)
	if !(ok && userinfo.Username != "") {
		if pathname != "/"+beego.AppConfig.String("adminPath")+"/login" &&
			pathname != "/"+beego.AppConfig.String("adminPath")+"/login/gologin" &&
			pathname != "/"+beego.AppConfig.String("adminPath")+"/login/verificode" {
			//检查当前请求路径是否是登录页面或登录相关的路径，如果不是，则重定向到登录页面
			ctx.Redirect(302, "/"+beego.AppConfig.String("adminPath")+"/login")
		}
	} else {
		pathname = strings.Replace(pathname, "/"+beego.AppConfig.String("adminPath"), "", 1)
		urlPath, _ := url.Parse(pathname)
		//如果用户不是超级管理员（userinfo.IsSuper == 0），并且请求路径不在排除权限检查的路径列表中，则进行权限检查。
		if userinfo.IsSuper == 0 && !excludeAuthPath(urlPath.Path) {
			roleId := userinfo.RoleId //获取用户的角色ID，并根据角色ID查询该角色拥有的权限。
			roleAuth := []models.RoleAuth{}
			models.DB.Where("role_id=?", roleId).Find(&roleAuth)
			roleAuthMap := make(map[int]int)
			for _, v := range roleAuth {
				roleAuthMap[v.AuthId] = v.AuthId
			}
			auth := models.Auth{}
			models.DB.Where("url=?", urlPath.Path).Find(&auth)
			if _, ok := roleAuthMap[auth.Id]; !ok {
				ctx.WriteString("没有权限")
				return
			}
		}
	}
}

// 检验路径权限
func excludeAuthPath(urlPath string) bool {
	excludeAuthPathSlice := strings.Split(beego.AppConfig.String("excludeAuthPath"), ",")
	for _, v := range excludeAuthPathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}
