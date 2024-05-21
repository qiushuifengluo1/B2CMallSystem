/*
该 ProductTypeAttrController 提供了完整的商品类型属性管理功能，包括列出、添加、编辑和删除商品类型属性。
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

type ProductTypeAttrController struct {
	BaseController
}

func (c *ProductTypeAttrController) Get() {

	cateId, err1 := c.GetInt("cate_id")
	if err1 != nil {
		c.Error("非法请求", "/productType")
	}
	//获取当前的类型
	productType := models.ProductType{Id: cateId}
	models.DB.Find(&productType)
	c.Data["productType"] = productType

	//查询当前类型下面的商品类型属性
	productTypeAttr := []models.ProductTypeAttribute{}
	models.DB.Where("cate_id=?", cateId).Find(&productTypeAttr)
	c.Data["productTypeAttrList"] = productTypeAttr

	c.TplName = "backend/productTypeAttribute/index.html"
}

func (c *ProductTypeAttrController) Add() {

	cateId, err1 := c.GetInt("cate_id")
	if err1 != nil {
		c.Error("非法请求", "/productType")
	}

	productType := []models.ProductType{}
	models.DB.Find(&productType)
	c.Data["productTypeList"] = productType
	c.Data["cateId"] = cateId
	c.TplName = "backend/productTypeAttribute/add.html"
}

func (c *ProductTypeAttrController) GoAdd() {

	title := c.GetString("title")
	cateId, err1 := c.GetInt("cate_id")
	attrType, err2 := c.GetInt("attr_type")
	attrValue := c.GetString("attr_value")
	sort, err4 := c.GetInt("sort")
	if err1 != nil || err2 != nil {
		c.Error("非法请求", "/productType")
		return
	}
	if strings.Trim(title, " ") == "" {
		c.Error("商品类型属性名称不能为空", "/productTypeAttribute/add?cate_id="+strconv.Itoa(cateId))
		return
	}
	if err4 != nil {
		c.Error("排序值错误", "/productTypeAttribute/add?cate_id="+strconv.Itoa(cateId))
		return
	}
	productTypeAttr := models.ProductTypeAttribute{
		Title:     title,
		CateId:    cateId,
		AttrType:  attrType,
		AttrValue: attrValue,
		Status:    1,
		AddTime:   int(common.GetUnix()),
		Sort:      sort,
	}
	models.DB.Create(&productTypeAttr)
	c.Success("增加成功", "/productTypeAttribute?cate_id="+strconv.Itoa(cateId))

}

func (c *ProductTypeAttrController) Edit() {
	id, err1 := c.GetInt("id")
	if err1 != nil {
		c.Error("非法请求", "/goodType")
		return
	}
	productTypeAttr := models.ProductTypeAttribute{Id: id}
	models.DB.Find(&productTypeAttr)
	c.Data["productTypeAttr"] = productTypeAttr
	productType := []models.ProductType{}
	models.DB.Find(&productType)
	c.Data["productTypeList"] = productType
	c.TplName = "backend/productTypeAttribute/edit.html"
}

func (c *ProductTypeAttrController) GoEdit() {
	id, err := c.GetInt("id")
	title := c.GetString("title")
	cateId, err1 := c.GetInt("cate_id")
	attrType, err2 := c.GetInt("attr_type")
	attrValue := c.GetString("attr_value")
	sort, err4 := c.GetInt("sort")
	if err != nil || err1 != nil || err2 != nil {
		c.Error("非法请求", "/productTypeAttribute")
		return
	}
	if strings.Trim(title, " ") == "" {
		c.Error("商品类型属性名称不能为空", "/productTypeAttribute/edit?cate_id="+strconv.Itoa(id))
		return
	}
	if err4 != nil {
		c.Error("排序值错误", "/productTypeAttribute/edit?cate_id="+strconv.Itoa(id))
		return
	}
	productTypeAttr := models.ProductTypeAttribute{Id: id}
	models.DB.Find(&productTypeAttr)
	productTypeAttr.Title = title
	productTypeAttr.CateId = cateId
	productTypeAttr.AttrType = attrType
	productTypeAttr.AttrValue = attrValue
	productTypeAttr.Sort = sort
	err3 := models.DB.Save(&productTypeAttr).Error
	if err3 != nil {
		c.Error("修改数据失败", "/productTypeAttribute/edit?cate_id="+strconv.Itoa(id))
	}
	c.Success("修改数据成功", "/productTypeAttribute?cate_id="+strconv.Itoa(cateId))
}
func (c *ProductTypeAttrController) Delete() {
	id, err := c.GetInt("id")
	cateId, err1 := c.GetInt("cate_id")
	if err != nil {
		c.Error("传入参数错误", "/productTypeAttribute?cate_id="+strconv.Itoa(cateId))
		return
	}
	if err1 != nil {
		c.Error("非法请求", "/productType")
	}
	productTypeAttr := models.ProductTypeAttribute{Id: id}
	models.DB.Delete(&productTypeAttr)
	c.Success("删除数据成功", "/productTypeAttribute?cate_id="+strconv.Itoa(cateId))
}
