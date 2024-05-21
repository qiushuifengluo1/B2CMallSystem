/*
MainController 提供了主页面管理的基本功能，包括：

显示主页面 (Get)
显示欢迎页面 (Welcome)
修改公共状态 (ChangeStatus)
编辑数量字段 (EditNum)
这些方法共同构成了一个完整的主页面管理模块，便于在后台系统中管理用户权限和系统设置。
*/

package backend

import (
	"B2CProject/models"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
)

type MainController struct {
	beego.Controller
}

/*
Get 方法用于显示主页面。
从会话中获取用户信息并检查其角色权限。
根据角色权限加载相应的权限项，并将其标记为选中。
将权限列表和是否超级管理员的信息设置到模板数据中，渲染 index.html 模板。
*/
func (c *MainController) Get() {
	userinfo, ok := c.GetSession("userinfo").(models.Administrator)
	if ok {
		c.Data["username"] = userinfo.Username
		roleId := userinfo.RoleId
		auth := []models.Auth{}
		models.DB.Preload("AuthItem", func(db *gorm.DB) *gorm.DB {
			return db.Order("auth.sort DESC")
		}).Order("sort desc").Where("module_id=?", 0).Find(&auth)
		//获取当前部门拥有的权限，并把权限ID放在一个MAP对象里面
		roleAuth := []models.RoleAuth{}
		models.DB.Where("role_id=?", roleId).Find(&roleAuth)
		roleAuthMap := make(map[int]int)
		for _, v := range roleAuth {
			roleAuthMap[v.AuthId] = v.AuthId
		}
		for i := 0; i < len(auth); i++ {
			if _, ok := roleAuthMap[auth[i].Id]; ok {
				auth[i].Checked = true
			}
			for j := 0; j < len(auth[i].AuthItem); j++ {
				if _, ok := roleAuthMap[auth[i].AuthItem[j].Id]; ok {
					auth[i].AuthItem[j].Checked = true
				}
			}
		}
		c.Data["authList"] = auth
		c.Data["isSuper"] = userinfo.IsSuper
	}
	c.TplName = "backend/main/index.html"
}

func (c *MainController) Welcome() {
	c.TplName = "backend/main/welcome.html"
}

// ChangeStatus 修改公共状态
func (c *MainController) ChangeStatus() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "非法请求",
		}
		c.ServeJSON()
		return
	}
	table := c.GetString("table")
	field := c.GetString("field")
	err1 := models.DB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id=?", id).Error
	if err1 != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "更新数据失败",
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"msg":     "更新数据成功",
	}
	c.ServeJSON()
}

func (c *MainController) EditNum() {
	id := c.GetString("id")
	table := c.GetString("table")
	field := c.GetString("field")
	num := c.GetString("num")
	err1 := models.DB.Exec("update " + table + " set " + field + "=" + num + " where id=" + id).Error
	if err1 != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "修改数量失败",
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"msg":     "修改数量成功",
	}
	c.ServeJSON()
}
