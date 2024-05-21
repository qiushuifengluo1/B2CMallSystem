/*
该 ProductTypeController 提供了完整的商品类型管理功能，包括列出、添加、编辑和删除商品类型。
使用 models.DB 进行数据库操作，结合 GORM 的特性简化操作。
通过继承 BaseController 提供的基础功能，如错误处理和成功提示，简化了控制器的实现。
渲染对应的模板页面，为前端提供数据支持。
*/

package backend

import (
	"B2CProject/common"
	"B2CProject/models"
	"strconv"
	"strings"
)

type ProductTypeController struct {
	BaseController
}

func (c *ProductTypeController) Get() {
	productType := []models.ProductType{}
	models.DB.Find(&productType)
	c.Data["productTypeList"] = productType
	c.TplName = "backend/productType/index.html"
}

func (c *ProductTypeController) Add() {
	c.TplName = "backend/productType/add.html"
}

func (c *ProductTypeController) GoAdd() {
	title := strings.Trim(c.GetString("title"), "")
	description := strings.Trim(c.GetString("description"), "")
	status, err := c.GetInt("status")
	if title == "" {
		c.Error("标题不能为空", "/productType/add")
		return
	}
	if err != nil {
		c.Error("传入参数不正确", "/productType/add")
		return
	}
	productTypeList := []models.ProductType{}
	models.DB.Where("title=?", title).Find(&productTypeList)
	if len(productTypeList) != 0 {
		c.Error("该商品已存在！", "/productType/add")
		return
	}
	productType := models.ProductType{}
	productType.Title = title
	productType.Description = description
	productType.Status = status
	productType.AddTime = int(common.GetUnix())
	err1 := models.DB.Create(&productType).Error
	if err1 != nil {
		c.Error("增加商品类型失败", "/productType/add")
	} else {
		c.Success("增加商品类型成功", "/productType")
	}
}

func (c *ProductTypeController) Edit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/productType")
		return
	}
	productType := models.ProductType{Id: id}
	models.DB.Find(&productType)
	c.Data["productType"] = productType
	c.TplName = "backend/productType/edit.html"
}

func (c *ProductTypeController) GoEdit() {
	title := strings.Trim(c.GetString("title"), "")
	description := strings.Trim(c.GetString("description"), "")
	status, err := c.GetInt("status")
	id, err1 := c.GetInt("id")
	if err != nil || err1 != nil {
		c.Error("传入参数错误", "/productType")
		return
	}
	if title == "" {
		c.Error("标题不能为空", "/productType/edit?id="+strconv.Itoa(id))
		return
	}
	productType := models.ProductType{Id: id}
	models.DB.Find(&productType)
	productType.Title = title
	productType.Description = description
	productType.Status = status
	err2 := models.DB.Save(&productType).Error
	if err2 != nil {
		c.Error("修改商品类型失败", "/productType/edit?id="+strconv.Itoa(id))
	} else {
		c.Success("修改商品类型成功", "/productType")
	}
}

func (c *ProductTypeController) Delete() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/productType")
		return
	}
	productType := models.ProductType{Id: id}
	models.DB.Delete(&productType)
	c.Success("删除数据成功", "/productType")
}
