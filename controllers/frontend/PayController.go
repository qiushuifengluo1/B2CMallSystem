/*
   上述代码展示了如何使用 Beego 框架实现支付宝和微信支付功能，包括支付请求的发起、异步通知的处理以及支付结果的确认。每个支付方法
都包含了对订单信息的处理、支付接口的调用以及支付状态的更新。
*/

package frontend

import (
	"B2CProject/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/objcoding/wxpay"
	"github.com/skip2/go-qrcode"
	"github.com/smartwalle/alipay/v3"
	"strconv"
	"strings"
	"time"
)

type PayController struct {
	BaseController
}

func (c *PayController) Alipay() {
	// 获取支付订单 ID
	AliId, err1 := c.GetInt("id")
	if err1 != nil {
		c.Redirect(c.Ctx.Request.Referer(), 302)
	}

	// 获取订单项
	orderitem := []models.OrderItem{}
	models.DB.Where("order_id=?", AliId).Find(&orderitem)

	//设置私钥和初始化支付宝客户端
	var privateKey = "`-----BEGIN RSA PRIVATE KEY-----\n" +
		"MIIEpAIBAAKCAQEAtJQJgqkRvFVPMZ5N6XJ2r+LZ5UGmUx9UjSknHIlzlsFd1xTx\n" +
		"vKz4fu1ovqR3bbFRPq1BqAaD8qFx0Pxi9VtHVR1h3uF5BcxL8tD3U1F0kzHoWxl2\n" +
		"NU7Y4KhQle0Zgr1a5cM3/RMdD9qU79T6AcwXNkiT7YudCqdi9V6TJK1yFGloHLjL\n" +
		"sPyGROEX9KUL9Rbq6mZZFH62F0fz5aQO4ynOGC+QotxBt03cPz+2mSztDq2ZXRCF\n" +
		"G2t+0vUI1UT6E7uB54ZZ1l3mBLDVOef2V63D6UuQK5EWB4IWBBEFYZy72ZlckuFZ\n" +
		"OeB7nI9uTRzAxy7Ndz5/7yGCPEYvaf2U0uNN5QIDAQABAoIBAQCjXJbQ0T6cvZo6\n" +
		"MSpUsR2q3mnCqZ5N9YFsyYb4J6cOge3Y79FRs/8HLY8VoYHZMSseWx7K25v1++NG\n" +
		"8LFThle4NhPoibVmIesgOGkVZnrtbzOzGhdZ+GcLJw3D1WmBQnFwRTy4GeE6FCJD\n" +
		"MO5ylKrrU43K0VzZ2tVoCZh3bFNGzWPGhhoJczJZ8CydNGOnnl5aiD9bH0F4rznM\n" +
		"gfPZdT64Y68K1Z9O1Xb63UMso97V9tFYwosT0P2TG9Mz/JON6YbbJW5BzX+esH6x\n" +
		"50wYm6FZr/OxgWxCMq5x+K6RhJMgLBtMhq2BZjP62LmsJgJXY2WtTkRpDh8LwzFJ\n" +
		"kB9+NPKBAoGBAPL4/1X8HdTvCp9MFQIRvFuL14FJj3szDN84QdpHtq9I/XYiODv6\n" +
		"UMaNc6dczt+Vx8TQL7cODUmEAZDLpmkiwhZLBrG3dxxqLB7oL9u0D/KrAyn8Ow2z\n" +
		"DrXaTBHMGEdRFg9LPjR4kQ3Z/ziICxrgEB/NznABv0SYK7FYZW7A7lyVAoGBAMql\n" +
		"QoZYieP2/UUsMZd0p1jRDRs1pQ5KUmh+EGNNrcJAXcIUBHJfIQhcIv3uTRn5tftQ\n" +
		"u0OmJeUPspX8Ic0SyJzYrWxz0lD/Fx0ZrEqqlxciUoTHO0Bu+F1HmvXZmYr/ZMMr\n" +
		"2Wtw/1fXjFVdZ9eDdBGlT0wIBfc+d0VEAHfEDMblAoGAH5buxhOPXB32Llpe8Trt\n" +
		"dP13iQLQ0aRfZkW2exTEp3URH/OIVmXhG1O9zQHH6BeHl6ufh0rTBL1V4dy+TOB4\n" +
		"O8a2rWXIWgG1vlETRWtD5m7jN2DRIsfuSn9a3BmkElVp5LBQoJzGLyr41GCtGrOV\n" +
		"jftdtNwO/l7UwJ0lJ1tEewUCgYEAiNY8dnklmDyDce6VAtVs5VSwUVLt3+xZ8Q7f\n" +
		"gK8m7N6wSbszRY/LX4qVR4bnlV9qfKlRuM0QpJYJddCNV+Q3hN4+leKSPWhfb5Zh\n" +
		"AYrZwGPtVe+bMGA7quO/0RQ1aRzpTDH5Yt2S0NExRWh8TvqEd7cZmCz26m2aHEgz\n" +
		"z7aYw9kCgYB5QFCQNU08CIvq7R6vOi+03qlKJvKcZPwtk3tO2+fy+yU1CXnD1qlk\n" +
		"AwBLDnpMkiH1U8byuSDZl6/FrUN+ZzZDh7GhLsRuGagC7R2fZn+76uF88F9c1VEE\n" +
		"hT0tFJ0kMgdVikMy7DkwUXH1bqrgSl16tMgEcHutF5zjKDtH9Rly0A==\n" +
		"-----END RSA PRIVATE KEY-----`" // 模拟生成的私钥
	var client, err = alipay.New("2021001186696588", privateKey, true)
	err = client.LoadAppPublicCertFromFile("certfile/appCertPublicKey_2021001186696588.certfile")
	if err != nil {
		return
	} // 加载应用公钥证书
	err = client.LoadAliPayRootCertFromFile("certfile/alipayRootCert.certfile")
	if err != nil {
		return
	} // 加载支付宝根证书
	err = client.LoadAliPayPublicCertFromFile("certfile/alipayCertPublicKey_RSA2.certfile")
	if err != nil {
		return
	} // 加载支付宝公钥证书

	// 将 key 的验证调整到初始化阶段
	if err != nil {
		fmt.Println(err)
		return
	}

	//计算总价格
	var TotalAmount float64
	for i := 0; i < len(orderitem); i++ {
		TotalAmount = TotalAmount + orderitem[i].ProductPrice
	}
	var p = alipay.TradePagePay{}
	p.NotifyURL = "http://localhost/alipay/notify" //模拟生成
	p.ReturnURL = "http://localhost/alipay/return" //模拟生成
	p.TotalAmount = "100.0"
	p.Subject = "订单order——" + time.Now().Format("200601021504")
	p.OutTradeNo = "WF" + time.Now().Format("200601021504") + "_" + strconv.Itoa(AliId)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	var url, err4 = client.TradePagePay(p)
	if err4 != nil {
		fmt.Println(err4)
	}
	var payURL = url.String()
	c.Redirect(payURL, 302)
}

func (c *PayController) AlipayNotify() {
	var privateKey = "`-----BEGIN RSA PRIVATE KEY-----\n" +
		"MIIEpAIBAAKCAQEAtJQJgqkRvFVPMZ5N6XJ2r+LZ5UGmUx9UjSknHIlzlsFd1xTx\n" +
		"vKz4fu1ovqR3bbFRPq1BqAaD8qFx0Pxi9VtHVR1h3uF5BcxL8tD3U1F0kzHoWxl2\n" +
		"NU7Y4KhQle0Zgr1a5cM3/RMdD9qU79T6AcwXNkiT7YudCqdi9V6TJK1yFGloHLjL\n" +
		"sPyGROEX9KUL9Rbq6mZZFH62F0fz5aQO4ynOGC+QotxBt03cPz+2mSztDq2ZXRCF\n" +
		"G2t+0vUI1UT6E7uB54ZZ1l3mBLDVOef2V63D6UuQK5EWB4IWBBEFYZy72ZlckuFZ\n" +
		"OeB7nI9uTRzAxy7Ndz5/7yGCPEYvaf2U0uNN5QIDAQABAoIBAQCjXJbQ0T6cvZo6\n" +
		"MSpUsR2q3mnCqZ5N9YFsyYb4J6cOge3Y79FRs/8HLY8VoYHZMSseWx7K25v1++NG\n" +
		"8LFThle4NhPoibVmIesgOGkVZnrtbzOzGhdZ+GcLJw3D1WmBQnFwRTy4GeE6FCJD\n" +
		"MO5ylKrrU43K0VzZ2tVoCZh3bFNGzWPGhhoJczJZ8CydNGOnnl5aiD9bH0F4rznM\n" +
		"gfPZdT64Y68K1Z9O1Xb63UMso97V9tFYwosT0P2TG9Mz/JON6YbbJW5BzX+esH6x\n" +
		"50wYm6FZr/OxgWxCMq5x+K6RhJMgLBtMhq2BZjP62LmsJgJXY2WtTkRpDh8LwzFJ\n" +
		"kB9+NPKBAoGBAPL4/1X8HdTvCp9MFQIRvFuL14FJj3szDN84QdpHtq9I/XYiODv6\n" +
		"UMaNc6dczt+Vx8TQL7cODUmEAZDLpmkiwhZLBrG3dxxqLB7oL9u0D/KrAyn8Ow2z\n" +
		"DrXaTBHMGEdRFg9LPjR4kQ3Z/ziICxrgEB/NznABv0SYK7FYZW7A7lyVAoGBAMql\n" +
		"QoZYieP2/UUsMZd0p1jRDRs1pQ5KUmh+EGNNrcJAXcIUBHJfIQhcIv3uTRn5tftQ\n" +
		"u0OmJeUPspX8Ic0SyJzYrWxz0lD/Fx0ZrEqqlxciUoTHO0Bu+F1HmvXZmYr/ZMMr\n" +
		"2Wtw/1fXjFVdZ9eDdBGlT0wIBfc+d0VEAHfEDMblAoGAH5buxhOPXB32Llpe8Trt\n" +
		"dP13iQLQ0aRfZkW2exTEp3URH/OIVmXhG1O9zQHH6BeHl6ufh0rTBL1V4dy+TOB4\n" +
		"O8a2rWXIWgG1vlETRWtD5m7jN2DRIsfuSn9a3BmkElVp5LBQoJzGLyr41GCtGrOV\n" +
		"jftdtNwO/l7UwJ0lJ1tEewUCgYEAiNY8dnklmDyDce6VAtVs5VSwUVLt3+xZ8Q7f\n" +
		"gK8m7N6wSbszRY/LX4qVR4bnlV9qfKlRuM0QpJYJddCNV+Q3hN4+leKSPWhfb5Zh\n" +
		"AYrZwGPtVe+bMGA7quO/0RQ1aRzpTDH5Yt2S0NExRWh8TvqEd7cZmCz26m2aHEgz\n" +
		"z7aYw9kCgYB5QFCQNU08CIvq7R6vOi+03qlKJvKcZPwtk3tO2+fy+yU1CXnD1qlk\n" +
		"AwBLDnpMkiH1U8byuSDZl6/FrUN+ZzZDh7GhLsRuGagC7R2fZn+76uF88F9c1VEE\n" +
		"hT0tFJ0kMgdVikMy7DkwUXH1bqrgSl16tMgEcHutF5zjKDtH9Rly0A==\n" +
		"-----END RSA PRIVATE KEY-----`" // 模拟生成的私钥
	var client, err = alipay.New("2021001186696588", privateKey, true)

	err = client.LoadAppPublicCertFromFile("certfile/appCertPublicKey_2021001186696588.certfile")
	if err != nil {
		return
	} // 加载应用公钥证书
	err = client.LoadAliPayRootCertFromFile("certfile/alipayRootCert.certfile")
	if err != nil {
		return
	} // 加载支付宝根证书
	err = client.LoadAliPayPublicCertFromFile("certfile/alipayCertPublicKey_RSA2.certfile")
	if err != nil {
		return
	} // 加载支付宝公钥证书

	if err != nil {
		fmt.Println(err)
		return
	}

	req := c.Ctx.Request
	req.ParseForm()
	err = client.VerifySign(req.Form)
	if err != nil {
		c.Redirect(c.Ctx.Request.Referer(), 302)
	}
	rep := c.Ctx.ResponseWriter
	var noti, _ = client.GetTradeNotification(req)
	if noti != nil {
		fmt.Println("交易状态为:", noti.TradeStatus)
		if string(noti.TradeStatus) == "TRADE_SUCCESS" {
			order := models.Order{}
			temp := strings.Split(noti.OutTradeNo, "_")[1]
			id, _ := strconv.Atoi(temp)
			models.DB.Where("id=?", id).Find(&order)
			order.PayStatus = 1
			order.OrderStatus = 1
			models.DB.Save(&order)
		}
	}
	alipay.AckNotification(rep) // 确认收到通知消息
}
func (c *PayController) AlipayReturn() {
	c.Redirect("/user/order", 302)
}

func (c *PayController) WxPay() {
	WxId, err := c.GetInt("id")
	if err != nil {
		c.Redirect(c.Ctx.Request.Referer(), 302)
	}
	orderitem := []models.OrderItem{}
	models.DB.Where("order_id=?", WxId).Find(&orderitem)
	//1、配置基本信息
	account := wxpay.NewAccount(
		"wx1234567890abcdef",               // 随机生成的测试appid
		"1234567890",                       // 随机生成的测试商户号
		"abcdef1234567890abcdef1234567890", // 随机生成的测试appkey
		false,
	)
	client := wxpay.NewClient(account)
	var price int64
	for i := 0; i < len(orderitem); i++ {
		price = 1
	}
	//2、获取ip地址,订单号等信息
	ip := strings.Split(c.Ctx.Request.RemoteAddr, ":")[0]
	template := "202001021504"
	tradeNo := time.Now().Format(template)
	//3、调用统一下单
	params := make(wxpay.Params)
	params.SetString("body", "order——"+time.Now().Format(template)).
		SetString("out_trade_no", tradeNo+"_"+strconv.Itoa(WxId)).
		SetInt64("total_fee", price).
		SetString("spbill_create_ip", ip).
		SetString("notify_url", "http://localhost/wxpay/notify"). //配置的回调地址
		SetString("trade_type", "APP").                           //APP端支付
		SetString("trade_type", "NATIVE")                         //网站支付需要改为NATIVE

	p, err1 := client.UnifiedOrder(params)
	logs.Info(p)
	if err1 != nil {
		logs.Error(err1)
		c.Redirect(c.Ctx.Request.Referer(), 302)
	}
	//4、获取code_url生成支付二维码
	var pngObj []byte
	logs.Info(p)
	pngObj, _ = qrcode.Encode(p["code_url"], qrcode.Medium, 256)
	c.Ctx.WriteString(string(pngObj))
}

func (c *PayController) WxPayNotify() {
	//1、获取表单传过来的xml数据，在配置文件里设置 copyrequestbody = true
	xmlStr := string(c.Ctx.Input.RequestBody)
	postParams := wxpay.XmlToMap(xmlStr)
	beego.Info(postParams)

	//2、校验签名
	account := wxpay.NewAccount(
		"wx1234567890abcdef",               // 随机生成的测试appid
		"1234567890",                       // 随机生成的测试商户号
		"abcdef1234567890abcdef1234567890", // 随机生成的测试appkey
		false,
	)
	client := wxpay.NewClient(account)
	isValidate := client.ValidSign(postParams)
	// xml解析
	params := wxpay.XmlToMap(xmlStr)
	beego.Info(params)
	if isValidate == true {
		if params["return_code"] == "SUCCESS" {
			idStr := strings.Split(params["out_trade_no"], "_")[1]
			id, _ := strconv.Atoi(idStr)
			order := models.Order{}
			models.DB.Where("id=?", id).Find(&order)
			order.PayStatus = 1
			order.PayType = 1
			order.OrderStatus = 1
			models.DB.Save(&order)
		}
	} else {
		c.Redirect(c.Ctx.Request.Referer(), 302)
	}
}
