package routers

import (
	"B2CProject/common"
	"B2CProject/controllers/frontend"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &frontend.IndexController{})
	beego.Router("/category_:id([0-9]+).html", &frontend.ProductController{}, "get:CategoryList")
	beego.Router("/item_:id([0-9]+).html", &frontend.ProductController{}, "get:ProductItem")
	beego.Router("/seckill/item_:id([0-9]+).html", &frontend.SecKillController{}, "get:ProductItem")
	beego.Router("/product/getImgList", &frontend.ProductController{}, "get:GetImgList")
	beego.Router("/cart", &frontend.CartController{})
	beego.Router("/cart/addCart", &frontend.CartController{}, "get:AddCart")
	beego.Router("/cart/incCart", &frontend.CartController{}, "get:IncCart")
	beego.Router("/cart/decCart", &frontend.CartController{}, "get:DecCart")
	beego.Router("/cart/delCart", &frontend.CartController{}, "get:DelCart")
	beego.Router("/cart/changeOneCart", &frontend.CartController{}, "get:ChangeOneCart")
	beego.Router("/cart/changeAllCart", &frontend.CartController{}, "get:ChangeAllCart")
	beego.Router("/product/collect", &frontend.ProductController{}, "get:Collect")

	beego.Router("/auth/sendCode", &frontend.AuthController{}, "get:SendCode")
	beego.Router("/auth/doRegister", &frontend.AuthController{}, "post:GoRegister")
	beego.Router("/auth/validateSmsCode", &frontend.AuthController{}, "get:ValidateSmsCode")
	beego.Router("/auth/login", &frontend.AuthController{}, "get:Login")
	beego.Router("/auth/registerStep1", &frontend.AuthController{}, "get:RegisterStep1")
	beego.Router("/auth/registerStep2", &frontend.AuthController{}, "get:RegisterStep2")
	beego.Router("/auth/registerStep3", &frontend.AuthController{}, "get:RegisterStep3")
	beego.Router("/auth/login", &frontend.AuthController{}, "get:Login")
	beego.Router("/auth/goLogin", &frontend.AuthController{}, "post:GoLogin")
	beego.Router("/auth/loginOut", &frontend.AuthController{}, "get:LoginOut")

	//配置中间件判断权限
	beego.InsertFilter("/buy/*", beego.BeforeRouter, common.FrontendAuth)
	beego.Router("/buy/checkout", &frontend.CheckoutController{}, "get:Checkout")
	beego.Router("/buy/doOrder", &frontend.CheckoutController{}, "post:GoOrder")
	beego.Router("/buy/confirm", &frontend.CheckoutController{}, "get:Confirm")
	beego.Router("/buy/orderPayStatus", &frontend.CheckoutController{}, "get:OrderPayStatus")
	//配置中间件判断权限
	beego.InsertFilter("/address/*", beego.BeforeRouter, common.FrontendAuth)
	beego.Router("/address/addAddress", &frontend.AddressController{}, "post:AddAddress")
	beego.Router("/address/getOneAddressList", &frontend.AddressController{}, "get:GetOneAddressList")
	beego.Router("/address/goEditAddressList", &frontend.AddressController{}, "post:GoEditAddressList")
	beego.Router("/address/changeDefaultAddress", &frontend.AddressController{}, "get:ChangeDefaultAddress")
	//支付宝支付
	beego.Router("/alipay", &frontend.PayController{}, "get:Alipay")
	beego.Router("/alipayNotify", &frontend.PayController{}, "post:AlipayNotify")
	beego.Router("/alipayReturn", &frontend.PayController{}, "get:AlipayReturn")

	//微信支付
	beego.Router("/wxpay", &frontend.PayController{}, "get:WxPay")
	beego.Router("/wxpay/notify", &frontend.PayController{}, "post:WxPayNotify")

	//配置中间件判断权限
	beego.InsertFilter("/user/*", beego.BeforeRouter, common.FrontendAuth)
	beego.Router("/user", &frontend.UserController{})
	beego.Router("/user/order", &frontend.UserController{}, "get:OrderList")
	beego.Router("/user/orderinfo", &frontend.UserController{}, "get:OrderInfo")

	//搜索
	// beego.Router("/search", &frontend.SearchController{})
	// beego.Router("/search/getOne", &frontend.SearchController{}, "get:GetOne")
	// beego.Router("/search/update", &frontend.SearchController{}, "get:Update")
	// beego.Router("/search/delete", &frontend.SearchController{}, "get:Delete")
	// beego.Router("/search/query", &frontend.SearchController{}, "get:Query")
	// beego.Router("/search/filterQuery", &frontend.SearchController{}, "get:FilterQuery")
	beego.Router("/search/productList", &frontend.SearchController{}, "get:ProductList")
}
