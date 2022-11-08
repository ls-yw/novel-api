package http

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/woodlsy/woodGin/config"
	"github.com/woodlsy/woodGin/log"
	"io/ioutil"
	"strconv"
)

func OssGetObject(bookId int, articleId int) string {
	client, err := oss.New(config.Configs.Aliyun.Oss.Endpoint, config.Configs.Aliyun.Oss.AccessKeyId, config.Configs.Aliyun.Oss.AccessKeySecret)
	if err != nil {
		log.Logger.Error("oss", "连接oss失败", err)
		return ""
	}

	// yourBucketName填写存储空间名称。
	bucket, err := client.Bucket("woodlsy-novel")
	if err != nil {
		log.Logger.Error("oss", "连接bucket失败", err)
		return ""
	}

	objectPath := "book/" + strconv.Itoa(bookId) + "/" + strconv.Itoa(articleId) + ".txt"
	// 下载文件到流。
	body, err := bucket.GetObject(objectPath)
	if err != nil {
		log.Logger.Error("oss", "找不到文件", objectPath, err)
		return ""
	}
	// 数据读取完成后，获取的流必须关闭，否则会造成连接泄漏，导致请求无连接可用，程序无法正常工作。
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		log.Logger.Error("oss", "读取文件内容失败", objectPath, err)
		return ""
	}
	return string(data)
}
