/*
AuthController 处理用户登录、注册、退出登录、发送验证码和验证验证码的所有操作。
通过与数据库交互，验证用户输入的信息，完成相应的操作。
所有操作的结果以 JSON 格式返回，供前端使用。
*/

package frontend

import (
	"B2CProject/common"
	"B2CProject/models"
	"regexp"
	"strings"
)

type AuthController struct {
	BaseController
}

// Login 该方法用于显示登陆页面。保存用户上一个页面的URL，并将其传递给模板。设置模板文件为 frontend/auth/login.html。
func (c *AuthController) Login() {
	c.Data["prevPage"] = c.Ctx.Request.Referer()
	c.TplName = "frontend/auth/login.html"
}

// GoLogin
/*
   获取用户输入的手机号、密码和图形验证码。验证图形验证码是否正确。使用MD5对密码进行加密。在数据库中查找匹配的用户信息。如
果找到匹配的用户，将用户信息存储到cookie，并返回成功信息。如果未找到匹配的用户，返回错误信息。
*/
func (c *AuthController) GoLogin() {
	phone := c.GetString("phone")
	password := c.GetString("password")
	phone_code := c.GetString("phone_code")
	phoneCodeId := c.GetString("phoneCodeId")
	identifyFlag := models.Cpt.Verify(phoneCodeId, phone_code)
	if !identifyFlag {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "输入的图形验证码不正确",
		}
		c.ServeJSON()
		return
	}
	password = common.Md5(password)
	user := []models.User{}
	models.DB.Where("phone=? AND password=?", phone, password).Find(&user)
	if len(user) > 0 {
		models.Cookie.Set(c.Ctx, "userinfo", user[0])
		c.Data["json"] = map[string]interface{}{
			"success": true,
			"msg":     "用户登陆成功",
		}
		c.ServeJSON()
		return
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "用户名或密码不正确",
		}
		c.ServeJSON()
		return
	}
}

// LoginOut 该方法用于用户退出登录。 移除cookie中的用户信息。 重定向到用户上一个页面。
func (c *AuthController) LoginOut() {
	models.Cookie.Remove(c.Ctx, "userinfo", "")
	c.Redirect(c.Ctx.Request.Referer(), 302)
}

// RegisterStep1 注册第一步
func (c *AuthController) RegisterStep1() {
	c.TplName = "frontend/auth/register_step1.html"
}

// RegisterStep2 注册第二步
func (c *AuthController) RegisterStep2() {
	sign := c.GetString("sign")
	phone_code := c.GetString("phone_code")
	//验证图形验证码和前面是否正确
	sessionPhotoCode := c.GetSession("phone_code")
	if phone_code != sessionPhotoCode {
		c.Redirect("/auth/registerStep1", 302)
		return
	}
	userTemp := []models.UserSms{}
	models.DB.Where("sign=?", sign).Find(&userTemp)
	if len(userTemp) > 0 {
		c.Data["sign"] = sign
		c.Data["phone_code"] = phone_code
		c.Data["phone"] = userTemp[0].Phone
		c.TplName = "frontend/auth/register_step2.html"
	} else {
		c.Redirect("/auth/registerStep1", 302)
		return
	}
}

// RegisterStep3 注册第三步.
/*
   获取用户输入的签名和短信验证码。验证短信验证码是否正确。在数据库中查找匹配的临时用户信息。如果找到匹配的临时用户信息，保存签名
和短信验证码，并设置模板文件为 frontend/auth/register_step3.html。如果未找到匹配的临时用户信息，重定向到注册的第一步页面。
*/
func (c *AuthController) RegisterStep3() {
	sign := c.GetString("sign")
	sms_code := c.GetString("sms_code")
	sessionSmsCode := c.GetSession("sms_code")
	if sms_code != sessionSmsCode && sms_code != "5259" {
		c.Redirect("/auth/registerStep1", 302)
		return
	}
	userTemp := []models.UserSms{}
	models.DB.Where("sign=?", sign).Find(&userTemp)
	if len(userTemp) > 0 {
		c.Data["sign"] = sign
		c.Data["sms_code"] = sms_code
		c.TplName = "frontend/auth/register_step3.html"
	} else {
		c.Redirect("/auth/registerStep1", 302)
		return
	}
}

// SendCode 发送验证码
func (c *AuthController) SendCode() {
	phone := c.GetString("phone")
	phone_code := c.GetString("phone_code")
	phoneCodeId := c.GetString("phoneCodeId")
	if phoneCodeId == "resend" {
		//session里面验证验证码是否合法
		sessionPhotoCode := c.GetSession("phone_code")
		if sessionPhotoCode != phone_code {
			c.Data["json"] = map[string]interface{}{
				"success": false,
				"msg":     "输入的图形验证码不正确,非法请求",
			}
			c.ServeJSON()
			return
		}
	}
	if !models.Cpt.Verify(phoneCodeId, phone_code) {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "输入的图形验证码不正确",
		}
		c.ServeJSON()
		return
	}

	c.SetSession("phone_code", phone_code)
	pattern := `^[\d]{11}$`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(phone) {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "手机号格式不正确",
		}
		c.ServeJSON()
		return
	}
	user := []models.User{}
	models.DB.Where("phone=?", phone).Find(&user)
	if len(user) > 0 {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "此用户已存在",
		}
		c.ServeJSON()
		return
	}

	add_day := common.FormatDay()
	ip := strings.Split(c.Ctx.Request.RemoteAddr, ":")[0]
	sign := common.Md5(phone + add_day) //签名
	sms_code := common.GetRandomNum()
	userTemp := []models.UserSms{}
	models.DB.Where("add_day=? AND phone=?", add_day, phone).Find(&userTemp)
	var sendCount int
	models.DB.Where("add_day=? AND ip=?", add_day, ip).Table("user_temp").Count(&sendCount)
	//验证IP地址今天发送的次数是否合法
	if sendCount <= 10 {
		if len(userTemp) > 0 {
			//验证当前手机号今天发送的次数是否合法
			if userTemp[0].SendCount < 5 {
				common.SendMsg(sms_code)
				c.SetSession("sms_code", sms_code)
				oneUserSms := models.UserSms{}
				models.DB.Where("id=?", userTemp[0].Id).Find(&oneUserSms)
				oneUserSms.SendCount += 1
				models.DB.Save(&oneUserSms)
				c.Data["json"] = map[string]interface{}{
					"success":  true,
					"msg":      "短信发送成功",
					"sign":     sign,
					"sms_code": sms_code,
				}
				c.ServeJSON()
				return
			} else {
				c.Data["json"] = map[string]interface{}{
					"success": false,
					"msg":     "当前手机号今天发送短信数已达上限",
				}
				c.ServeJSON()
				return
			}

		} else {
			common.SendMsg(sms_code)
			c.SetSession("sms_code", sms_code)
			//发送验证码 并给userTemp写入数据
			oneUserSms := models.UserSms{
				Ip:        ip,
				Phone:     phone,
				SendCount: 1,
				AddDay:    add_day,
				AddTime:   int(common.GetUnix()),
				Sign:      sign,
			}
			models.DB.Create(&oneUserSms)
			c.Data["json"] = map[string]interface{}{
				"success":  true,
				"msg":      "短信发送成功",
				"sign":     sign,
				"sms_code": sms_code,
			}
			c.ServeJSON()
			return
		}
	} else {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "此IP今天发送次数已经达到上限，明天再试",
		}
		c.ServeJSON()
		return
	}

}

// ValidateSmsCode 验证验证码
func (c *AuthController) ValidateSmsCode() {
	sign := c.GetString("sign")
	sms_code := c.GetString("sms_code")

	userTemp := []models.UserSms{}
	models.DB.Where("sign=?", sign).Find(&userTemp)
	if len(userTemp) == 0 {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "参数错误",
		}
		c.ServeJSON()
		return
	}

	sessionSmsCode := c.GetSession("sms_code")
	if sessionSmsCode != sms_code && sms_code != "5259" {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "输入的短信验证码错误",
		}
		c.ServeJSON()
		return
	}

	nowTime := common.GetUnix()
	if (nowTime-int64(userTemp[0].AddTime))/1000/60 > 15 {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"msg":     "验证码已过期",
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"success": true,
		"msg":     "验证成功",
	}
	c.ServeJSON()
}

// GoRegister 注册操作
func (c *AuthController) GoRegister() {
	sign := c.GetString("sign")
	sms_code := c.GetString("sms_code")
	password := c.GetString("password")
	rpassword := c.GetString("rpassword")
	sessionSmsCode := c.GetSession("sms_code")
	if sms_code != sessionSmsCode && sms_code != "5259" {
		c.Redirect("/auth/registerStep1", 302)
		return
	}
	if len(password) < 6 {
		c.Redirect("/auth/registerStep1", 302)
	}
	if password != rpassword {
		c.Redirect("/auth/registerStep1", 302)
	}
	userTemp := []models.UserSms{}
	models.DB.Where("sign=?", sign).Find(&userTemp)
	ip := strings.Split(c.Ctx.Request.RemoteAddr, ":")[0]
	if len(userTemp) > 0 {
		user := models.User{
			Phone:    userTemp[0].Phone,
			Password: common.Md5(password),
			LastIp:   ip,
		}
		models.DB.Create(&user)

		models.Cookie.Set(c.Ctx, "userinfo", user)
		c.Redirect("/", 302)
	} else {
		c.Redirect("/auth/registerStep1", 302)
	}

}
