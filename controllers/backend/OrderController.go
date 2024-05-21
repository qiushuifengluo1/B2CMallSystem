/*
OrderController 提供了订单管理的基本功能，包括：

列出订单 (Get)
显示订单详情 (Detail)
显示编辑订单页面 (Edit)
处理编辑订单请求 (GoEdit)
删除订单 (Delete)
这些方法共同构成了一个完整的订单管理模块，便于在后台系统中管理订单。
*/

package backend

import (
	"B2CProject/models"
	"math"
	"strconv"
)

type OrderController struct {
	BaseController
}

/*
Get 方法用于获取订单列表。
处理分页逻辑，根据当前页码和每页显示的数量从数据库中获取订单数据。
计算总页数，并将订单列表和分页信息设置到模板数据中。
渲染模板 order.html。
*/
func (c *OrderController) Get() {
	page, _ := c.GetInt("page")
	if page == 0 {
		page = 1
	}
	pageSize := 5
	keyword := c.GetString("keyword")
	order := []models.Order{}
	var count int
	if keyword == "" {
		models.DB.Table("order").Count(&count)
		models.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&order)
	} else {
		models.DB.Where("phone=?", keyword).Offset((page - 1) * pageSize).Limit(pageSize).Find(&order)
		models.DB.Where("phone=?", keyword).Table("order").Count(&count)
	}
	c.Data["totalPages"] = math.Ceil(float64(count) / float64(pageSize))
	c.Data["page"] = page
	c.Data["order"] = order
	c.TplName = "backend/order/order.html"
}

func (c *OrderController) Detail() {
	c.Ctx.WriteString("详情页面")
}

func (c *OrderController) Edit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/order")
		return
	}
	order := models.Order{}
	models.DB.Where("id=?", id).Find(&order)
	c.Data["order"] = order
	c.TplName = "backend/order/edit.html"
}

/*
GoEdit 方法处理编辑订单的请求。
获取并验证表单输入的数据。
更新订单记录并保存到数据库。
根据操作结果显示成功或错误消息。
*/
func (c *OrderController) GoEdit() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/order")
		return
	}
	orderId := c.GetString("order_id")
	allPrice := c.GetString("all_price")
	name := c.GetString("name")
	phone := c.GetString("phone")
	address := c.GetString("address")
	zipcode := c.GetString("zipcode")
	payStatus, _ := c.GetInt("pay_status")
	payType, _ := c.GetInt("pay_type")
	orderStatus, _ := c.GetInt("order_status")
	order := models.Order{}
	models.DB.Where("id=?", id).Find(&order)
	order.OrderId = orderId
	order.AllPrice, _ = strconv.ParseFloat(allPrice, 64)
	order.Name = name
	order.Phone = phone
	order.Address = address
	order.Zipcode = zipcode
	order.PayStatus = payStatus
	order.PayType = payType
	order.OrderStatus = orderStatus
	models.DB.Save(&order)
	c.Success("订单修改成功", "/order")
}

func (c *OrderController) Delete() {
	id, err := c.GetInt("id")
	if err != nil {
		c.Error("传入参数错误", "/order")
		return
	}
	order := models.Order{}
	models.DB.Where("id=?", id).Delete(&order)
	c.Success("删除订单记录成功", "/order")
}
