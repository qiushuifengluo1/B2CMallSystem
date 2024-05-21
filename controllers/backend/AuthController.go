/*
AuthController 提供了权限管理的基本功能，包括：
列出权限列表 (Get)
显示添加权限页面 (Add)
处理添加权限请求 (GoAdd)
显示编辑权限页面 (Edit)
处理编辑权限请求 (GoEdit)
删除权限 (Delete)
这些方法共同构成了一个完整的权限管理模块，便于在后台系统中对权限进行管理。
*/

package backend

import (
	"B2CProject/models"
	"strconv"
)

type AuthController struct {
	BaseController
}

//Get 方法用于获取权限列表，并预加载其子权限项。
func (c *AuthController) Get() {
	auth := []models.Auth{}
	models.DB.Preload("AuthItem").Where("module_id=0").Find(&auth)
	//将权限列表数据设置到 c.Data["authList"]，并渲染模板 index.html。
	c.Data["authList"] = auth
	c.TplName = "backend/auth/index.html"
}

//Add 方法用于显示添加权限的页面
func (c *AuthController) Add() {
	auth := []models.Auth{}
	models.DB.Where("module_id=0").Find(&auth)
	//获取所有顶级模块权限，并设置到 c.Data["authList"]，渲染模板 add.html。
	c.Data["authList"] = auth
	c.TplName = "backend/auth/add.html"
}

/*
GoAdd 方法处理添加权限的请求。
验证输入数据的合法性。
创建并保存新的权限记录。
根据操作结果显示成功或错误消息。
*/
func (c *AuthController) GoAdd() {
	moduleName := c.GetString("module_name")
	iType, err1 := c.GetInt("type")
	actionName := c.GetString("action_name")
	url := c.GetString("url")
	moduleId, err2 := c.GetInt("module_id")
	sort, err3 := c.GetInt("sort")
	description := c.GetString("description")
	status, err4 := c.GetInt("status")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		c.Error("传入参数错误", "/auth/add")
		return
	}
	auth := models.Auth{
		ModuleName:  moduleName,
		Type:        iType,
		ActionName:  actionName,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
	}
	err := models.DB.Create(&auth).Error
	if err != nil {
		c.Error("增加数据失败", "/auth/add")
		return
	}
	c.Success("增加数据成功", "/auth")
}

//Edit 方法用于显示编辑权限的页面
func (c *AuthController) Edit() {
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("传入参数错误", "/auth")
		return
	}
	auth := models.Auth{Id: id}
	models.DB.Find(&auth)
	//根据传入的权限 ID 获取权限信息，并设置到 c.Data["auth"]。
	c.Data["auth"] = auth
	authList := []models.Auth{}
	models.DB.Where("module_id=0").Find(&authList)
	//获取所有顶级模块权限，并设置到 c.Data["authList"]，渲染模板 edit.html。
	c.Data["authList"] = authList
	c.TplName = "backend/auth/edit.html"
}

/*
GoEdit 方法处理编辑权限的请求。
验证输入数据的合法性。
更新权限记录。
根据操作结果显示成功或错误消息。
*/
func (c *AuthController) GoEdit() {
	moduleName := c.GetString("module_name")
	iType, err1 := c.GetInt("type")
	actionName := c.GetString("action_name")
	url := c.GetString("url")
	moduleId, err2 := c.GetInt("module_id")
	sort, err3 := c.GetInt("sort")
	description := c.GetString("description")
	status, err4 := c.GetInt("status")
	id, err5 := c.GetInt("id")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
		c.Error("传入参数错误", "/auth")
		return
	}
	auth := models.Auth{Id: id}
	models.DB.Find(&auth)
	auth.ModuleName = moduleName
	auth.Type = iType
	auth.ActionName = actionName
	auth.Url = url
	auth.ModuleId = moduleId
	auth.Sort = sort
	auth.Description = description
	auth.Status = status
	err6 := models.DB.Save(&auth).Error
	if err6 != nil {
		c.Error("修改权限失败", "/auth/edit?id="+strconv.Itoa(id))
		return
	}
	c.Success("修改权限成功", "/auth")
}

/*
Delete 方法处理删除权限的请求。
根据传入的权限 ID 删除权限记录。
如果删除的是顶级模块权限，首先检查其是否包含子权限，若有子权限则不允许删除。
显示删除成功消息。
*/
func (c *AuthController) Delete() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/role")
		return
	}
	auth := models.Auth{Id: id}
	models.DB.Find(&auth)
	if auth.ModuleId == 0 {
		auth2 := []models.Auth{}
		models.DB.Where("module_id=?", auth.Id).Find(&auth2)
		if len(auth2) > 0 {
			c.Error("请删除当前顶级模块下面的菜单或操作！", "/auth")
			return
		}
	}
	models.DB.Delete(&auth)
	c.Success("删除成功", "/auth")
}
