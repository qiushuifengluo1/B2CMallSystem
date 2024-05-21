/*
BannerController 提供了轮播图管理的基本功能，包括：

列出轮播图 (Get)
显示添加轮播图页面 (Add)
处理添加轮播图请求 (GoAdd)
显示编辑轮播图页面 (Edit)
处理编辑轮播图请求 (GoEdit)
删除轮播图 (Delete)
这些方法共同构成了一个完整的轮播图管理模块，便于在后台系统中管理轮播图。
*/

package backend

import (
	"B2CProject/common"
	"B2CProject/models"
	"github.com/astaxie/beego/logs"
	"os"
	"strconv"
)

type BannerController struct {
	BaseController
}

func (c *BannerController) Get() {
	banner := []models.Banner{}
	models.DB.Find(&banner)
	c.Data["bannerList"] = banner
	c.TplName = "backend/banner/index.html"
}

func (c *BannerController) Add() {
	c.TplName = "backend/banner/add.html"
}

func (c *BannerController) GoAdd() {
	bannerType, err1 := c.GetInt("banner_type")
	title := c.GetString("title")
	link := c.GetString("link")
	sort, err2 := c.GetInt("sort")
	status, err3 := c.GetInt("status")
	if err1 != nil || err3 != nil {
		c.Error("非法请求", "/banner")
		return
	}
	if err2 != nil {
		c.Error("排序表单里面输入的数据不合法", "/banner/add")
		return
	}
	bannerImgSrc, err4 := c.UploadImg("banner_img")
	if err4 == nil {
		banner := models.Banner{
			Title:      title,
			BannerType: bannerType,
			BannerImg:  bannerImgSrc,
			Link:       link,
			Sort:       sort,
			Status:     status,
			AddTime:    int(common.GetUnix()),
		}
		models.DB.Create(&banner)
		c.Success("增加轮播图成功", "/banner")
	} else {
		c.Error("增加轮播图失败", "/banner/add")
		return
	}
}

func (c *BannerController) Edit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("非法请求", "/banner")
		return
	}
	banner := models.Banner{Id: id}
	models.DB.Find(&banner)
	c.Data["banner"] = banner
	c.TplName = "backend/banner/edit.html"
}

/*
GoEdit 方法处理编辑轮播图的请求。
获取并验证表单输入的数据。
如果有新上传的轮播图图片，则更新图片路径。
更新轮播图记录并保存到数据库。
根据操作结果显示成功或错误消息。
*/
func (c *BannerController) GoEdit() {
	id, err := c.GetInt("id")
	bannerType, err1 := c.GetInt("banner_type")
	title := c.GetString("title")
	link := c.GetString("link")
	sort, err2 := c.GetInt("sort")
	status, err3 := c.GetInt("status")
	if err != nil || err1 != nil || err3 != nil {
		c.Error("非法请求", "/banner")
		return
	}
	if err2 != nil {
		c.Error("排序表单里面输入的数据不合法", "/banner/edit?id="+strconv.Itoa(id))
		return
	}
	bannerImgSrc, _ := c.UploadImg("banner_img")
	banner := models.Banner{Id: id}
	models.DB.Find(&banner)
	banner.Title = title
	banner.BannerType = bannerType
	banner.Link = link
	banner.Sort = sort
	banner.Status = status
	if bannerImgSrc != "" {
		banner.BannerImg = bannerImgSrc
	}
	err5 := models.DB.Save(&banner).Error
	if err5 != nil {
		c.Error("修改轮播图失败", "/banner/edit?id="+strconv.Itoa(id))
		return
	}
	c.Success("修改轮播图成功", "/banner")
}

func (c *BannerController) Delete() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/banner")
		return
	}
	banner := models.Banner{Id: id}
	models.DB.Find(&banner)
	address := "/home/qsfl/桌面/B2CProject/" + banner.BannerImg
	test := os.Remove(address)
	if test != nil {
		logs.Error(test)
		c.Error("删除物理机上图片错误", "/banner")
		return
	}
	models.DB.Delete(&banner)
	c.Success("删除轮播图成功", "/banner")
}
