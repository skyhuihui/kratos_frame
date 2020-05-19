package oss

import (
	"io"

	"github.com/bilibili/kratos/pkg/log"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	OssCli *oss.Client
)

const (
	BucketUrl       = ""
	AccessKeyID     = ""
	AccessKeySecret = ""
	Bucket          = ""
)

func New() (err error) {
	OssCli, err = oss.New(BucketUrl, AccessKeyID, AccessKeySecret)

	return err
}

// path 为上传路径， path[0]是一级目录， path[1] 是二级目录 以此类推
func PutFile(path []string, fileName string, fileContent io.Reader) (string, error) {
	objectName := ""
	for _, v := range path {
		objectName += v + "/"
	}
	objectName = objectName + fileName
	// 获取存储空间。
	bucket, err := OssCli.Bucket(Bucket)
	if err != nil {
		log.Error("OSS 获取储存空间出错", err)
		return "", err
	}

	// 上传文件。
	err = bucket.PutObject(objectName, fileContent)
	if err != nil {
		log.Error("OSS 上传文件失败", err)
		return "", err
	}

	imageUrl := "https://.oss-cn-beijing.aliyuncs.com/" + objectName

	return imageUrl, nil
}

func DelFile(fileUrl string) (bool, error) {
	// 获取存储空间。
	bucket, err := OssCli.Bucket(Bucket)
	if err != nil {
		log.Error("OSS 获取储存空间出错", err)
		return false, err
	}

	// 删除文件
	err = bucket.DeleteObject(fileUrl)
	if err != nil {
		log.Error("OSS 删除文件失败", err)
		return false, err
	}
	return true, nil
}
