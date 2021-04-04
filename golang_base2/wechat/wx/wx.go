package wx

import (
	"crypto/sha1"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/clbanning/mxj"
)

type weixinQuery struct {
	Signature    string `json:"signature"`
	Timestamp    string `json:"timestamp"`
	Nonce        string `json:"nonce"`
	EncryptType  string `json:"encrypt_type"`
	MsgSignature string `json:"msg_signature"`
	Echostr      string `json:"echostr"`
}

type WeixinClient struct {
	Token          string
	Query          weixinQuery // 请求的一些参数
	Message        map[string]interface{}
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	Methods        map[string]func() bool
}

/// 请求数据Request， 返回数据ResponseWriter， token是自己的
func NewClient(r *http.Request, w http.ResponseWriter, token string) (*WeixinClient, error) {

	weixinClient := new(WeixinClient)

	weixinClient.Token = token // 获取本地的token
	weixinClient.Request = r
	weixinClient.ResponseWriter = w

	weixinClient.initWeixinQuery()
	logger.Info("Signature:", weixinClient.Query.Signature)
	if weixinClient.Query.Signature != weixinClient.hashcode() { // 签名认证
		return nil, errors.New("Invalid Signature.")
	}

	return weixinClient, nil
}

func (this *WeixinClient) initWeixinQuery() {

	var q weixinQuery
	logger.Info("URL:", this.Request.URL.Path, ", RawQuery:", this.Request.URL.RawPath)
	q.Nonce = this.Request.URL.Query().Get("nonce")
	q.Echostr = this.Request.URL.Query().Get("echostr")
	q.Signature = this.Request.URL.Query().Get("signature")
	q.Timestamp = this.Request.URL.Query().Get("timestamp")
	q.EncryptType = this.Request.URL.Query().Get("encrypt_type")
	q.MsgSignature = this.Request.URL.Query().Get("msg_signature")

	this.Query = q
}

// 根据 Token Timestamp Nonce 生成对应的校验码， Token是不能明文传输的
func (this *WeixinClient) hashcode() string {

	strs := sort.StringSlice{this.Token, this.Query.Timestamp, this.Query.Nonce} // 使用本地的token生成校验
	sort.Strings(strs)
	str := ""
	for _, s := range strs {
		str += s
	}
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// 读取消息，解析XML
func (this *WeixinClient) initMessage() error {

	body, err := ioutil.ReadAll(this.Request.Body)

	if err != nil {
		return err
	}

	m, err := mxj.NewMapXml(body)

	if err != nil {
		return err
	}

	if _, ok := m["xml"]; !ok {
		return errors.New("Invalid Message.")
	}

	message, ok := m["xml"].(map[string]interface{})

	if !ok {
		return errors.New("Invalid Field `xml` Type.")
	}

	this.Message = message // 保存消息

	logger.Info(this.Message)

	return nil
}

func (this *WeixinClient) text() {

	inMsg, ok := this.Message["Content"].(string) // 读取内容

	if !ok {
		return
	}

	if inMsg == "【收到不支持的消息类型，暂无法显示】"{
		inMsg = "大兄弟，你好坏， 居然发表情包给我/::D"
	}

	var reply TextMessage

	reply.InitBaseData(this, "text")
	reply.Content = value2CDATA(fmt.Sprintf("%s", inMsg)) // 把消息再次封装

	replyXml, err := xml.Marshal(reply) // 序列化

	if err != nil {
		logger.Info(err)
		this.ResponseWriter.WriteHeader(403)
		return
	}

	this.ResponseWriter.Header().Set("Content-Type", "text/xml") // 数据类型text/xml
	this.ResponseWriter.Write(replyXml)                          // 回复微信平台
}

func (this *WeixinClient) image() {

	inMsg, ok := this.Message["MediaId"].(string) // 读取内容

	logger.Info("inMsg == ",inMsg)

	if !ok {
		return
	}

	var reply ImageMessage

	reply.InitBaseData(this, "image")
	reply.Image.MediaId = value2CDATA(fmt.Sprintf("%s", inMsg)) // 把消息再次封装

	replyXml, err := xml.Marshal(reply) // 序列化

	if err != nil {
		logger.Info(err)
		this.ResponseWriter.WriteHeader(403)
		return
	}

	this.ResponseWriter.Header().Set("Content-Type", "text/xml") // 数据类型text/xml
	this.ResponseWriter.Write(replyXml)                          // 回复微信平台
}

func (this *WeixinClient) voice() {

	inMsg, ok := this.Message["MediaId"].(string) // 读取内容

	logger.Info("inMsg == ",inMsg)

	if !ok {
		return
	}

	var reply VoiceMessage

	reply.InitBaseData(this, "voice")
	reply.Voice.MediaId = value2CDATA(fmt.Sprintf("%s", inMsg)) // 把消息再次封装

	replyXml, err := xml.Marshal(reply) // 序列化

	if err != nil {
		logger.Info(err)
		this.ResponseWriter.WriteHeader(403)
		return
	}

	this.ResponseWriter.Header().Set("Content-Type", "text/xml") // 数据类型text/xml
	this.ResponseWriter.Write(replyXml)                          // 回复微信平台
}

func (this *WeixinClient) othreInfo(msg string) {


	var info string

	switch msg {

	case "video":
		info = "我知道你给我发视频，虽然我还看不了，相信你很喜欢吧"
	case "shortvideo":
		info = "我知道你给我发短视频，虽然我还看不了，相信你很喜欢吧"
	case "link":
		info = "我知道你给我发小链接，虽然我还看不了，相信你很喜欢吧"
	case "location":
		info = "你给我发的地位，我可能暂时还不能去接你..."

	default:
		info = "大兄弟，我还在长身体，还在学习中..."
	}


	var reply TextMessage

	reply.InitBaseData(this, "text")
	reply.Content = value2CDATA(fmt.Sprintf("%s", info)) // 把消息再次封装

	replyXml, err := xml.Marshal(reply) // 序列化

	if err != nil {
		logger.Info(err)
		this.ResponseWriter.WriteHeader(403)
		return
	}

	this.ResponseWriter.Header().Set("Content-Type", "text/xml") // 数据类型text/xml
	this.ResponseWriter.Write(replyXml)                          // 回复微信平台
}

func (this *WeixinClient) Run() {

	err := this.initMessage()

	if err != nil {

		logger.Info(err)
		this.ResponseWriter.WriteHeader(403)
		return
	}

	MsgType, ok := this.Message["MsgType"].(string)

	if !ok {
		this.ResponseWriter.WriteHeader(403)
		return
	}

	switch MsgType {
	case "text":
		logger.Info(" text 消息类型")
		this.text() // 处理文本消息
	case "image":
		logger.Info(" image 消息类型")
		this.image() //处理图片消息
	case "voice":
		logger.Info(" voice 消息类型")
		this.voice()
	default:
		logger.Warn("其他的消息类型")
		this.othreInfo(MsgType) //处理不支持的消息类型
	}

	return
}
