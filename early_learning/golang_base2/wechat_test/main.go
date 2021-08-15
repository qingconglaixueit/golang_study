package main

import (
	"crypto/sha1"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

type myData struct {
	Token   string
	TimeSta string
	RandNum string
}

type SConfig struct {
	XMLName      xml.Name   `xml:"config"`     // 指定最外层的标签为config
	SmtpServer   string     `xml:"smtpServer"` // 读取smtpServer配置项，并将结果保存到SmtpServer变量中
	SmtpPort     int        `xml:"smtpPort"`
	Sender       string     `xml:"sender"`
	SenderPasswd string     `xml:"senderPasswd"`
	Receivers    SReceivers `xml:"receivers"` // 读取receivers标签下的内容，以结构方式获取
}

type SReceivers struct {
	Age    int      `xml:"age"`
	Flag   string   `xml:"flag,attr"` // 读取flag属性
	User   []string `xml:"user"`      // 读取user数组
	Script string   `xml:"script"`    // 读取 <![CDATA[ xxx ]]> 数据
}

func GenerateRandNum(len int) string {

	var res string
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	for i := 0; i < len; i++ {
		tmp := rand.Intn(len)
		res += string(str[tmp])
	}

	return res
}

//生成token
//微信的规则
//token + 时间戳 + 随机数  通过sha1加密 生成签名
func GenerateSignatrue(data *myData) (string, error) {
	if data == nil {
		logger.Info("data is nil")
		return "", errors.New("data is nil")
	}

	//获取当前时间戳
	data.TimeSta = strconv.FormatInt(time.Now().Unix(), 10)

	//获取随机数
	data.RandNum = GenerateRandNum(5)

	//生成签名
	var str string
	//str = fmt.Sprint("%s%s%s", data.Token, data.TimeSta, data.RandNum)
	//传参可以任意顺序
	strs := sort.StringSlice{data.Token, data.TimeSta, data.RandNum}
	sort.Strings(strs)

	logger.Info("111 strs == ", strs)

	for _, v := range strs {
		str += v
	}

	h := sha1.New()
	h.Write([]byte(str))

	signature := fmt.Sprintf("%x", h.Sum(nil))

	logger.Info("111 generate signatur is ", signature)

	return signature, nil
}

//检验签名
func VerifySig(data *myData, sig string) error {
	if data == nil {
		logger.Info("data is nil")
		return errors.New("data is nil")
	}

	//生成签名
	var str string
	//str = fmt.Sprint("%s%s%s", data.Token, data.TimeSta, data.RandNum)

	//传参可以任意顺序
	strs := sort.StringSlice{data.Token, data.RandNum, data.TimeSta}
	sort.Strings(strs)

	logger.Info("222 strs == ", strs)

	for _, v := range strs {
		str += v
	}

	h := sha1.New()
	h.Write([]byte(str))

	signature := fmt.Sprintf("%x", h.Sum(nil))

	logger.Info("2222 generate signatur is ", signature)

	if sig == signature {
		return nil
	}
	return errors.New("sig != signature")
}

func checkSig() {
	//客户端生成签名
	token := "qqqqbbbqwe"

	mydata := &myData{}
	mydata.Token = token
	signature, _ := GenerateSignatrue(mydata)

	//服务器检验签名
	err := VerifySig(mydata, signature)
	if err != nil {
		logger.Error("VerifySig failed")
		return
	}
	logger.Info("VerifySig success")
}

//xml 解析
func main() {

	//得到xml的byte数组

	fp, err := os.Open("./test.xml")
	if err != nil {
		logger.Info("Open error  ", err)
		return
	}

	data, err := ioutil.ReadAll(fp)
	if err != nil {
		logger.Info("ReadAll error  ", err)
		return
	}
	//解析xml
	res := &SConfig{}
	err = xml.Unmarshal(data, res)
	if err != nil {
		logger.Info("Unmarshal error  ", err)
		return
	}

	logger.Info("res == ", res)

}
