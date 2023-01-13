package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	_ "github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	_ "github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/viper"
	"mime/multipart"
	_ "mime/multipart"
	"net/http"
	"strings"
	"time"
)

type Qiniu struct {
	AccessKey string
	SecretKey string
	Bucket    string
	Domain    string
}

var Q Qiniu

func InitConfig(prefix string) error {
	viper.AddConfigPath("./conf")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml") // 默认配置文件类型为yaml
	viper.AutomaticEnv()        // 读取默认的环境变量
	viper.SetEnvPrefix(prefix)  // 读取环境变量的前缀为APISERVER
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}
func Load() {
	Q = Qiniu{
		AccessKey: viper.GetString("oss.access_key"),
		SecretKey: viper.GetString("oss.secret_key"),
		Bucket:    viper.GetString("oss.bucket_name"),
		Domain:    viper.GetString("oss.domain_name"),
	}
}

func UploadQiniu(file *multipart.FileHeader) (int, string) {
	src, err := file.Open()
	if err != nil {
		return 10011, err.Error()
	}

	defer src.Close()

	putPolicy := storage.PutPolicy{
		Scope: Q.Bucket,
	}

	mac := qbox.NewMac(Q.AccessKey, Q.SecretKey)

	// 获取上传凭证
	upToken := putPolicy.UploadToken(mac)

	// 配置参数
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}        // 上传返回后的结果
	putExtra := storage.PutExtra{} // 额外参数

	// 自定义文件名及后缀
	key := "(" + time.Now().String() + ")" + file.Filename

	if err := formUploader.Put(context.Background(), &ret,
		upToken, key, src, file.Size, &putExtra); err != nil {
		return 501, err.Error()
	}

	return 0, Q.Domain + "/" + ret.Key
}

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 10010,
			"msg":  err.Error(),
		})
		return
	}

	code, url := UploadQiniu(file)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"url":  url,
	})

}

func Route() *gin.Engine {
	r := gin.New()
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router.")
	})
	g1 := r.Group("/api/v1")
	{
		g1.POST("/image", UploadImage)
	}

	return r
}
func main() {
	InitConfig("qiniu")
	Load()
	gin.SetMode(gin.TestMode)
	r := Route()
	r.Run(":8080")
}
