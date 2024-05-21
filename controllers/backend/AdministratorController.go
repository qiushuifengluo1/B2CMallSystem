/*
AdministratorController 提供了管理员管理的基本功能，包括：
列出管理员列表 (Get)
显示添加管理员页面 (Add)
处理添加管理员请求 (GoAdd)
显示编辑管理员页面 (Edit)
处理编辑管理员请求 (GoEdit)
删除管理员 (Delete)
这些方法共同构成了一个完整的管理员管理模块，便于在后台系统中对管理员进行管理。
*/

package backend

import (
	"B2CProject/common"
	"B2CProject/models"
	"strconv"
	"strings"
)

// AdministratorController 嵌入 BaseController 结构体，以便使用其中的方法。
type AdministratorController struct {
	BaseController
}

func (c *AdministratorController) Get() {
	administrator := []models.Administrator{}
	models.DB.Preload("Role").Find(&administrator)
	//将管理员列表数据设置到 c.Data["administratorList"]，并渲染模板 index.html。
	c.Data["administratorList"] = administrator
	c.TplName = "backend/administrator/index.html"
}

func (c *AdministratorController) Add() {
	role := []models.Role{}
	models.DB.Find(&role)
	c.Data["roleList"] = role
	c.TplName = "backend/administrator/add.html"
}


/*
GoAdd 方法处理添加管理员的请求。
验证输入数据的合法性（用户名、密码长度和邮箱格式）。
检查用户名是否已存在。
创建并保存新的管理员记录。
根据操作结果显示成功或错误消息。
*/
func (c *AdministratorController) GoAdd() {
	username := strings.Trim(c.GetString("username"), "")
	password := strings.Trim(c.GetString("password"), "")
	mobile := strings.Trim(c.GetString("mobile"), "")
	email := strings.Trim(c.GetString("email"), "")
	roleId, err1 := c.GetInt("role_id")
	if err1 != nil {
		c.Error("非法请求", "/administrator/add")
	}
	if len(username) < 2 || len(password) < 6 {
		c.Error("用户名或密码长度不合法", "/administrator/add")
		return
	} else if common.VerifyEmail(email) == false {
		c.Error("邮箱格式不正确，请重新填写!", "/administrator/add")
		return
	}
	administratorList := []models.Administrator{}
	models.DB.Where("username=?", username).Find(&administratorList)
	if len(administratorList) > 0 {
		c.Error("用户名已存在", "/administrator/add")
		return
	}

	administrator := models.Administrator{}
	administrator.Username = username
	administrator.Password = common.Md5(password)
	administrator.Mobile = mobile
	administrator.Email = email
	administrator.Status = 1
	administrator.AddTime = int(common.GetUnix())
	administrator.RoleId = roleId
	err := models.DB.Create(&administrator).Error
	if err != nil {
		c.Error("增加管理员失败", "/administrator/add")
		return
	}
	c.Success("增加管理员成功", "/administrator")
}

//Edit 方法用于显示编辑管理员页面。
func (c *AdministratorController) Edit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/administrator")
		return
	}
	administrator := models.Administrator{Id: id}
	models.DB.Find(&administrator)
	//根据传入的管理员 ID 获取管理员信息，并设置到 c.Data["administrator"]。
	c.Data["administrator"] = administrator
	//获取所有角色信息并设置到 c.Data["roleList"]，渲染模板 edit.html
	role := []models.Role{}
	models.DB.Find(&role)
	c.Data["roleList"] = role
	c.TplName = "backend/administrator/edit.html"
}

/*
GoEdit 方法处理编辑管理员的请求。
验证输入数据的合法性（密码长度和邮箱格式）。
更新管理员记录。
根据操作结果显示成功或错误消息。
*/
func (c *AdministratorController) GoEdit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/administrator")
		return
	}
	username := strings.Trim(c.GetString("Username"), "")
	password := strings.Trim(c.GetString("Password"), "")
	mobile := strings.Trim(c.GetString("Mobile"), "")
	email := strings.Trim(c.GetString("Email"), "")
	roleId, err1 := c.GetInt("role_id")
	if err1 != nil {
		c.Error("非法请求", "/administrator")
		return
	}
	if password != "" {
		if len(password) < 6 {
			c.Error("密码长度不合法！", "/administrator/add?id="+strconv.Itoa(id))
			return
		} else if common.VerifyEmail(email) == false {
			c.Error("邮箱格式不正确，请重新填写!", "/administrator/add?id="+strconv.Itoa(id))
			return
		}
		password = common.Md5(password)
	}
	administrator := models.Administrator{Id: id}
	models.DB.Find(&administrator)
	administrator.Username = username
	administrator.Password = password
	administrator.Mobile = mobile
	administrator.Email = email
	administrator.RoleId = roleId
	err2 := models.DB.Save(&administrator).Error
	if err2 != nil {
		c.Error("修改管理员失败", "/administrator/edit?id="+strconv.Itoa(id))
	} else {
		c.Success("修改管理员成功", "/administrator")
	}
}

//Delete 方法处理删除管理员的请求。
func (c *AdministratorController) Delete() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/role")
		return
	}
	administrator := models.Administrator{Id: id}
	models.DB.Delete(&administrator)
	c.Success("删除管理员成功", "/administrator")
}
