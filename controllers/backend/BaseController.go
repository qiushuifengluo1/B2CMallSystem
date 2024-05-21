//BaseController 提供了消息处理、页面跳转和图片上传等基础功能，便于后台管理系统的开发。

package backend

import (
	"B2CProject/common"
	"B2CProject/models"
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/astaxie/beego"
	"os"
	"path"
	"strconv"
	"strings"
)

type BaseController struct {
	beego.Controller
}

// Success 方法用于显示成功消息，并重定向用户到指定页面。
func (c *BaseController) Success(message string, redirect string) {
	c.Data["Message"] = message
	if strings.Contains(redirect, "http") {
		c.Data["Redirect"] = redirect
	} else {
		c.Data["Redirect"] = "/" + beego.AppConfig.String("adminPath") + redirect
	}
	c.TplName = "backend/public/success.html"
}

// Error 方法用于显示错误消息，并重定向用户到指定页面。
func (c *BaseController) Error(message string, redirect string) {
	c.Data["Message"] = message
	if strings.Contains(redirect, "http") {
		c.Data["Redirect"] = redirect
	} else {
		c.Data["Redirect"] = "/" + beego.AppConfig.String("adminPath") + redirect
	}
	c.TplName = "backend/public/error.html"
}

// Goto 方法用于重定向用户到指定页面。
func (c *BaseController) Goto(redirect string) {
	c.Redirect("/"+beego.AppConfig.String("adminPath")+redirect, 302)
}

// UploadImg 方法根据配置选择本地上传或阿里云 OSS 上传图片。
func (c *BaseController) UploadImg(picName string) (string, error) {
	ossStatus, _ := beego.AppConfig.Bool("ossStatus")
	if ossStatus == true {
		return c.OssUploadImg(picName)
	}
	return c.LocalUploadImg(picName)
}

// LocalUploadImg 方法处理本地图片上传，包括检查文件类型、创建保存目录和保存文件。
func (c *BaseController) LocalUploadImg(picName string) (string, error) {
	f, h, err := c.GetFile(picName)
	if err != nil {
		return "", err
	}
	//2、关闭文件流
	defer f.Close()
	//3、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(h.Filename)

	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("图片后缀名不合法")
	}
	//4、创建图片保存目录  static/upload/20200623
	day := common.FormatDay()
	dir := "static/upload/" + day

	if err := os.MkdirAll(dir, 0666); err != nil {
		return "", err
	}
	//5、生成文件名称   144325235235.png
	fileUnixName := strconv.FormatInt(common.GetUnixNano(), 10)
	//static/upload/20200623/144325235235.png
	saveDir := path.Join(dir, fileUnixName+extName)
	//6、保存图片

	c.SaveToFile(picName, saveDir)
	return saveDir, nil

}

// OssUploadImg 方法处理将图片上传到阿里云 OSS，包括创建 OSS 客户端、获取存储桶和上传文件。
func (c *BaseController) OssUploadImg(picName string) (string, error) {
	setting := models.Setting{}
	models.DB.First(&setting)
	f, h, err := c.GetFile(picName)
	if err != nil {
		return "", err
	}
	//2、关闭文件流
	defer f.Close()
	//3、获取后缀名 判断类型是否正确  .jpg .png .gif .jpeg
	extName := path.Ext(h.Filename)

	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("图片后缀名不合法")
	}
	//把文件流上传值OSS

	//4.1 创建OSS实例
	client, err := oss.New(setting.EndPoint, setting.Appid, setting.AppSecret)
	if err != nil {
		return "", err
	}

	// 4.2获取存储空间。
	bucket, err := client.Bucket(setting.BucketName)
	if err != nil {
		return "", err
	}
	//4.3创建图片保存目录  static/upload/20200623
	day := common.FormatDay()
	dir := "static/upload/" + day
	fileUnixName := strconv.FormatInt(common.GetUnixNano(), 10)
	//static/upload/20200623/144325235235.png
	saveDir := path.Join(dir, fileUnixName+extName)
	// 4.4上传文件流。
	err = bucket.PutObject(saveDir, f)
	if err != nil {
		return "", err
	}
	return saveDir, nil
}

// GetSetting 方法从数据库中获取设定项并返回。
func (c *BaseController) GetSetting() models.Setting {
	setting := models.Setting{Id: 1}
	models.DB.First(&setting)
	return setting
}
