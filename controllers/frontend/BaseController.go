/*
   BaseController 通过 BaseInit 方法初始化了一些通用的数据，包括：顶部导航菜单,左侧分类,中间导航菜单,用户登录信息.这些数据通
过缓存机制（Redis）进行存储和读取，以提高系统性能。根据用户是否登录，显示不同的用户信息。最终，这些数据会被传递到视图模板进行渲染，
生成最终的前端页面。
*/

package frontend

import (
	"B2CProject/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/jinzhu/gorm"
	"net/url"
	"strings"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) BaseInit() {
	//获取顶部导航
	topMenu := []models.Menu{}
	if hasTopMenu := models.CacheDb.Get("topMenu", &topMenu); hasTopMenu == true {
		c.Data["topMenuList"] = topMenu
	} else {
		models.DB.Where("status=1 AND position=1").Order("sort desc").Find(&topMenu)
		c.Data["topMenuList"] = topMenu
		models.CacheDb.Set("topMenu", topMenu)
	}

	//左侧分类（预加载）
	productCate := []models.ProductCate{}

	if hasProductCate := models.CacheDb.Get("productCate", &productCate); hasProductCate == true {
		c.Data["productCateList"] = productCate
	} else {
		models.DB.Preload("ProductCateItem", func(db *gorm.DB) *gorm.DB {
			return db.Where("product_cate.status=1").
				Order("product_cate.sort DESC")
		}).Where("pid=0 AND status=1").Order("sort desc", true).
			Find(&productCate)
		c.Data["productCateList"] = productCate
		models.CacheDb.Set("productCate", productCate)
	}

	//获取中间导航的数据
	middleMenu := []models.Menu{}
	if hasMiddleMenu := models.CacheDb.Get("middleMenu", &middleMenu); hasMiddleMenu == true {
		c.Data["middleMenuList"] = middleMenu
	} else {
		models.DB.Where("status=1 AND position=2").Order("sort desc").
			Find(&middleMenu)

		for i := 0; i < len(middleMenu); i++ {
			//获取关联商品
			middleMenu[i].Relation = strings.ReplaceAll(middleMenu[i].Relation, "，", ",")
			relation := strings.Split(middleMenu[i].Relation, ",")
			product := []models.Product{}
			models.DB.Where("id in (?)", relation).Limit(6).Order("sort ASC").
				Select("id,title,product_img,price").Find(&product)
			middleMenu[i].ProductItem = product
		}
		c.Data["middleMenuList"] = middleMenu
		models.CacheDb.Set("middleMenu", middleMenu)
	}

	//判断用户是否登陆
	user := models.User{}
	models.Cookie.Get(c.Ctx, "userinfo", &user)
	if len(user.Phone) == 11 {
		str := fmt.Sprintf(`<ul>
			<li class="userinfo">
				<a href="#">%v</a>

				<i class="i"></i>
				<ol>
					<li><a href="/user">个人中心</a></li>

					<li><a href="#">我的收藏</a></li>

					<li><a href="/auth/loginOut">退出登录</a></li>
				</ol>

			</li>
		</ul> `, user.Phone)
		c.Data["userinfo"] = str
	} else {
		str := fmt.Sprintf(`<ul>
			<li><a href="/auth/login" target="_blank">登录</a></li>
			<li>|</li>
			<li><a href="/auth/registerStep1" target="_blank" >注册</a></li>
		</ul>`)
		c.Data["userinfo"] = str
	}
	urlPath, _ := url.Parse(c.Ctx.Request.URL.String())
	c.Data["pathname"] = urlPath.Path
}
