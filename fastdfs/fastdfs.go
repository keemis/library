package fastdfs

import (
	"errors"
	"fmt"

	"github.com/astaxie/beego"
	fdfs "github.com/keemis/go-fastdfs-client"
)

// Client 全局
var Client *fdfs.Client

// InitFastDfs 初始化
func Init() {
	maxConn := beego.AppConfig.DefaultInt("MaxConn", 100)
	trackerAddrs := beego.AppConfig.DefaultStrings("TrackerAddrs", []string{"127.0.0.1:22122"})
	if maxConn <= 0 {
		panic("fastdfs MaxConn must > 0")
	}
	if len(trackerAddrs) == 0 {
		panic("fastdfs TrackerAddrs len must > 0")
	}
	conf := &fdfs.Config{
		TrackerAddrs: trackerAddrs,
		MaxConn:      maxConn,
	}
	c, err := fdfs.New(conf)
	if err != nil {
		panic(fmt.Sprintf("fastdfs new client error: %v", err))
	}
	Client = c
}

// UploadByBuffer 上传
func UploadByBuffer(buffer []byte) (string, error) {
	if Client == nil {
		return "", errors.New("client is nil, must be init first")
	}
	fileID, err := Client.UploadByBuffer(buffer, "")
	if err != nil {
		return "", err
	}
	return fileID, nil
}

// DownloadToBuffer 下载
func DownloadToBuffer(fileID string) ([]byte, error) {
	if Client == nil {
		return nil, errors.New("client is nil, must be init first")
	}
	byt, err := Client.DownloadToBuffer(fileID, 0, 0)
	if err != nil {
		return nil, err
	}
	return byt, nil
}

// DeleteFile 删除
func DeleteFile(fileID string) error {
	if Client == nil {
		return errors.New("client is nil, must be init first")
	}
	return Client.DeleteFile(fileID)
}
