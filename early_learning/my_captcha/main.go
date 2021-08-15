package main

import (
	"github.com/dchest/captcha"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

const filePath = "./my_captcha.html"

func readHtml() string {
	var bytes []byte
	var err error
	if bytes, err = ioutil.ReadFile(filePath); err != nil {
		log.Fatalf("ioutil.ReadFile error filePath =  %s , err :"+filePath, err)
		return ""
	}

	return string(bytes)
}

// 读取html 文件，转成template.Template 指针
var formTemplate = template.Must(template.New("myCaptcha").Parse(readHtml()))

// 显示验证码
func showCaptcha(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	d := struct {
		CaptchaId string
	}{
		captcha.New(),
	}
	// Execute将解析后的模板应用到指定的数据对象，并将输出写入wr
	if err := formTemplate.Execute(w, &d); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// 处理验证码，跳转结果页面
func resultPage(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if !captcha.VerifyString(r.FormValue("captchaId"), r.FormValue("captchaSolution")) {
		io.WriteString(w, "错误的验证码，请重新输入\n")
	} else {
		io.WriteString(w, "验证吗正确，你很棒哦！！\n")
	}
	io.WriteString(w, "<br><a href='/'>再试一下</a>")
}

func main() {

	// 简单设置log参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)


	// HandleFunc registers the handler function for the given pattern
	// in the DefaultServeMux.
	// The documentation for ServeMux explains how patterns are matched.
	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	//	DefaultServeMux.HandleFunc(pattern, handler)
	//}
	http.HandleFunc("/", showCaptcha)
	http.HandleFunc("/processCapcha", resultPage)

	http.Handle("/captcha/", captcha.Server(captcha.StdWidth, captcha.StdHeight))

	log.Println("starting server : 8888")

	if err := http.ListenAndServe("localhost:8888", nil); err != nil {
		log.Fatal(err)
	}
}
