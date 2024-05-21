/*
SettingController 提供了系统设置管理的基本功能，包括：
查看系统设置 (Get)
处理编辑系统设置的请求 (GoEdit)
这些方法共同构成了一个完整的系统设置管理模块，便于在后台系统中对系统设置进行管理。
*/

package backend

import "B2CProject/models"

type SettingController struct {
	BaseController
}

func (c *SettingController) Get() {
	setting := models.Setting{}
	models.DB.First(&setting)
	c.Data["setting"] = setting
	c.TplName = "backend/setting/index.html"
}

func (c *SettingController) GoEdit() {
	setting := models.Setting{}
	models.DB.Find(&setting)
	err := c.ParseForm(&setting)
	if err != nil {
		return
	}
	siteLogo, err := c.UploadImg("site_logo")
	if len(siteLogo) > 0 && err == nil {
		setting.SiteLogo = siteLogo
	}
	noPicture, err := c.UploadImg("no_picture")
	if len(noPicture) > 0 && err == nil {
		setting.NoPicture = noPicture
	}
	err = models.DB.Where("id=1").Save(&setting).Error
	if err != nil {
		c.Error("修改数据失败", "/setting")
		return
	}
	c.Success("修改数据成功", "/setting")
}

