package http

import (
	"encoding/json"
	"io/ioutil"
	"kratos_frame/internal/model"
	err_code "kratos_frame/internal/pkg/ecode"
	"kratos_frame/internal/pkg/oss"
	"kratos_frame/internal/pkg/sha256"
	"net/url"
	"os"
	"regexp"
	"strings"

	"github.com/bilibili/kratos/pkg/log"

	"github.com/gin-gonic/gin/render"

	bm "github.com/bilibili/kratos/pkg/net/http/blademaster"
)

func FileUpload(c *bm.Context) {
	var paths []string
	path, _ := url.QueryUnescape(c.Request.FormValue("path"))
	if err := json.Unmarshal([]byte(path), &paths); err != nil { // path : [\"img\",\"tu\"]
		log.Error("上传文件 解析PATH出现Err：", err, "PATH解码后值", path, "PATH原值", c.Request.FormValue("path"))
		c.JSON(nil, err_code.ArgsErr)
		return
	}
	files := c.Request.MultipartForm.File["files"]
	log.Info("文件上传路径：", paths, "   上传数量", len(files))
	var filePath []string
	for _, file := range files {
		hFile, err := file.Open()
		if err != nil {
			log.Error("上传文件 读文件出现Err：", err)
			c.JSON(nil, err_code.ArgsErr)
			return
		}
		log.Info("上传文件 文件名：", file.Filename)
		fileNames := strings.Split(file.Filename, ".")
		if name, e := oss.PutFile(paths, sha256.Sha(file.Filename)+"."+fileNames[len(fileNames)-1], hFile); e != nil {
			log.Error("上传文件 上传到OSS出现Err：", err)
			c.JSON(nil, err_code.ArgsErr)
			return
		} else {
			filePath = append(filePath, name)
		}
	}

	c.JSON(filePath, err_code.Success)
	return
}

func DelFile(c *bm.Context) {
	var params model.DelFileParams
	if bindArgs(c, &params) != nil {
		return
	}

	b, e := oss.DelFile(params.Path)
	if e != nil {
		c.JSON(nil, err_code.ArgsErr)
		return
	}

	c.JSON(struct {
		IsDel bool
	}{IsDel: b}, err_code.Success)
	return
}

// 富文本图片上传至oss
func Ueditor(c *bm.Context) {
	action := c.Request.FormValue("action")

	switch action {
	//自动读入配置文件，只要初始化UEditor即会发生
	case "config":
		jsonByte, _ := ioutil.ReadFile(os.Getenv("GOPATH") + "/src/kratos_frame/configs/ueditor.json")
		re, _ := regexp.Compile("\\/\\*[\\S\\s]+?\\*\\/")
		jsonByte = re.ReplaceAll(jsonByte, []byte(""))
		var data interface{}
		if json.Unmarshal(jsonByte, &data) != nil {
			c.JSON(nil, err_code.ArgsErr)
			return
		}
		callback := c.Request.FormValue("callback")
		if callback == "" {
			c.Render(200, render.JSON{Data: data})
			return
		}
		c.Render(200, render.JsonpJSON{Data: data})

	default:
		ueditorUp(c)
	}
}

// 富文本上传文件
func ueditorUp(c *bm.Context) {
	file, header, err := c.Request.FormFile("upfile")
	if err != nil || file == nil {
		c.JSON(nil, err_code.ArgsErr)
		return
	}
	fileNames := strings.Split(header.Filename, ".")
	name, e := oss.PutFile([]string{"file", "ueditor"}, sha256.Sha(header.Filename)+"."+fileNames[len(fileNames)-1], file)
	if e != nil {
		c.JSON(nil, err_code.ArgsErr)
		return
	}

	c.JSON(name, err_code.Success)
	return
}
