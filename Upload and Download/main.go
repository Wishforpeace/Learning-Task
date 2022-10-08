package main

import (
	"Up_DownLoad/config"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	_ "github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	_ "github.com/qiniu/go-sdk/v7/storage"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func upload(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	file, fileHeader, err := c.Request.FormFile("file")
	fileSize := fileHeader.Size
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "no file"})
	}
	var (
		AccessKey = viper.GetString("qiniu.AccessKey")
		SecretKey = viper.GetString("qiniu.SecretKey")
		Bucket    = viper.GetString("qiniu.Bucket")
		ImgUrl    = viper.GetString("qiniu.QiniuServer")
	)

	putPlicy := storage.PutPolicy{
		Scope: Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPlicy.UploadToken(mac)
	log.Println("token", upToken)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	err = formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Failed to upload",
			"error": err.Error(),
		})
		return
	}
	url := ImgUrl + ret.Key
	c.JSON(http.StatusOK, gin.H{"msg": "Succeed",
		"url": url,
	})
	return
}

func download(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
}

func main() {
	var err error

	err = config.Init("./conf/config.yaml", "")

	if err != nil {
		panic(err)
	}

	r := Router()
	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}

}
func Router() *gin.Engine {
	r := gin.New()
	g1 := r.Group("/api/v1")
	{
		g1.POST("/upload", upload)
		g1.GET("download", download)
	}
	return r
}
