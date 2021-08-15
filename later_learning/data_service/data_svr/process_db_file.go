package data_svr

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// 获取图片数据
func GetPic(url string) string {
	// 获取具体的图片名字
	fileName := GetFilename(url)
	// 下载图片
	return DownloadPic(url, fileName)
}

// 获取到 文件的 名字
func GetFilename(url string) (filename string) {
	// 找到最后一个 = 的索引
	lastIndex := strings.LastIndex(url, "/")
	// 获取 / 后的字符串 ，这就是源文件名
	filename = url[lastIndex+1:]

	// 把时间戳 加 在原来名字前，拼一个新的名字
	prefix := fmt.Sprintf("%d", time.Now().Unix())
	filename = prefix + "_" + filename

	return filename
}

func DownloadPic(url string, filename string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("http.Get error : ", err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ioutil.ReadAll error : ", err)
	}

	// 文件存放的路径
	filename = "./src/" + filename

	// 写文件 并设置文件权限
	err = ioutil.WriteFile(filename, bytes, 0666)
	if err != nil {
		log.Fatal("wirte failed !!", err)
	} else {
		log.Println("ioutil.WriteFile successfully , filename = ", filename)
	}

	return filename
}
